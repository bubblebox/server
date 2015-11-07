package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port = ":8042"
)

// Router configures a new mux.Router and returns it for routing HTTP requests.
func Router(db *DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/{code}", ItemHandler(db))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	return r
}

func main() {
	db := &DB{}
	db.Open("firedragon.db")
	defer db.Close()

	http.Handle("/", Router(db))

	log.Printf("Starting Fire Dragon server on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
