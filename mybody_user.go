package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type userInfo struct {
	ID       string `json:"id"`
	NAME     string `json:"name"`
	PASSWORD string `json:"password"`
}

func CheckUser(id string, pw string) int {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/mybodydiary")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}

	rows, err := db.Query("select id, password, name from user where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user = userInfo{}

	if !rows.Next() {
		log.Fatal("Cannot find match id.", err)
		return -201
	} else {
		_ = rows.Scan(&user.ID, &user.PASSWORD, &user.NAME)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(pw))
	if err != nil {
		//log.Fatal("Wrong password. err : ", err)
		return -202
	}

	return 0
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
	var user userInfo
	var obj map[string]interface{}
	err_json := json.Unmarshal([]byte("{}"), &obj)
	if err_json != nil {
		fmt.Println(err_json)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/mybodydiary")
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}

	bs, err := bcrypt.GenerateFromPassword([]byte(user.PASSWORD), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Query("insert into user values (?, ?, ?, now())",
		user.ID, bs, user.NAME)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	defer db.Close()

	obj["result"] = 0
	obj["msg"] = "success"

	jsonVal, _ := json.Marshal(obj)
	fmt.Fprint(w, string(jsonVal))
}

type customErr struct {
	Code    string
	Message string
}

func (e *customErr) Error() string {
	return e.Code + ", " + e.Message
}

func (e *customErr) StatusCode() int {
	result, _ := strconv.Atoi(e.Code)
	return result
}
