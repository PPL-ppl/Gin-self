package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadHTMLGlob1(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website1",
	})
}
func LoadHTMLGlob2(c *gin.Context) {
	c.HTML(http.StatusOK, "post/index.tmpl", gin.H{
		"title": "Main website2",
	})
}
func LoadHTMLGlob3(c *gin.Context) {
	c.HTML(http.StatusOK, "user/index.tmpl", gin.H{
		"title": "Main website3",
	})
}
