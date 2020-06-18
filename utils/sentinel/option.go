package middleware

import "github.com/gin-gonic/gin"

type (
	Option  func(*options)
	options struct {
		resourcePrefix     func(*gin.Context) string
		blockFallbackMap   map[string]func(*gin.Context) // 流量被限制回调函数
		breakerFallbackMap map[string]func(*gin.Context) // 熔断之后的回调函数
	}
)

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	for _, opt := range opts {
		opt(optCopy)
	}

	return optCopy
}

// WithResourcePrefix sets the resource prefix
func WithResourcePrefix(fn func(*gin.Context) string) Option {
	return func(opts *options) {
		opts.resourcePrefix = fn
	}
}

// WithBlockFallback sets the fallback handler when requests are blocked.
func WithBlockFallback(fnMap map[string]func(ctx *gin.Context)) Option {
	return func(opts *options) {
		opts.blockFallbackMap = fnMap
	}
}

// WithBreakerFallback sets the fallback handler when requests are breaker.
func WithBreakerFallback(fnMap map[string]func(ctx *gin.Context)) Option {
	return func(opts *options) {
		opts.breakerFallbackMap = fnMap
	}
}
