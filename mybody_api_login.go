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
	ID       string `json:"id"`
	NAME     string `json:"name"`
	PASSWORD string `json:"password"`
}

func (p *userInfo) createUserInfo(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	var id string
	var name string
	var jsonArr []interface{}

	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/mybodydiary")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}

	rows, err := db.Query("SELECT id, name FROM user")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}

		var obj map[string]interface{}
		err_json := json.Unmarshal([]byte("{}"), &obj)
		if err_json != nil {
			fmt.Println(err_json)
			return
		}

		obj["id"] = id
		obj["name"] = name

		jsonArr = append(jsonArr, obj)
	}

	defer db.Close()

	jsonArrVal, _ := json.Marshal(jsonArr)
	fmt.Fprint(w, string(jsonArrVal))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var jsonArr []interface{}

	var user userInfo

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/mybodydiary")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}

	stmt, err := db.Query("insert into user values (?, SHA2(?, 256), ?, now())",
		user.ID, user.PASSWORD, user.NAME)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	defer db.Close()

	var obj map[string]interface{}
	err_json := json.Unmarshal([]byte("{}"), &obj)
	if err_json != nil {
		fmt.Println(err_json)
		return
	}

	obj["result"] = 0
	obj["msg"] = "success"

	jsonArr = append(jsonArr, obj)

	jsonArrVal, _ := json.Marshal(jsonArr)
	fmt.Fprint(w, string(jsonArrVal))
}
