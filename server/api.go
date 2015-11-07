package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ItemHandler returns a handler function for redirecting or
// rendering items to end users.
func ItemHandler(db *DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		item, err := db.GetItem(vars["code"])
		if err != nil {
			http.Error(w, "Cannot load that content", 404)
		}

		http.Redirect(w, r, item.Content, 301)

	}
}
