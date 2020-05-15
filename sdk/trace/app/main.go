package main

import (
	"github.com/gin-gonic/gin"
	"go-demo/sdk/trace/app/db"
	"go-demo/sdk/trace/app/middleware"
	"math/rand"
)

type User struct {
	Id       int    `json:"id"`
	Age      int    `json:"age"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (*User) TableName() string {
	return "user"
}

func main() {
	r := gin.Default()

	// 使用链路追踪
	r.Use(middleware.GinTrace())
	r.Use(middleware.DBTrace())

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "this is index")

		// 调用数据库
		user := &User{
			Id:       rand.Intn(100),
			Age:      rand.Intn(100),
			UserName: "test",
			Password: "123456",
		}
		err := db.DB.Create(user).Error
		if err != nil {
			ctx.String(200, err.Error())
		}
	})

	r.Run(":8081")
}
