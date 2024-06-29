package main

import (
	"net/http"

	"github.com/a-h/templ/examples/integration-gin/gintemplrenderer"
	"github.com/gin-gonic/gin"

	"dobefu/search-page/functions"
	"dobefu/search-page/structs"
	"dobefu/search-page/templates"
	"dobefu/search-page/templates/htmx"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	ginHtmlRenderer := router.HTMLRender
	router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	router.SetTrustedProxies([]string{"localhost"})

	var streets = functions.GetStreets()

	router.GET("/", func(c *gin.Context) {
		var query structs.QueryParams
		if c.ShouldBindQuery(&query) == nil {
			streets = functions.FilterStreets(streets, query)
		}

		c.HTML(http.StatusOK, "", templates.Index(streets, query))
	})

	router.GET("/htmx/results.html", func(c *gin.Context) {
		var query structs.QueryParams
		if c.ShouldBindQuery(&query) == nil {
			streets = functions.FilterStreets(streets, query)
		}

		c.HTML(http.StatusOK, "", htmx.HtmxResults(streets))
	})

	router.Run("localhost:6060")
}
