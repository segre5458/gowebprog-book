package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	b := r.Body
	fmt.Fprintln(w, h)
	fmt.Fprintln(w, " ")
	fmt.Fprintln(w, b)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
