package api_test

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/ariejan/firedragon/server/api"
	"github.com/ariejan/firedragon/server/db"
	"github.com/ariejan/firedragon/server/model"
	"github.com/gin-gonic/gin"
)

var database = &db.DB{}
var router *gin.Engine

func setup() {
	database.Open("testing.db")

	timeStamp, _ := time.Parse(time.RFC3339Nano, "2015-11-19T12:19:33.865042825+01:00")

	items := []*model.Item{
		&model.Item{
			Code:      "url",
			Type:      model.URLItemType,
			Content:   "https://ariejan.net",
			CreatedAt: timeStamp,
		},
		&model.Item{
			Code:      "txt",
			Type:      model.TextItemType,
			Content:   "Lorem ipsum",
			CreatedAt: timeStamp,
		},
	}

	for _, item := range items {
		database.SaveItem(item)
	}

	router = gin.Default()
	api.Setup(router.Group("/api/v1"), database)
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
