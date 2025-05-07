package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leonardo-luz/basic-goolang-api/pkg/service/user"
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

	// NOTE: Routes
	userService := user.NewHandler()
	userService.RegisterRoutes(subrouter.PathPrefix("/users").Subrouter())

	log.Println("Server listenning on http://localhost", server.address)

	return http.ListenAndServe(server.address, router)
}
