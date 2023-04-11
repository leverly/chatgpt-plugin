package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"plugin/handler"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/logo.png", handler.PluginLogo)
	router.GET("/.well-known/ai-plugin.json", handler.PluginManifest)
	router.GET("/openapi.yaml", handler.OpenapiSpec)

	router.POST("/todos/:username", handler.AddTodo)
	router.GET("/todos/:username", handler.GetTodos)
	router.DELETE("/todos/:username", handler.DeleteTodo)

	if err := router.Run("localhost:5003"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
