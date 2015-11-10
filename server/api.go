package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shwoodard/jsonapi"
)

// ItemHandler returns a handler function for redirecting or
// rendering items to end users.
func ItemHandler(db *DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		item, err := db.GetItem(vars["code"])
		if err != nil {
			http.Error(w, "Not found", 404)
		}

		http.Redirect(w, r, item.Content, 301)
	}
}

// ViewItemHandler handles viewing a single item
func ViewItemHandler(db *DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		item, err := db.GetItem(vars["code"])

		if err != nil {
			http.Error(w, "Not found", 404)
			return
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/vnd.api+json")
		if err := jsonapi.MarshalOnePayload(w, item); err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}

// CreateItemHandler handles creating a new Item
func CreateItemHandler(db *DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		item := new(Item)

		if err := jsonapi.UnmarshalPayload(r.Body, item); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Save it

		w.WriteHeader(201)
		w.Header().Set("Content-Type", "application/vnd.api+json")

		if err := jsonapi.MarshalOnePayload(w, item); err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}
