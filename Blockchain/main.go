package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleGetBlockchain).Methods("GET")
	r.HandleFunc("/", handleWriteBlockchain).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
