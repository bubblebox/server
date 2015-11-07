package main

import (
	"encoding/json"
	"log"

	"github.com/boltdb/bolt"
)

// DB wraps BoltDB and handles storing/retrieving data
type DB struct {
	*bolt.DB
}

// GetItem will retrieve an item from the database, given it's unique Code
func (db *DB) GetItem(code string) (*Item, error) {
	item := &Item{Code: code}

	err := db.View(func(tx *bolt.Tx) error {
		data := tx.Bucket([]byte("items")).Get([]byte(code))

		err := json.Unmarshal(data, item)

		if err != nil {
			log.Fatal("Could not unmarshal JSON")
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return item, nil
}

// SaveItem will save the item, if the item has no code yet, it will be
// automatically assigned one. If a code is provided and it already
// exists, an error will be returned.
func (db *DB) SaveItem(item *Item) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("items"))

		data, err := json.Marshal(item)
		if err != nil {
			log.Fatal("Could not marshal JSON")
			return err
		}

		err = bucket.Put([]byte(item.Code), data)

		return nil
	})

	if err != nil {
		log.Println("Error saving item", err)
		return err
	}

	return nil
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
