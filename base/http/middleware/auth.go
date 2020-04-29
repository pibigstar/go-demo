package middleware

import "net/http"

// 鉴权，判断用户是否登录
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
