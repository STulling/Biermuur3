package api

import (
	"STulling/video/display/controller"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setAction(c *gin.Context) {
	action := c.Param("action")
	go controller.SetCallback(action)
	c.String(http.StatusOK, "OK")
}

func Run() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/DJ/:action", setAction)
	fmt.Println("Starting...")
	router.Run("0.0.0.0:5000")
}
