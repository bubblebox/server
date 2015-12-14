package api

import (
	"github.com/bubblebox/server/db"
	"github.com/gin-gonic/gin"
)

// Setup will setup all API routes.
func Setup(group *gin.RouterGroup, db *db.DB) {
	group.Use()
	{
		group.GET("/items/:code", itemsShowHandler(db))
		group.DELETE("/items/:code", itemsDestroyHandler(db))
		group.GET("/items", itemsIndexHandler(db))
		group.POST("/items", itemsCreateHandler(db))
	}
}
