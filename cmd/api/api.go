package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Splucheviy/TiagoEcomm/service/user"
	"github.com/gorilla/mux"
)

// APIServer...
type APIServer struct {
	addr string
	db   *sql.DB
}

// NewAPIServer ...
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Run ...
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}