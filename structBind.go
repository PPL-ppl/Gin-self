package main

import "github.com/gin-gonic/gin"

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func structBind(ctx *gin.Context) {
	var log Login
	ctx.BindJSON(&log)
}
