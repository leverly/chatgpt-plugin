package handler

import "github.com/gin-gonic/gin"

func PluginLogo(c *gin.Context) {
	c.File("logo.png")
}

func PluginManifest(c *gin.Context) {
	c.File("well-known/ai-plugin.json")
}

func OpenapiSpec(c *gin.Context) {
	c.File("openapi.yaml")
}
