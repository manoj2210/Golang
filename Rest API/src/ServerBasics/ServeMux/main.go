package main

import (
	"io"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome")
}
func curse(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Go to hell")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/curse", curse)
	http.ListenAndServe(":8080", mux)
}
