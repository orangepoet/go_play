package cmd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpServer() {
	route := gin.Default()
	route.GET("/health", health)
	route.Run("localhost:8080")
}

func health(context *gin.Context) {
	context.JSON(http.StatusOK, "ok")
}
