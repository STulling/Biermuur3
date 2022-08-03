package api

import (
	"STulling/video/display/controller"
	"embed"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getEffects(c *gin.Context) {
	locale := c.Param("locale")
	c.JSON(http.StatusOK, effects[locale])
}

func getActivations(c *gin.Context) {
	locale := c.Param("locale")
	c.JSON(http.StatusOK, activations[locale])
}

func setAction(c *gin.Context) {
	action := c.Param("action")
	go controller.SetCallback(action)
	c.String(http.StatusOK, "OK")
}

func setActivation(c *gin.Context) {
	activation := c.Param("activation")
	go controller.SetActivation(activation)
	c.String(http.StatusOK, "OK")
}

func serveFile(file string) func(*gin.Context) {
	return func(c *gin.Context) {
		c.FileFromFS(fmt.Sprintf("static%s", file), http.FS(static))
	}
}

func serveDir(dir string) func(*gin.Context) {
	return func(c *gin.Context) {
		path := c.Param("path")
		c.FileFromFS(fmt.Sprintf("static/%s/%s", dir, path), http.FS(static))
	}
}

//go:embed static/*
var static embed.FS

func Run() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", serveFile("/"))
	router.GET("/favicon.png", serveFile("/favicon.png"))
	router.GET("/style.css", serveFile("/style.css"))
	router.GET("/components/:path", serveDir("components"))

	router.GET("/api/DJ/effects/:locale", getEffects)
	router.GET("/api/DJ/activations/:locale", getEffects)

	router.GET("/api/DJ/action/:action", setAction)
	router.GET("/api/DJ/activation/:activation", setActivation)
	fmt.Println("Starting...")
	router.Run("0.0.0.0:80")
}
