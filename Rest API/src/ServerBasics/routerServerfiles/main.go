package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	server := httprouter.New()
	server.ServeFiles("/*filepath", http.Dir("/Users/manoj/Documents/"))
	log.Fatal(http.ListenAndServe(":8080", server))
}
