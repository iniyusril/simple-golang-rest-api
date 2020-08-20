package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/users", returnAllUsers).Methods("GET")
	http.Handle("/", router)
	router.HandleFunc("/user", insertUsersMultipart).Methods("POST")
	router.HandleFunc("/users", updateUsersMultipart).Methods("PUT")
	router.HandleFunc("/users", deleteUsersMultipart).Methods("DELETE")
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))

}
