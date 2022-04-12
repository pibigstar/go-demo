package main

import (
	_ "go-demo/sdk/gf/boot"
	_ "go-demo/sdk/gf/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
