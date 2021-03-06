package api

import (
	"net/http"

	"github.com/bubblebox/server/db"
	"github.com/bubblebox/server/model"
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

func itemsCreateHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item *model.Item
		if c.BindJSON(&item) == nil {
			if storedItem, err := db.SaveItem(item); err == nil {
				c.JSON(http.StatusCreated, gin.H{
					"item": storedItem,
				})
			} else {
				c.String(http.StatusInternalServerError, "Could not create item.")
			}
		}
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

		c.JSON(http.StatusOK, gin.H{})
	}
}
