package web_test

import (
	"os"
	"testing"
	"time"

	"github.com/bubblebox/server/db"
	"github.com/bubblebox/server/model"
	"github.com/bubblebox/server/web"
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
	web.Setup(router.Group("/"), database)
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
