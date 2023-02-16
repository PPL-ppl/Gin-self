package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Query query?id=1234&page=2
func Query(context *gin.Context) {
	id := context.Query("id")
	page := context.DefaultQuery("page", "0")
	name := context.PostForm("name")
	message := context.PostForm("message")
	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
}
