package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func loginUserHandler(w http.ResponseWriter, r *http.Request) {
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

	if r.Method == http.MethodPost {
		err := CheckUser(user.ID, user.PASSWORD)
		if err != 0 {
			obj["result"] = err
			obj["msg"] = "Failed to login."

			jsonVal, _ := json.Marshal(obj)
			fmt.Fprint(w, string(jsonVal))
		}

		//http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	obj["result"] = 0
	obj["msg"] = "success"

	jsonVal, _ := json.Marshal(obj)
	fmt.Fprint(w, string(jsonVal))
}
