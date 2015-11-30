package db_test

import (
	"os"
	"testing"
	"time"

	"github.com/ariejan/firedragon/server/db"
	"github.com/ariejan/firedragon/server/model"
)

var database = &db.DB{}

func TestGetItem(t *testing.T) {
	item, err := database.GetItem("url")

	if err != nil {
		t.Error(err)
	}

	if item.Code != "url" {
		t.Error("GetItem() did not return an Item with the proper ID.")
	}
}

func TestItemExists(t *testing.T) {
	if !database.DoesItemExist("url") {
		t.Error("Expected item 'url' to exist.")
	}

	if database.DoesItemExist("not-found") {
		t.Error("Expected item 'not found' not to exist.")
	}
}

func TestGetItems(t *testing.T) {
	items, err := database.GetItems()

	if err != nil {
		t.Error(err)
	}

	if len(items) != 2 {
		t.Errorf("Expected to find 2 items, got %d instead.", len(items))
	}
}

func TestGetItemNotFound(t *testing.T) {
	item, err := database.GetItem("nope")

	if item != nil {
		t.Error("Expected item to be `nil` when it does not exist.")
	}

	if err == nil {
		t.Error("Expected an error when unable to find an Item by ID.")
	}
}

func TestSaveItem(t *testing.T) {
	var err error
	var actual *model.Item

	item := &model.Item{
		Code:      "bam",
		Content:   "Lorem Ipsum",
		Type:      model.TextItemType,
		CreatedAt: time.Now(),
	}

	_, err = database.SaveItem(item)

	if err != nil {
		t.Error("Did not expect an error saving Item", err)
	}

	actual, err = database.GetItem("bam")
	if actual == nil || err != nil {
		t.Error("Expected to be able to retrieve just saved Item.")
	}

	database.DeleteItem(item.Code)
}

func TestSaveItemAutoShortCode(t *testing.T) {
	var err error
	var actual *model.Item

	// Bolt's NextSequence starts at 1 by default, let's assume "1" and "2"
	// have already been taken.
	database.SaveItem(&model.Item{
		Code:      "1",
		Content:   "Original 1",
		Type:      model.TextItemType,
		CreatedAt: time.Now(),
	})
	database.SaveItem(&model.Item{
		Code:      "2",
		Content:   "Original 2",
		Type:      model.TextItemType,
		CreatedAt: time.Now(),
	})

	// Test items with explicit and implicit blank codes
	item1 := &model.Item{
		Code:      "",
		Content:   "Lorem Ipsum",
		Type:      model.TextItemType,
		CreatedAt: time.Now(),
	}

	item2 := &model.Item{
		Content:   "Lorem Ipsum",
		Type:      model.TextItemType,
		CreatedAt: time.Now(),
	}

	// Test with explicit blank value
	actual, err = database.SaveItem(item1)
	if err != nil {
		t.Error("Did not expect an error saving Item", err)
	}
	if actual.Code != "3" {
		t.Error("Expected saveItem to auto-generate shortcode '3', but got '%s' instead.", err, actual.Code)
	}

	// TEst with implicit blank value
	actual, err = database.SaveItem(item2)
	if err != nil {
		t.Error("Did not expect an error saving Item", err)
	}
	if actual.Code != "4" {
		t.Error("Expected saveItem to auto-generate shortcode '4', but got '%s' instead.", err, actual.Code)
	}
}

func TestDeleteItem(t *testing.T) {
	database.SaveItem(&model.Item{
		Code:      "bam",
		Content:   "Lorem Ipsum",
		Type:      model.TextItemType,
		CreatedAt: time.Now(),
	})

	if !database.DoesItemExist("bam") {
		t.Error("Expected 'bam' to be saved to database.")
	}

	err := database.DeleteItem("bam")

	if err != nil {
		t.Error("failed deleting item")
		t.Error(err)
	}

	if database.DoesItemExist("bam") {
		t.Error("Expected 'bam' to be deleted.")
	}
}

func setup() {
	database.Open("testing.db")

	items := []*model.Item{
		&model.Item{Code: "url", Type: model.URLItemType, Content: "https://ariejan.net", CreatedAt: time.Now()},
		&model.Item{Code: "txt", Type: model.TextItemType, Content: "Lorem ipsum", CreatedAt: time.Now()},
	}

	for _, item := range items {
		database.SaveItem(item)
	}
}

func teardown() {
	os.Remove(database.Path())
	database.Close()
}

func TestMain(m *testing.M) {
	setup()

	retCode := m.Run()

	teardown()

	os.Exit(retCode)
}
