package decorator

import (
	"net/http"
)

// 方法的装饰器模式, 包装某个方法
// http 鉴权
func Auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("token")
		if name == "pi" {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(403)
		}

	}
}

func f(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
