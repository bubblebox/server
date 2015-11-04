package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port = ":8042"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", r)
	log.Printf("Starting Fire Dragon server on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
