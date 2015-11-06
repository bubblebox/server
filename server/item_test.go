package main_test

import (
	"os"
	"testing"

	"github.com/ariejan/firedragon/server"
)

var item = &main.Item{
	Code:    "abc",
	Type:    main.TextItemType,
	Content: "Lorem ipsum",
}

func TestItemPersistence(t *testing.T) {
	var err error

	db := main.DB{}
	db.Open("testing.db")
	defer os.Remove(db.Path())
	defer db.Close()

	err = db.SaveItem(item)
	if err != nil {
		t.Error("Expected to save a new item successfully")
	}

	result, err := db.GetItem("abc")
	if err != nil {
		t.Error("Expected to retrieve item successfully")
	}

	if *result == *item {
		t.Errorf("Expected %v but got %v", *item, *result)
	}

}
