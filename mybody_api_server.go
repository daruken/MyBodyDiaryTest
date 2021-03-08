package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Post struct {
	NAME    string `json:"name"`
	STADIUM string `json:"stadium"`
}

var db *sql.DB
var err error

func apiServer() {
	fmt.Printf("[API Server Section]\n")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", apiServerStatus)
	router.HandleFunc("/lookup", testShow).Methods("GET")
	router.HandleFunc("/user_info", showUserInfo).Methods("GET")
	log.Fatal(http.ListenAndServe(":28080", router))
	fmt.Println("[" + time.Now().Format(time.RFC3339) + "][API Server started")
}

func testShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	json.NewEncoder(w).Encode(map[string]string{"Club": "Manu"})
}

func apiServerStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	message := "[" + time.Now().Format(time.RFC3339) + "][API Server is online]\n"
	fmt.Fprintf(w, message)
}

func showUserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Header().Set("Content-Type", "application/json")

	db, err = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/mybodydiary")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Query("select id, password, email from user_info")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var posts []Post

	for result.Next() {
		var post Post
		err := result.Scan(&post.NAME, &post.STADIUM)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
		fmt.Print(post.NAME)
		fmt.Println(" ", post.STADIUM)
	}

	json.NewEncoder(w).Encode(posts)
}

func main() {
	fmt.Printf("[Application started]\n")
	apiServer()
	fmt.Printf("[Application exit]\n")
}
