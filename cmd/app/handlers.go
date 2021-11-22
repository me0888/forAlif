package app

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/me0888/forAlif/pkg"
	"io"
	"log"
	"net/http"
	"strconv"
)

const secret = "~,6a=HFaC/P]R5Zp"

func (s *Server) handleGetAccount(writer http.ResponseWriter, request *http.Request) {
	id := Auth(writer, request)
	if id == 0 {
		log.Println("you are not authorized")
		return
	}

	var account *accounts.Account
	body, _ := io.ReadAll(request.Body)

	if ok, _ := Verify(body, secret, request.Header.Get("X-Digest")); !ok {
		log.Println("Tampered")
		status := map[string]string{"status": "fail", "message": "tampered"}
		writeJSON(writer, status, http.StatusBadRequest)
		return
	} else {
		log.Println("OK")
	}

	err := json.Unmarshal(body, &account)
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item := s.accountsSvc.CheckAccount(request.Context(), account.Phone)
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if item {
		status := map[string]string{"status": "ok", "exsist": "true"}
		writeJSON(writer, status, 200)
	} else {
		status := map[string]string{"status": "fail", "exist": "false"}
		writeJSON(writer, status, 404)
	}

}

func (s *Server) handleDeposit(writer http.ResponseWriter, request *http.Request) {

	id := Auth(writer, request)
	if id == 0 {
		log.Println("you are not authorized")
		return
	}

	body, _ := io.ReadAll(request.Body)

	if ok, _ := Verify(body, secret, request.Header.Get("X-Digest")); !ok {
		log.Println("Tampered")

		status := map[string]string{"status": "fail", "message": "tampered"}
		writeJSON(writer, status, http.StatusBadRequest)
		return
	} else {
		log.Println("OK")
	}

	identified, err := s.accountsSvc.Identified(request.Context(), strconv.Itoa(int(id)))
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var maxBalance float64 = 10.00

	if identified {
		maxBalance = 100.00
	}

	var item1 struct {
		Ammount float64 `json:"ammount"`
	}
	err = json.Unmarshal(body, &item1)
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	balance, err := s.accountsSvc.Balance(request.Context(), strconv.Itoa(int(id)))
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if balance+item1.Ammount > maxBalance {
		status := map[string]string{"status": "fail", "message": "maximum balance for your account " +
			strconv.FormatFloat(maxBalance, 'f', 2, 64)}
		writeJSON(writer, status, http.StatusBadRequest)
		return
	}

	_, err = s.accountsSvc.Deposit(request.Context(), id, item1.Ammount)
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	status := map[string]string{"status": "ok", "ammount": strconv.FormatFloat(item1.Ammount, 'f', 2, 64)}
	writeJSON(writer, status, 200)

}

func (s *Server) handleCountAndSum(writer http.ResponseWriter, request *http.Request) {

	id := Auth(writer, request)
	if id == 0 {
		log.Println("you are not authorized")
		return
	}

	count, sum, err := s.accountsSvc.Amount(request.Context(), strconv.Itoa(int(id)))
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	status := map[string]string{"status": "ok", "count": strconv.Itoa(int(count)), "sum": strconv.FormatFloat(sum, 'f', 2, 64)}
	writeJSON(writer, status, 200)

}

func (s *Server) handleBalance(writer http.ResponseWriter, request *http.Request) {

	id := Auth(writer, request)
	if id == 0 {
		log.Println("you are not authorized")
		return
	}

	balance, err := s.accountsSvc.Balance(request.Context(), strconv.Itoa(int(id)))
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	status := map[string]string{"status": "ok", "balance": strconv.FormatFloat(balance, 'f', 2, 64)}
	writeJSON(writer, status, 200)

}

func Auth(writer http.ResponseWriter, request *http.Request) int64 {
	id, err := Authentication(request.Context())
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 0
	}

	if id == 0 {
		log.Print("you are not authorized")
		http.Error(writer, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return 0
	}
	return id

}

func writeJSON(writer http.ResponseWriter, item interface{}, code int) {
	writer.WriteHeader(code)
	data, err := json.Marshal(item)
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)
	if err != nil {
		log.Print(err)
	}
}

func Verify(msg []byte, key, hash string) (bool, error) {
	sig, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}

	mac := hmac.New(sha256.New, []byte(key))
	mac.Write(msg)

	return hmac.Equal(sig, mac.Sum(nil)), nil
}
