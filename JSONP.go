package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSONP(context *gin.Context) {
	data := map[string]interface{}{
		"foo": "bar",
	}
	// 输出 : {"foo":"bar"}
	context.JSON(http.StatusOK, data)
}
