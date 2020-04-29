package middleware

import (
	"golang.org/x/time/rate"
	"net/http"
	"sync"
)

/**
基于 IP 限制 HTTP 访问频率
*/
var (
	ipLimitMaps = make(map[string]*rate.Limiter)
	mu          sync.Mutex
	rateLimit   = 1 // 每秒往池子填充的令牌数
	rateMax     = 5 // 池子最大的令牌数
)

func GetIPLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if limiter, ok := ipLimitMaps[ip]; ok {
		return limiter
	}
	limiter := rate.NewLimiter(rate.Limit(rateLimit), rateMax)
	ipLimitMaps[ip] = limiter

	return limiter
}

// 限制IP访问频率
// 作用在Server, 对所有访问进行限制
func IPRateLimit(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := GetIPLimiter(r.RemoteAddr)
		// 如果想不丢掉此次请求，请使用Wait方法
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
