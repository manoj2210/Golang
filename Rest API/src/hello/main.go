package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hellofunc)
	http.HandleFunc("/mmk", kk)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hellofunc(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hello World\n")
}

func kk(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello\n")
}
