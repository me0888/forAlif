package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/me0888/forAlif/cmd/app"
	"github.com/me0888/forAlif/pkg"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	host := "0.0.0.0"
	port := "9999"
	dns := "postgres://app:pass@localhost:5432/db"
	if err := Execute(host, port, dns); err != nil {
		os.Exit(1)
	}
}

func Execute(host string, port string, dns string) (err error) {
	connCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	pool, err := pgxpool.Connect(connCtx, dns)
	if err != nil {
		log.Println(err)
		return
	}

	defer pool.Close()

	mux := mux.NewRouter()
	managersSvc := accounts.NewService(pool)
	server := app.NewServer(mux, managersSvc)
	server.Init()

	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: server,
	}
	println("srv starts")
	return srv.ListenAndServe()
}
