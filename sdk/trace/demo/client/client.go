package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"go-demo/sdk/trace/demo"
	"io/ioutil"
	"net/http"
)

const URL = "http://localhost:8081/hello"

func main() {
	tracer, closer := demo.NewTracer("hello-client")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// 设置请求名
	span := tracer.StartSpan("hello trace")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	span, _ = opentracing.StartSpanFromContext(ctx, "client request")
	defer span.Finish()

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		panic(err)
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, URL)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	fmt.Println(req.URL.String())
	span.LogFields(
		log.String("event", "hello"),
		log.String("value", "hello"),
	)

	reqPrepareSpan, _ := opentracing.StartSpanFromContext(ctx, "Client_sendRequest")
	defer reqPrepareSpan.Finish()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Do send requst failed(%s)\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll error(%s)\n", err)
		return
	}

	fmt.Printf("Response:%s\n", string(body))

	respSpan, _ := opentracing.StartSpanFromContext(ctx, "Client_response")
	defer respSpan.Finish()
	respSpan.LogFields(
		log.String("event", "getResponse"),
		log.String("value", string(body)),
	)
}
