package main

import (
	"Gin-self/gorm"
	"Gin-self/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func U(ctx *gin.Context) {
	var user model.User
	conn := gorm.GetConn()
	conn.First(&user, 10)
	user.Age = user.Age - 1
	if user.Age > 0 {
		fmt.Println(user.Age)
		conn.Save(&user)
	}
	ctx.JSON(http.StatusOK, user.Age)
}

func C(ctx *gin.Context) {
	user := model.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	conn := gorm.GetConn()
	conn.Create(&user) // 通过数据的指针来创建
	ctx.JSON(http.StatusOK, user.ID)
}

func D(ctx *gin.Context) {
	var user model.User
	user.ID = 1
	conn := gorm.GetConn()
	conn.Delete(&user)
	ctx.JSON(http.StatusOK, user.ID)
}
