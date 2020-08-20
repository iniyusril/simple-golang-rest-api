package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var users Users
var arr_user []Users
var response Response

func returnAllUsers(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select id,first_name,last_name from person")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_user = append(arr_user, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

//code lanjutan dari tutorial part 1

func insertUsersMultipart(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t Users
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.FirstName)
	log.Println(t.LastName)
	// Do something with the Person struct...

	// var response Response

	db := connect()
	defer db.Close()

	_, err = db.Exec("INSERT INTO person (first_name, last_name) values (?,?)",
		t.FirstName,
		t.LastName,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Add"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func updateUsersMultipart(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var t Users
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.FirstName)
	log.Println(t.LastName)

	_, err = db.Exec("UPDATE person set first_name = ?, last_name = ? where id = ?",
		t.FirstName,
		t.LastName,
		t.Id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Update Data"
	log.Print("Update data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func deleteUsersMultipart(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var t Users
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.FirstName)
	log.Println(t.LastName)

	_, err = db.Exec("DELETE from person where id = ?",
		t.Id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Delete Data"
	log.Print("Delete data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
