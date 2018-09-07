package main

import (
	"github.com/zabawaba99/firego"
	"log"
)

func main() {
	fb := firego.New("https://go-demo-2d035.firebaseio.com", nil)

	v := map[string]interface{}{
		"name":     "SuperWang",
		"location": "Beijing",
		"age":      28,
		"Likes":    []string{"Movie", "Football"},
	}
	if err := fb.Set(v); err != nil {
		log.Fatal(err)
	}
}
