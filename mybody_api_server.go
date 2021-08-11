package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {}

func (a *App) Run(addr string) {}

func NewHandler() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/users", usersHandler).Methods("GET")

	return mux
}
