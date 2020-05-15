package middleware

import (
	"github.com/gin-gonic/gin"
	"go-demo/sdk/trace/app/db"
)

func DBTrace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db.AddContext(ctx)
	}
}
