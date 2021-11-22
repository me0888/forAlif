package accounts

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

var ErrNotFound = errors.New("item not found")
var ErrNoSuchUser = errors.New("no such user")
var ErrInvalidPassword = errors.New("invalid password")
var ErrInternal = errors.New("internal error")

type Service struct {
	pool *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

type Account struct {
	ID         int64     `json:"id"`
	Phone      string    `json:"phone"`
	Password   string    `json:"password"`
	Balance    float64   `json:"balance"`
	Identified bool      `json:"identified"`
	Created    time.Time `json:"created"`
}

func (s *Service) CheckAccount(ctx context.Context, phone string) bool {
	var item string
	err := s.pool.QueryRow(ctx, `SELECT phone FROM accounts WHERE phone = $1;`, phone).Scan(&item)
	if err != nil {
		if item != phone {
			return false
		}
	}

	return true
}

func (s *Service) Deposit(ctx context.Context, id int64, ammount float64) (float64, error) {
	var balance float64
	_, err := s.pool.Exec(ctx, `INSERT INTO replenishment (accounts_id, ammount) VALUES ($1, $2);
		`, id, ammount)
	if err != nil {
		log.Print(err)
		return 0, ErrInternal
	}

	log.Println(ammount)
	log.Println(id)
	err = s.pool.QueryRow(ctx, `UPDATE accounts SET balance=balance+$1 WHERE id=$2 returning balance;`,
		ammount, id).Scan(&balance)
	if err != nil {
		log.Print(err)
		return 0, ErrInternal
	}

	return balance, nil
}

func (s *Service) Amount(ctx context.Context, id string) (int64, float64, error) {
	var count int64
	var sum float64

	err := s.pool.QueryRow(ctx, `
	SELECT COUNT(ammount), SUM(ammount) FROM replenishment 
	WHERE accounts_id = $1 
	AND created>=date_trunc('month', CURRENT_DATE) 
	AND created<=date_trunc('month',CURRENT_DATE)+'1month'::interval-'1sec'::INTERVAL;`, id).Scan(&count, &sum)

	if err != nil {
		log.Println(err)
		return 0, 0, ErrInternal
	}

	return count, sum, nil
}

func (s *Service) Balance(ctx context.Context, id string) (float64, error) {
	var balance float64
	err := s.pool.QueryRow(ctx, `SELECT balance FROM accounts WHERE id = $1;`, id).
		Scan(&balance)
	if err != nil {
		return 0, ErrInternal
	}
	return balance, nil
}

func (s *Service) IDByToken(ctx context.Context, token string) (id int64, err error) {
	var expire bool

	err = s.pool.QueryRow(ctx, `
	SELECT accounts_id, now()>expire as expire FROM accounts_tokens WHERE token =$1;`,
		token).Scan(&id, &expire)

	if errors.Is(err, pgx.ErrNoRows) {
		log.Println(err)
		return 0, ErrNoSuchUser
	}
	if err != nil {
		log.Println(err)
		return 0, ErrInternal
	}

	return id, nil
}

func (s *Service) Identified(ctx context.Context, id string) (bool, error) {
	var identified bool

	err := s.pool.QueryRow(ctx, `SELECT identified FROM accounts WHERE id =$1;`, id).Scan(&identified)

	if errors.Is(err, pgx.ErrNoRows) {
		log.Println(err)
		return false, ErrNoSuchUser
	}
	if err != nil {
		log.Println(err)
		return false, ErrInternal
	}

	return identified, nil
}
