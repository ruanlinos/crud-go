package controller

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

//ListAllUsers return all users in the database
func ListAllUsers(rw http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		rw.Write([]byte("Error on connect to the db."))
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from users")

	if err != nil {
		rw.Write([]byte("Error on get the users!"))
		return
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user

		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			rw.Write([]byte("Error on scan the list of users"))
			return
		}

		users = append(users, user)
	}
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(users); err != nil {
		rw.Write([]byte("Error on encode slice."))
		return
	}

}

//ListOneUser return a user based on your id.
func ListOneUser(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		rw.Write([]byte("Error on convert the param to integer"))
	}

	db, err := db.Connect()

	if err != nil {
		rw.Write([]byte("Error on connect to db."))
	}

	row, err := db.Query("select * from users where id = ?", ID)

	if err != nil {
		rw.Write([]byte("Error on get the specified user"))
	}
	var user user
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			rw.Write([]byte("Error on scan the user"))
			return
		}
	}
	if err := json.NewEncoder(rw).Encode(user); err != nil {
		rw.Write([]byte("Error on encode user."))
		return
	}
}
