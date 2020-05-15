package middleware

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/zipkin"
)

func NewTracer() opentracing.Tracer {
	// 判断是否已注册了，单例模式
	if opentracing.IsGlobalTracerRegistered() {
		return opentracing.GlobalTracer()
	}
	cfg := &config.Configuration{
		ServiceName: "pibigstar",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: "127.0.0.1:6831",
			LogSpans:           true,
		},
		Headers: &jaeger.HeadersConfig{
			JaegerDebugHeader:        "x-debug-id",
			JaegerBaggageHeader:      "x-baggage",
			TraceContextHeaderName:   "x-trace-id",
			TraceBaggageHeaderPrefix: "x-ctx",
		},
	}

	propagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	tracer, _, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
		config.Injector(opentracing.HTTPHeaders, propagator),
		config.Extractor(opentracing.HTTPHeaders, propagator),
		config.ZipkinSharedRPCSpan(true),
		config.MaxTagValueLength(256),
		config.PoolSpans(true),
	)
	if err != nil {
		panic(fmt.Sprintf("Init failed: %v\n", err))
	}
	// 设置到全局里
	opentracing.SetGlobalTracer(tracer)

	return tracer
}
