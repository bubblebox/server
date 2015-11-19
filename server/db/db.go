package db

import (
	"encoding/json"
	"log"

	"github.com/ariejan/firedragon/server/model"
	"github.com/boltdb/bolt"
)

// DB wraps BoltDB and handles storing/retrieving data
type DB struct {
	*bolt.DB
}

// GetItem will retrieve an item from the database, given it's unique Code
func (db *DB) GetItem(code string) (*model.Item, error) {
	item := &model.Item{Code: code}

	err := db.View(func(tx *bolt.Tx) error {
		data := tx.Bucket([]byte("items")).Get([]byte(code))

		if err := json.Unmarshal(data, item); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return item, nil
}

// GetItems returns all items, yes, all of them.
// TODO: Add pagination
func (db *DB) GetItems() ([]model.Item, error) {
	var items []model.Item

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("items"))
		bucket.ForEach(func(k, v []byte) error {
			item := model.Item{Code: string(k)}
			if err := json.Unmarshal(v, &item); err != nil {
				return err
			}

			items = append(items, item)

			return nil
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return items, nil
}

// SaveItem will save the item, if the item has no code yet, it will be
// automatically assigned one. If a code is provided and it already
// exists, an error will be returned.
func (db *DB) SaveItem(item *model.Item) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("items"))

		data, err := json.Marshal(item)
		if err != nil {
			log.Fatal("Could not marshal JSON")
			return err
		}

		err = bucket.Put([]byte(item.Code), data)

		return err
	})

	if err != nil {
		log.Println("Error saving item", err)
		return err
	}

	return nil
}

// DeleteItem delete the item with `id` from the database
func (db *DB) DeleteItem(id string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("items"))

		return bucket.Delete([]byte(id))
	})

	return err
}

// Open will open a new Bolt database and create the necessary buckets
// for using it elsewhere in Fire Dragon.
func (db *DB) Open(name string) error {
	var err error
	db.DB, err = bolt.Open(name, 0600, nil)
	if err != nil {
		return err
	}

	// Create buckets
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("items"))
		if err != nil {
			log.Panic("Error creating bucket for items", err)
		}

		return nil
	})

	if err != nil {
		db.Close()
		return err
	}

	return nil
}
