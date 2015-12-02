package web

import (
	"net/http"

	"github.com/ariejan/firedragon/server/db"
	"github.com/gin-gonic/gin"
)

// Setup configures how to handle routes
func Setup(group *gin.RouterGroup, db *db.DB) {
	group.Use()
	{
		group.GET("/", renderItemHandler(db))
		group.GET("/:shortcode", renderItemHandler(db))
	}
}

func renderItemHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement rendering content...
		c.String(http.StatusTeapot, "I'm a teapot")
	}
}
