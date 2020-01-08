package main

import (
	"io"
	"net/http"
)

type customMux int

func (customMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		io.WriteString(w, "Hi this is from customMux you came to /")
		return
	}
	http.NotFound(w, r)
	return
}
func main() {
	var c customMux
	http.ListenAndServe(":8080", c)
}
