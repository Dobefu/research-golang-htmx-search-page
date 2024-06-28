package main

import (
	"net/http"

	"github.com/a-h/templ/examples/integration-gin/gintemplrenderer"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	ginHtmlRenderer := router.HTMLRender
	router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	router.SetTrustedProxies([]string{"localhost"})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", Index())
	})

	router.Run("localhost:6060")
}
