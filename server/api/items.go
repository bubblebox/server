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
		}

		c.JSON(http.StatusOK, gin.H{
			"items": items,
		})
	}
}
