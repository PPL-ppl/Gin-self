package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回值是数组值加前缀
// engine.SecureJsonPrefix(")]}',\n")
// 将输出：)]}',
//["lena","austin","foo"]

func someJSON(c *gin.Context) {
	names := []string{"lena", "austin", "foo"}
	// 将输出：while(1);["lena","austin","foo"]
	c.SecureJSON(http.StatusOK, names)
}
