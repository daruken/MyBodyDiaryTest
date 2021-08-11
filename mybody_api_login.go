package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type userInfo struct {
	EMAIL    string `json:"email"`
	ID       string `json:"id"`
	PASSWORD string `json:"password"`
}

func (p *userInfo) createUserInfo(db *sql.DB) error {
	return errors.New("Not implemented")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	var email string
	var id string
	var jsonArr []interface{}

	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/mybodydiary")

	rows, err := db.Query("SELECT email, id FROM user")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&email, &id)

		if err != nil {
			log.Fatal(err)
		}

		var obj map[string]interface{}
		err_json := json.Unmarshal([]byte("{}"), &obj)
		if err_json != nil {
			fmt.Println(err_json)
			return
		}

		obj["email"] = email
		obj["id"] = id

		jsonArr = append(jsonArr, obj)
	}

	defer db.Close()

	jsonArrVal, _ := json.Marshal(jsonArr)
	fmt.Fprint(w, string(jsonArrVal))
}
