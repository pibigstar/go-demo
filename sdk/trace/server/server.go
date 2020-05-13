package main

import (
	"go-demo/sdk/trace"
	"io"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

var (
	tracer opentracing.Tracer
	closer io.Closer
)

func init() {
	tracer, closer = trace.NewTracer("hello-server", "const", 1)
}

func hello(w http.ResponseWriter, req *http.Request) {
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	span := tracer.StartSpan("hello", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	io.WriteString(w, "Hello World!")
}

func main() {
	defer closer.Close()

	http.HandleFunc("/hello", hello)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
