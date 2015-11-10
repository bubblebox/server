package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	port = ":8042"
)

// Router configures a new mux.Router and returns it for routing HTTP requests.
func Router(db *DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/items", CreateItemHandler(db)).Methods("POST")
	r.HandleFunc("/api/v1/items/{code}", ViewItemHandler(db)).Methods("GET")

	r.HandleFunc("/{code}", ItemHandler(db))

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	return r
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	db := &DB{}
	db.Open("firedragon.db")
	defer db.Close()

	seedData(db)

	http.Handle("/", Router(db))

	log.Printf("Starting Fire Dragon server on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func seedData(db *DB) {
	var items = []*Item{
		&Item{Code: "url", Type: URLItemType, Content: "https://ariejan.net", CreatedAt: time.Now()},
		&Item{Code: "txt", Type: TextItemType, Content: "Lorem ipsum", CreatedAt: time.Now()},
	}

	for _, item := range items {
		db.SaveItem(item)
	}
}
