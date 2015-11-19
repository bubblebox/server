package api

import (
	"net/http"

	"github.com/ariejan/firedragon/server/db"
	"github.com/gin-gonic/gin"
)

func itemsIndexHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := db.GetItems()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"items": []string{},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"items": items,
		})
	}
}

func itemsShowHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")
		item, err := db.GetItem(code)

		if err != nil {
			c.String(http.StatusNotFound, "Item not found.")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"item": item,
		})
	}
}

func itemsDestroyHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")
		err := db.DeleteItem(code)

		if err != nil {
			c.String(http.StatusNotFound, "Item not found.")
			return
		}

		c.String(http.StatusOK, "")
	}
}
