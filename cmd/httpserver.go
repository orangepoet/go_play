package cmd

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var dataChan = make(chan string, 10)

func HttpServer() {
	route := gin.Default()
	route.GET("/health", health)
	route.GET("/pk", pk)
	route.GET("/genpk", genpk)
	route.Run("localhost:8080")
}

func pk(context *gin.Context) {
	select {
	case data := <-dataChan:
		context.JSON(http.StatusOK, data)
	case <-time.After(60 * time.Second):
		context.JSON(http.StatusOK, "timeout")
	}
}

func genpk(context *gin.Context) {
	data := context.Query("data")
	dataChan <- data
	context.JSON(http.StatusOK, "post data: "+data)
}

func health(context *gin.Context) {
	context.JSON(http.StatusOK, "ok")
}
