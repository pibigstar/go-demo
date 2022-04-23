package ctxkit

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"google.golang.org/grpc/metadata"
	"net/http"
)

type ctxHeaderKey struct{}

func (ctxHeaderKey) String() string {
	return "Header-Ctx-Key"
}

func SetHeaderMiddleware(c *gin.Context) {
	c.Request = c.Request.WithContext(SetValue(c.Request.Context(), ctxHeaderKey{}, c.Request.Header))
	c.Next()
}

func AppendHeader(ctx context.Context, key string, value interface{}) context.Context {
	header := make(http.Header)
	if value := GetValue(ctx, ctxHeaderKey{}); value != nil {
		if err := gconv.Struct(value, &header); err != nil {
			return ctx
		}
	}
	header.Add(key, gconv.String(value))
	return SetValue(ctx, ctxHeaderKey{}, header)
}

func SetValue(ctx context.Context, key interface{}, value interface{}) context.Context {
	newCtx := context.WithValue(ctx, key, value)
	md := metadata.New(map[string]string{
		gconv.String(key): gconv.String(value),
	})
	return metadata.NewOutgoingContext(newCtx, md)
}

func GetValue(ctx context.Context, key interface{}) interface{} {
	if value := ctx.Value(key); value != nil {
		return value
	}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if values := md.Get(gconv.String(key)); len(values) > 0 {
			return values[0]
		}
	}

	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		if values := md.Get(gconv.String(key)); len(values) > 0 {
			return values[0]
		}
	}
	return nil
}

func GetHeader(ctx context.Context, key string) interface{} {
	if value := GetValue(ctx, ctxHeaderKey{}); value != nil {
		var header http.Header
		if err := gconv.Struct(value, &header); err != nil {
			return nil
		}
		return header.Get(key)
	}
	return nil
}
