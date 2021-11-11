package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type ResponseUserInfo struct {
	RET  int        `json:"result"`
	USER []UserInfo `json:"user"`
}

type UserInfo struct {
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

	var user = UserInfo{}

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

func (p *UserInfo) createUserInfo(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	var id string
	var name string
	var responseUserInfo ResponseUserInfo

	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/mybodydiary")
	if err != nil {
		customError(w, -101, err.Error())

		return
	}

	rows, err := db.Query("SELECT id, name FROM user")
	if err != nil {
		defer db.Close()
		customError(w, -102, err.Error())

		return
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			defer db.Close()
			return
		}

		responseUserInfo.USER = append(responseUserInfo.USER, UserInfo{ID: id, NAME: name})
	}

	defer db.Close()

	responseUserInfo.RET = 0
	jsonArrVal, _ := json.Marshal(responseUserInfo)
	fmt.Fprint(w, string(jsonArrVal))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user UserInfo
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

func customError(w http.ResponseWriter, errCode int, errMsg string) {
	var errObj map[string]interface{}
	err_json := json.Unmarshal([]byte("{}"), &errObj)
	if err_json != nil {
		fmt.Println(err_json)
		return
	}

	errObj["result"] = errCode
	errObj["msg"] = errMsg

	jsonArrVal, _ := json.Marshal(errObj)
	fmt.Fprint(w, string(jsonArrVal))
}
