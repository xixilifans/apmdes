package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin/v2"
)

func main() {
	engine := gin.New()
	engine.Use(apmgin.Middleware(engine))

	engine.GET("/albums", getAlbums)
	engine.Run(":8081")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "ok")
}
