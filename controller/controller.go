package controller

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//CreateUser creates a user.. kkkkkk
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.Write([]byte("Error on body request!"))
		return
	}

	var userRequest user

	if err = json.Unmarshal(bodyRequest, &userRequest); err != nil {
		rw.Write([]byte("Error on converting user!"))
		return
	}

	database, err := db.Connect()

	if err != nil {
		rw.Write([]byte("Error on connect on db!"))
		return
	}
	defer database.Close()

	statement, err := database.Prepare("insert into users (name,email) values (?,?)")

	if err != nil {
		rw.Write([]byte("Error on create the statement!"))
	}
	defer statement.Close()

	insert, err := statement.Exec(userRequest.Name, userRequest.Email)

	if err != nil {
		rw.Write([]byte("Error on create the user!"))
	}

	userID, err := insert.LastInsertId()
	if err != nil {
		rw.Write([]byte("Erro on show the user ID"))
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(fmt.Sprintf("User successfully created with id: %d", userID)))
}
