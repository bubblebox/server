package resource

import (
	"fmt"
	"net/http"

	"github.com/ariejan/firedragon/server/storage"
	"github.com/manyminds/api2go"
)

// ItemResource for api2go routes
type ItemResource struct {
	ItemStorage *storage.ItemStorage
}

// FindOne item
func (i ItemResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := i.ItemStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	return &Response{Res: res}, nil
}

func (i ItemResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	res, err := i.ItemStorage.GetAll()
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	return &Response{Res: res}, nil
}

func (i ItemResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (i ItemResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (i ItemResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return nil, fmt.Errorf("Not implemented")
}
