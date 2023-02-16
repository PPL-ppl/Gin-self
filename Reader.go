package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Reader(c *gin.Context) {
	response, err := http.Get("https://bkimg.cdn.bcebos.com/pic/eac4b74543a98226cffca68ba4ceae014a90f603bd41?x-bce-process=image/watermark,image_d2F0ZXIvYmFpa2UyNzI=,g_7,xp_5,yp_5")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}
