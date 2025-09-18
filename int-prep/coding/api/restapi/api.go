package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// simple rest api server using standard net http library

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Employees []Employee

func addEmployee(w http.ResponseWriter, r *http.Request) {

	var emp Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	Employees = append(Employees, emp)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(emp)
}

func listEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Employees)
}

func main() {

	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			listEmployees(w, r)
		} else {
			addEmployee(w, r)
		}
	})

	port := 8080
	fmt.Println("Starting http server on port: ", port)

	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Panic(err)
	}

}
