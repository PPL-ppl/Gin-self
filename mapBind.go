package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func mapBind(ctx *gin.Context) {
	ids := ctx.QueryMap("ids")
	names := ctx.PostFormMap("names")
	fmt.Printf("ids: %v; names: %v", ids, names)
}
