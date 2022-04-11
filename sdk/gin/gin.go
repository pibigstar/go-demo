package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"log"
	"net/http"
	"reflect"
)

func main() {
	r := gin.Default()

	// 设置gin mode
	gin.SetMode(gin.DebugMode)

	// 修改默认路由打印格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("【%v】====> %v\n", httpMethod, absolutePath)
	}

	// 全局中间件，阻止panic
	r.Use(gin.Recovery())

	// 为某个group下的路由注册中间件
	r.Group("/admin").Use(func(context *gin.Context) {
		if id := context.GetHeader("adminUserId"); id == "" {
			context.JSON(http.StatusForbidden, "请登录管理员")
		}
	})

	// 开启pprof
	pprof.Register(r)

	// 自定义校验
	// 将我们自定义的校验方法注册到 validator中
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("NotNullAndAdmin", nameNotNullAndAdmin); err != nil {
			fmt.Println("RegisterValidation NotNullAndAdmin", err.Error())
		}
	}

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func nameNotNullAndAdmin(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if value, ok := field.Interface().(string); ok {
		// 字段不能为空，并且不等于  admin
		return value != "" && "admin" != value
	}
	return true
}

type Req struct {
	Name string `validate:"NotNullAndAdmin"`
}
