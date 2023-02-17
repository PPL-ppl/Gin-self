package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// MyBenchLogger 自定义中间件
func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("自定义中间件")
	}
}
