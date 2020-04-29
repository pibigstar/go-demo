package middleware

import (
	"encoding/json"
	"io"
	"net/http"
)

// 限制body长度
func BodyLimit(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var maxLength int64 = 128
		var body = make(map[string]interface{})

		err := json.NewDecoder(io.LimitReader(r.Body, maxLength)).Decode(&body)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Body length illegal"))
		} else {
			handler.ServeHTTP(w, r)
		}
	}
}
