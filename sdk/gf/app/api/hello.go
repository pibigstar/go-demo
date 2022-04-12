package api

import (
	"github.com/gogf/gf/net/ghttp"
)

var Hello = helloApi{}

type helloApi struct {}

// Index is a demonstration route handler for output "Hello World!".
func (*helloApi) Index(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}
