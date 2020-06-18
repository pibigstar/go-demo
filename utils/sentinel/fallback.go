package middleware

import "github.com/gin-gonic/gin"

// 熔断回调函数
var BreakerFallbackMap = map[string]func(*gin.Context){
	"/ping": func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(403, map[string]interface{}{
			"err":  "服务暂时不可用，请稍后重试",
			"code": 403,
		})
	},
}

// 流量控制回调函数
var BlockFallbackMap = map[string]func(*gin.Context){
	"/ping": func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(400, map[string]interface{}{
			"err":  "请求太多，来不及处理了，请稍后重试",
			"code": 429,
		})
	},
}
