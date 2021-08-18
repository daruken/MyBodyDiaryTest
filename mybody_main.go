package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Start server ...")
	router := mux.NewRouter()
	router.HandleFunc("/users", getUserInfoHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}
