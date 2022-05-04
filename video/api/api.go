package api

import (
	"STulling/video/display/controller"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getEffects(c *gin.Context) {
	locale := c.Param("locale")
	c.JSON(http.StatusOK, effects[locale])
}

func setAction(c *gin.Context) {
	action := c.Param("action")
	go controller.SetCallback(action)
	c.String(http.StatusOK, "OK")
}

func Run() {
	router := gin.Default()
	router.Use(cors.Default())

	router.StaticFile("/", "./api/static/index.html")
	router.StaticFile("/favicon.png", "./api/static/favicon.png")
	router.StaticFile("/style.css", "./api/static/style.css")
	router.Static("/components", "./api/static/components")

	router.GET("/api/DJ/effects/:locale", getEffects)
	router.GET("/api/DJ/:action", setAction)
	fmt.Println("Starting...")
	router.Run("0.0.0.0:80")
}
