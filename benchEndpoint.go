package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func benchEndpoint(c *gin.Context) {
	fmt.Println("benchEndpoint")
}
