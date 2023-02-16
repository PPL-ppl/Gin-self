package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func someXML(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}
