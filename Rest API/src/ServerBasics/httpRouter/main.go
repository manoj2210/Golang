package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func greet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	io.WriteString(w, "Welcome")
	fmt.Println(params)
}
func curse(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	io.WriteString(w, "Go to hell")
}

func main() {
	server := httprouter.New()
	server.GET("/greet/:name/:class", greet)
	server.GET("/curse", curse)

	log.Fatal(http.ListenAndServe(":8080", server))
}
