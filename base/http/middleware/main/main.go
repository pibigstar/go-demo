package main

import (
	"fmt"
	"go-demo/base/http/middleware"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", middleware.BodyLimit(hello))

	mux.HandleFunc("/admin", middleware.Auth(hello))

	go fmt.Println("server staring...")

	if err := http.ListenAndServe(":8081", middleware.IPRateLimit(mux)); err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
