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

	log.Fatal(http.ListenAndServe(":5000", router))

}
