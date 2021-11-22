package app

import (
	"github.com/gorilla/mux"
	"github.com/me0888/forAlif/pkg"
	"net/http"
)

type Server struct {
	mux         *mux.Router
	accountsSvc *accounts.Service
}

func NewServer(mux *mux.Router, accountsSvc *accounts.Service) *Server {
	return &Server{mux: mux, accountsSvc: accountsSvc}
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.mux.ServeHTTP(writer, request)
}

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
)

func (s *Server) Init() {
	s.mux.Use(Authenticate(s.accountsSvc.IDByToken))
	s.mux.HandleFunc("/account", s.handleGetAccount).Methods(POST)
	s.mux.HandleFunc("/deposit", s.handleDeposit).Methods(POST)
	s.mux.HandleFunc("/getsum", s.handleCountAndSum).Methods(GET)
	s.mux.HandleFunc("/balance", s.handleBalance).Methods(GET)

}
