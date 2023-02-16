package main

import "github.com/gin-gonic/gin"

func PostFormBind(context *gin.Context) {
	// 在这种情况下，将自动选择合适绑定  表单提交
	user := context.PostForm("user")
	password := context.PostForm("password")
	if user == "user" && password == "password" {
		context.JSON(200, gin.H{"status": "you are logged in"})
	} else {
		context.JSON(401, gin.H{"status": "unauthorized"})
	}
}
