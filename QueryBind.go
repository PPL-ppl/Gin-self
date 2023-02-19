package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func queryBind(ctx *gin.Context) {
	firstname := ctx.DefaultQuery("hello", "2")
	lastname := ctx.Query("lastname")
	ctx.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
