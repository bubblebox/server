package api_test

import (
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/ariejan/firedragon/server/api"
	"github.com/ariejan/firedragon/server/db"
	"github.com/ariejan/firedragon/server/model"
)

var database = &db.DB{}
var handler http.Handler

func setup() {
	database.Open("testing.db")

	items := []*model.Item{
		&model.Item{ID: "url", Type: model.URLItemType, Content: "https://ariejan.net", CreatedAt: time.Now()},
		&model.Item{ID: "txt", Type: model.TextItemType, Content: "Lorem ipsum", CreatedAt: time.Now()},
	}

	for _, item := range items {
		database.SaveItem(item)
	}

	handler = api.Handler(8888, database)
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

func matchJSON(actual string, expected string) bool {
	var aval interface{}
	var eval interface{}

	json.Unmarshal([]byte(actual), &aval)
	json.Unmarshal([]byte(expected), &eval)

	return reflect.DeepEqual(aval, eval)
}
