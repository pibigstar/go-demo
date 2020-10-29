package elastic

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// Transport for tracing Elastic operations.
type Transport struct {
	rt    http.RoundTripper
	debug bool
}

// Option signature for specifying options, e.g. WithRoundTripper.
type Option func(t *Transport)

// WithRoundTripper specifies the http.RoundTripper to call
// next after this transport. If it is nil (default), the
// transport will use http.DefaultTransport.
func WithRoundTripper(rt http.RoundTripper) Option {
	return func(t *Transport) {
		t.rt = rt
	}
}

// WithDebug if debug is true then print es body
func WithDebug(debug bool) Option {
	return func(t *Transport) {
		t.debug = debug
	}
}

// NewTransport specifies a transport that will trace Elastic
// and report back via OpenTracing.
func NewTransport(opts ...Option) *Transport {
	t := &Transport{}
	for _, o := range opts {
		o(t)
	}
	return t
}

// RoundTrip captures the request and starts an OpenTracing span
// for Elastic PerformRequest operation.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	span, ctx := opentracing.StartSpanFromContext(req.Context(), "ElasticSearch")
	req = req.WithContext(ctx)
	defer span.Finish()

	ext.HTTPUrl.Set(span, req.URL.String())
	ext.HTTPMethod.Set(span, req.Method)
	ext.PeerHostname.Set(span, req.URL.Hostname())
	if t.debug {
		data, _ := ioutil.ReadAll(req.Body)
		span.SetTag("http.body", string(data))
		req.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	}

	var (
		resp *http.Response
		err  error
	)
	if t.rt != nil {
		resp, err = t.rt.RoundTrip(req)
	} else {
		resp, err = http.DefaultTransport.RoundTrip(req)
	}
	if err != nil {
		ext.Error.Set(span, true)
	}
	if resp != nil {
		ext.HTTPStatusCode.Set(span, uint16(resp.StatusCode))
	}

	return resp, err
}
