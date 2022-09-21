package main

import (
	"encoding/json"

	"log"
	"net/http"
)

type Todo struct {
	ID   int
	Task string
}

var todos []Todo

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.Header().Add("Access-Control-Allow-Origin", "*")

		switch r.Method {
		case "GET":
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(todos)

		}
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
