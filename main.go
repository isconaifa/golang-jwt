package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"jwt-token/handlers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("The server is running on port 4000")
	router := mux.NewRouter()

	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/protected", handlers.ProtectedHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", router))
}
