package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func moreJSON(ctx *gin.Context) {
	// 你也可以使用一个结构体
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 123
	// 注意 msg.Name 在 JSON 中变成了 "user"
	// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}
	ctx.JSON(http.StatusOK, msg)
}
