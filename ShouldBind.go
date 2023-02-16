package main

import "github.com/gin-gonic/gin"

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func ShouldBind(context *gin.Context) {
	var form LoginForm
	// 在这种情况下，将自动选择合适绑定  Json
	if context.ShouldBind(&form) == nil {
		if form.User == "user" && form.Password == "password" {
			context.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			context.JSON(401, gin.H{"status": "unauthorized"})
		}
	}
}
