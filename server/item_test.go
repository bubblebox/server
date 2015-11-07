package main_test

import (
	"testing"

	"github.com/ariejan/firedragon/server"
)

func TestItemPersistence(t *testing.T) {
	var err error

	item := &main.Item{
		Code:    "abc",
		Type:    main.TextItemType,
		Content: "Lorem ipsum",
	}

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
