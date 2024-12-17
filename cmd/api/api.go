package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leonardo-luz/basic-goolang-api/service/user"
)

type ApiServer struct {
	address string
	db      *sql.DB
}

func NewApiServer(address string, db *sql.DB) *ApiServer {
	return &ApiServer{
		address: address,
		db:      db,
	}
}

func (server *ApiServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userService := user.NewHandler()
	userService.RegisterRoutes(subrouter)

	log.Println("Server listenning on", server.address)

	return http.ListenAndServe(server.address, router)
}
