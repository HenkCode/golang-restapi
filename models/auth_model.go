package models

import (
	"database/sql"
	"fmt"
	"github.com/HenkCode/golang-restapi/db"
	"github.com/HenkCode/golang-restapi/helpers"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckAuth(username, password string) (bool, error) {
	var obj User
	var pass string
	connection := db.CreateConf()
	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := connection.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pass,
	)

	if err == sql.ErrNoRows {
		fmt.Println("username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("query error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pass)
	if !match {
		fmt.Println("hash password doesn't match.")
		return false, err
	}

	return true, nil
}