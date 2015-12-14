package web

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/bubblebox/server/db"
	"github.com/bubblebox/server/model"
	"github.com/gin-gonic/gin"
)

const tpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>BubbleBox</title>
	<link href="https://cdnjs.cloudflare.com/ajax/libs/pure/0.6.0/pure-min.css" rel="stylesheet">
	<style>
		.wrapper {
			width: 980px;
			margin-left: auto;
			margin-right: auto;
		}
	</style>
  </head>
  <body>
    <div class="wrapper">
		<h1>Fire Dragon</h1>
		<hr>
		<div class="content">
			<pre>
     		{{.Content}}
			</pre>
	    </div>
		<hr>
		<footer>Hosted by <a href="https://github.com/bubblebox">BubbleBox</a></footer>
	</div>
  </body>
</html>`

var textTmpl = template.Must(template.New("textContent").Parse(tpl))

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
		case model.TextItemType:
			var output bytes.Buffer
			textTmpl.Execute(&output, item)
			if output.Len() == 0 {
				c.String(http.StatusNotFound, "Something went wrong rendering")
				return
			}
			c.Data(http.StatusOK, "text/html; charset=utf-8", output.Bytes())
			return
		default:
			c.String(http.StatusNotFound, "Not found")
			return
		}
	}
}
