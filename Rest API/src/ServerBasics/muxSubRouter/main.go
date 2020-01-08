package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func greet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Welcome %v %v", vars["name"], vars["id"])
}

func main() {
	Mux := mux.NewRouter()
	s := Mux.PathPrefix("/greet").Subrouter()
	s.HandleFunc("/{name}/{id:[0-9][6-9]}", greet)
	server := &http.Server{
		Handler: Mux,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
