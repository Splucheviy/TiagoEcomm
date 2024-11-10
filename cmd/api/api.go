package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Splucheviy/TiagoEcomm/service/product"
	"github.com/Splucheviy/TiagoEcomm/service/user"
	"github.com/gorilla/mux"
)

// Server ...
type Server struct {
	addr string
	db   *sql.DB
}

// NewServer ...
func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

// Run ...
func (s *Server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
