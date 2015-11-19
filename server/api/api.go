package api

import (
	"net/http"

	"github.com/ariejan/firedragon/server/db"
	"github.com/ariejan/firedragon/server/model"
	"github.com/ariejan/firedragon/server/resource"
	"github.com/ariejan/firedragon/server/storage"
	"github.com/gorilla/mux"
	"github.com/manyminds/api2go"
)

// Handler returns the http handler for Fire Dragon
func Handler(port int, db *db.DB) http.Handler {
	itemStorage := storage.NewItemStorage(db)

	api := api2go.NewAPI("/api/v1/")
	api.AddResource(model.Item{}, resource.ItemResource{ItemStorage: itemStorage})

	handler := api.Handler()

	r := mux.NewRouter()
	r.PathPrefix("/api/v1/").Handler(handler)
	r.PathPrefix("/dashboard/").Handler(http.StripPrefix("/dashboard/", http.FileServer(http.Dir("./public/"))))

	return r
}
