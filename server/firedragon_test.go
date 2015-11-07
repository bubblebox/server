package main_test

import (
	"os"
	"testing"
	"time"

	"github.com/ariejan/firedragon/server"
)

var db = &main.DB{}

var items = []*main.Item{
	&main.Item{Code: "url", Type: main.URLItemType, Content: "https://ariejan.net", CreatedAt: time.Now()},
	&main.Item{Code: "txt", Type: main.TextItemType, Content: "Lorem ipsum", CreatedAt: time.Now()},
}

func setup() {
	db.Open("testing.db")
	loadFixtureData()
}

func teardown() {
	os.Remove(db.Path())
	db.Close()
}

func TestMain(m *testing.M) {
	setup()

	retCode := m.Run()

	teardown()

	os.Exit(retCode)
}

func loadFixtureData() {
	for _, item := range items {
		db.SaveItem(item)
	}
}
