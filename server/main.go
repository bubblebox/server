package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ariejan/firedragon/server/api"
	"github.com/ariejan/firedragon/server/db"
	"github.com/ariejan/firedragon/server/ember"
	"github.com/ariejan/firedragon/server/model"
	"github.com/ariejan/firedragon/server/web"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
)

var (
	port           = 8042
	domain         = "127.0.0.1.xip.io"
	apiSubdomain   = "api"   // api.domain
	adminSubdomain = "admin" // admin.domain
	dbName         = "firedragon.db"

	// Generated domain names
	portString  = fmt.Sprintf(":%d", port)
	contentHost = fmt.Sprintf("%s:%d", domain, port)
	apiHost     = fmt.Sprintf("%s.%s:%d", apiSubdomain, domain, port)
	adminHost   = fmt.Sprintf("%s.%s:%d", adminSubdomain, domain, port)

	corsConfig = cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}
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

	// Configure routers
	apiRouter := gin.Default()
	apiRouter.Use(cors.Middleware(corsConfig))
	api.Setup(apiRouter.Group("/api/v1"), db)

	adminRouter := gin.Default()
	adminRouter.Use(cors.Middleware(corsConfig))
	ember.Setup(adminRouter.Group("/"))

	contentRouter := gin.Default()
	web.Setup(contentRouter.Group("/"), db)

	// Setup subdomain routing
	hs := make(HostSwitch)
	hs[apiHost] = apiRouter
	hs[adminHost] = adminRouter
	hs[contentHost] = contentRouter

	// Start HTTP server
	log.Print(">> Server up and running")
	log.Printf(">> Content server - http://%s", contentHost)
	log.Printf(">> API server     - http://%s", apiHost)
	log.Printf(">> Admin server   - http://%s", adminHost)
	http.ListenAndServe(portString, hs)
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
