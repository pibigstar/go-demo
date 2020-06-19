package middleware

import (
	"net/http"
	"strings"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
)

func SentinelMiddleware(opts ...Option) gin.HandlerFunc {
	options := evaluateOptions(opts)
	return func(ctx *gin.Context) {
		resource := ctx.FullPath()

		if options.resourcePrefix != nil {
			resource = strings.ReplaceAll(resource, options.resourcePrefix(ctx), "")
		}
		flowResource := FlowPrefix + resource
		breakerResource := BreakerPrefix + resource

		// 流量控制
		flowEntry, err := api.Entry(
			flowResource,
			api.WithResourceType(base.ResTypeWeb),
			api.WithTrafficType(base.Inbound),
		)
		if err != nil {
			if fn, ok := options.blockFallbackMap[resource]; ok {
				fn(ctx)
			} else {
				// 默认失败回调
				ctx.AbortWithStatus(http.StatusTooManyRequests)
			}

			return
		}
		defer flowEntry.Exit()

		// 熔断检查
		breakerEntry, err := api.Entry(breakerResource)
		if err != nil {
			if fn, ok := options.breakerFallbackMap[resource]; ok {
				fn(ctx)
			} else {
				// 默认失败回调
				ctx.AbortWithStatus(http.StatusForbidden)
			}
			return
		}
		defer breakerEntry.Exit()

		ctx.Next()
	}
}
