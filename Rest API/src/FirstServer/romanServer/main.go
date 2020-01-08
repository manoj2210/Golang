package main

import (
	"FirstServer/romanNumeral"
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Create a server and run it on 8080 port
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/", romanfunc)
	s.ListenAndServe()

}

//Use Supervisor or Gulps
func romanfunc(w http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/")

	if urlPathElements[1] == "roman_number" {
		number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
		if number == 0 || number > 10 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not Found"))
		} else {
			fmt.Fprintf(w, "%q",
				html.EscapeString(romanNumeral.Numerals[number]))
		}
	} else {
		// For all other requests, tell that Client sent a bad request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad request"))
	}
	fmt.Println(urlPathElements)
}
