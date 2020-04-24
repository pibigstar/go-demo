package main

import (
	"fmt"
	"net/http"
)

// 每秒可接收1个请求，最大可运行5个请求
var limiter = NewIPRateLimiter(1, 5)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	if err := http.ListenAndServe(":8888", limitMiddleware(mux)); err != nil {
		panic(err)
	}
}

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 获取IP限速器
		limiter := limiter.GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World")
}
