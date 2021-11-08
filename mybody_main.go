package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Start server ...")
	router := mux.NewRouter()
	router.HandleFunc("/gateway/users", getUserInfoHandler).Methods("GET")
	router.HandleFunc("/gateway/users", createUserHandler).Methods("POST")
	router.HandleFunc("/gateway/login", loginUserHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}
