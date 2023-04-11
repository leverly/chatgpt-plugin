package main

import "github.com/gin-gonic/gin"

func pluginLogo(c *gin.Context) {
	c.File("logo.png")
}

func pluginManifest(c *gin.Context) {
	c.File(".well-known/ai-plugin.json")
}

func openapiSpec(c *gin.Context) {
	c.File("openapi.yaml")
}
