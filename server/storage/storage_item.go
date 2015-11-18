package storage

import (
	"fmt"

	"github.com/ariejan/firedragon/server/db"
	"github.com/ariejan/firedragon/server/model"
)

// ItemStorage handle storage of Items
type ItemStorage struct {
	db *db.DB
}

// NewItemStorage returns a new ItemStorage using the specified database.
func NewItemStorage(db *db.DB) *ItemStorage {
	return &ItemStorage{db: db}
}

// GetOne item
func (is ItemStorage) GetOne(id string) (*model.Item, error) {
	item, err := is.db.GetItem(id)
	if err != nil {
		return nil, fmt.Errorf("Item with ID '%s' not found", id)
	}

	return item, nil
}

func (is ItemStorage) GetAll() ([]*model.Item, error) {
	items, err := is.db.GetItems()
	if err != nil {
		return nil, fmt.Errorf("Cannot fetch items")
	}

	return items, nil
}
