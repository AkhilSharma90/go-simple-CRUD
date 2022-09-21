package main

import (
	"fmt"
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
}

func main() {

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
