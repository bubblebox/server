package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ariejan/firedragon/server/api"
	"github.com/ariejan/firedragon/server/db"
	"github.com/ariejan/firedragon/server/ember"
	"github.com/ariejan/firedragon/server/model"
	"github.com/gin-gonic/gin"
)

var (
	port   = 8042
	dbName = "firedragon.db"
)

func main() {
	// Configure logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	// Setup bolt database
	db := &db.DB{}
	db.Open(dbName)
	defer db.Close()

	// Create seed data
	// TODO: Replace this with an optional CLI command to seed data.
	seedData(db)

	// Setup Gin
	router := gin.Default()
	api.Setup(router.Group("/api/v1"), db)
	ember.Setup(router.Group("/dashboard"))

	// Start HTTP server
	log.Printf("Listening on :%d", port)
	router.Run(fmt.Sprintf(":%d", port))
}

func seedData(db *db.DB) {
	var items = []*model.Item{
		&model.Item{Code: "url", Type: model.URLItemType, Content: "https://ariejan.net", CreatedAt: time.Now()},
		&model.Item{Code: "txt", Type: model.TextItemType, Content: "Lorem ipsum", CreatedAt: time.Now()},
	}

	for _, item := range items {
		db.SaveItem(item)
	}
}
