package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	fmt.Println("Start server ...")
	router := mux.NewRouter()
	router.HandleFunc("/users", usersHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}
