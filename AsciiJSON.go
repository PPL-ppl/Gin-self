package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AsciiJSON(context *gin.Context) {
	data := map[string]interface{}{
		"lang": "go语言",
		"tag":  "<br>",
	}
	// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	//context.AsciiJSON(http.StatusOK, data)
	// 输出 : {"lang":"go语言","tag":"\u003cbr\u003e"}
	context.JSON(http.StatusOK, data)
}
