package gin_play

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var dataChan = make(chan string, 10)

func GinServer() {
	route := gin.Default()
	route.GET("/health", health)
	route.GET("/pk", pk)
	route.GET("/genpk", genpk)
	route.GET("long-polling", longPolling)
	_ = route.Run("localhost:8080")
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

func longPolling(c *gin.Context) {
	ch := make(chan string, 1)
	defer close(ch)

	select {
	case result := <-ch:
		log.Println("polling result", result)
		c.String(200, "Long polling response: %s", result)
	case <-time.After(5 * time.Second):
		log.Println("polling timeout")
		c.String(200, "Timeout waiting for data")
	}
}
