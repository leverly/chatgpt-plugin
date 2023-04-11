package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type Todo struct {
	Text string `json:"todo"`
}

var todos = make(map[string][]string)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/logo.png", pluginLogo)
	router.GET("/.well-known/ai-plugin.json", pluginManifest)
	router.GET("/openapi.yaml", openapiSpec)

	router.POST("/todos/:username", addTodo)
	router.GET("/todos/:username", getTodos)
	router.DELETE("/todos/:username", deleteTodo)

	if err := router.Run("localhost:5003"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
