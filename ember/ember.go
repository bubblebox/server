package ember

import "github.com/gin-gonic/gin"

// Setup will configure the router group for severing a static
// Ember.js application.
func Setup(group *gin.RouterGroup) {
	group.Use()
	{
		group.Static("", "./public/")
	}
}
