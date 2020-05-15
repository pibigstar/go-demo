package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

// gin 集成 trace
func GinTrace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tracer := NewTracer()
		if tracer == nil {
			ctx.Next()
			return
		}
		var span opentracing.Span
		extract, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		if err == nil {
			// 继承父span
			span = opentracing.StartSpan(ctx.Request.URL.Path, opentracing.ChildOf(extract))
		} else {
			// 初始化新的span
			span = opentracing.StartSpan(ctx.Request.URL.Path)
		}
		defer span.Finish()

		// 记录body信息
		if b := ctx.Request.GetBody; b != nil {
			if body, err := b(); err == nil {
				var buff bytes.Buffer
				if _, err = buff.ReadFrom(body); err == nil {
					span.LogKV("form-body", buff.String())
				}
			}
		}
		// 设置context
		if sp, ok := span.Context().(jaeger.SpanContext); ok {
			// 这里设置一个root span，本次请求的其他span挂在root span下面
			ctx.Set("root_span", span)
			ctx.Writer.Header().Set("x-request-id", sp.TraceID().String())
			ctx.Writer.Header().Set("x-trace-id", sp.TraceID().String())
			ctx.Writer.Header().Set("X-Span-id", sp.SpanID().String())
		}
		ctx.Next()
	}
}
