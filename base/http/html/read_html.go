package main

import (
	"net/http"
)

// 访问静态资源
func main() {
	dir := http.Dir("base/http/html/root")
	staticHandler := http.FileServer(dir)
	http.Handle("/", http.StripPrefix("/", staticHandler))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
