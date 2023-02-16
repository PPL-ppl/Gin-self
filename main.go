package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	engine := gin.Default()
	//Default 使用 Logger 和 Recovery 中间件
	//engine := gin.Default()
	//不使用
	//engine:=gin.New()
	// 你也可以使用自己的 SecureJSON 前缀
	engine.GET("/AsciiJson", AsciiJSON)

	//返回Html
	engine.LoadHTMLGlob("templates/**/*")
	engine.GET("/LoadHTMLGlob1", LoadHTMLGlob1)
	engine.GET("/LoadHTMLGlob2", LoadHTMLGlob2)
	engine.GET("/LoadHTMLGlob3", LoadHTMLGlob3)

	//返回Html
	engine.Delims("{[{", "}]}")
	engine.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	engine.LoadHTMLFiles("testdata/template/raw.tmpl")
	engine.GET("/LoadHTMLFile1", LoadHTMLFile1)

	engine.GET("/JSONP", JSONP)

	//Json
	engine.POST("/login1byJson", ShouldBind)
	//表单提交
	engine.POST("/login2byForm", PostFormBind)

	//"html": "<b>Hello, world!</b>"
	// 提供 unicode 实体  {"html":"\u003cb\u003eHello, world!\u003c/b\u003e"}
	engine.GET("/JSON", JSON)
	// 提供字面字符  {"html":"<b>Hello, world!</b>"}
	engine.GET("/pureJson", pureJson)

	// http://localhost:8080/query?id=1234&page=2
	engine.POST("/query", Query)
	// 返回值是数组值加前缀
	engine.GET("/someJSON", someJSON)

	// 自定义返回结构
	engine.GET("/someJSON1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	// 返回自定义结构体
	engine.GET("/moreJSON", moreJSON)

	engine.GET("/someXML", someXML)

	engine.GET("/someYAML", someYAML)

	engine.GET("/someProtoBuf", someProtoBuf)
	//单文件上传
	engine.POST("/uploadOne", uploadOne)
	//多文件上传
	engine.POST("/uploadMore", uploadMore)
	
	//读取Reader数据
	engine.GET("/Reader", Reader)

	engine.Run(":8080")
}
