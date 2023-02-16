package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func someYAML(context *gin.Context) {
	context.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}
