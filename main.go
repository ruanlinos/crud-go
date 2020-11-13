package main

import (
	"crud/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", controller.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", controller.ListAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", controller.ListOneUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", controller.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", controller.DeleteUser).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":5000", router))

}
