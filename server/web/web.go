package web

import (
	"net/http"

	"github.com/ariejan/firedragon/server/db"
	"github.com/ariejan/firedragon/server/model"
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
		code := c.Param("shortcode")
		if code == "" {
			c.String(http.StatusOK, "Nothing to see here. Move along now, people.")
			return
		}

		item, _ := db.GetItem(code)
		if item == nil {
			c.String(http.StatusNotFound, "")
			return
		}

		switch item.Type {
		case model.URLItemType:
			c.Redirect(http.StatusMovedPermanently, item.Content)
			return
		default:
			c.String(http.StatusNotFound, "Not found")
			return
		}
	}
}
