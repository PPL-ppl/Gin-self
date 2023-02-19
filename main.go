package main

import (
	"Gin-self/middleware"
	"context"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	engine := gin.Default()
	//记录到文件
	create, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(create)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	//gin.DefaultWriter = io.MultiWriter(create, os.Stdout)

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

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

	authorized := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	//Auth
	authorized.GET("/Auth", Auth)
	//engine.Use(gin.Logger())
	// 你可以为每个路由添加任意数量的中间件。
	engine.GET("/benchmark", middleware.MyBenchLogger(), benchEndpoint)
	//// urlBind?name=22&address=2
	engine.GET("/urlBind", urlBind)

	engine.GET("/longAsync", longAsync)
	engine.GET("/longSync", longSync)

	//匹配多个不同body
	engine.POST("/SomeHandler", SomeHandler)
	engine.POST("/SomeHandler1", SomeHandler1)

	//绑定map
	engine.POST("/mapBind", mapBind)
	engine.GET("/queryBind", queryBind)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		if err := autotls.Run(engine, "example1.com", "example2.com"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		//log.Fatal()
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
