package main

import (
	"database/sql"
	"errors"
)

type userInfo struct {
	EMAIL    string `json:"email"`
	ID       string `json:"id"`
	PASSWORD string `json:"password"`
}

func (p *userInfo) createUserInfo(db *sql.DB) error {
	return errors.New("Not implemented")
}
