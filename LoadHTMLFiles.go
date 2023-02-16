package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
func LoadHTMLFile1(c *gin.Context) {
	c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
		"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
	})
}
