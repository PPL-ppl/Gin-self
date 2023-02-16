package main

import (
	"github.com/gin-gonic/gin"
)

func JSON(context *gin.Context) {
	context.JSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}
func pureJson(context *gin.Context) {
	context.PureJSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}
