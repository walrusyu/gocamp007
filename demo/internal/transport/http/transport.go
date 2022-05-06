package http

import (
	"github.com/walrusyu/gocamp007/demo/internal/transport"
	"net/http"
)

var _ Transporter = &Transport{}

type Transporter interface {
	transport.Transporter
	Request() *http.Request
	PathTemplate() string
}

type Transport struct {
	endpoint     string
	operation    string
	reqHeader    headerCarrier
	replyHeader  headerCarrier
	request      *http.Request
	pathTemplate string
}

func (t *Transport) Kind() string {
	return transport.KindHTTP
}

func (t *Transport) Endpoint() string {
	return t.endpoint
}

func (t *Transport) Operation() string {
	return t.operation
}

func (t *Transport) RequestHeader() transport.Header {
	return t.reqHeader
}

func (t *Transport) ReplyHeader() transport.Header {
	return t.replyHeader
}

// Request returns the HTTP request.
func (tr *Transport) Request() *http.Request {
	return tr.request
}

// PathTemplate returns the http path template.
func (tr *Transport) PathTemplate() string {
	return tr.pathTemplate
}

type headerCarrier http.Header

// Get returns the value associated with the passed key.
func (hc headerCarrier) Get(key string) string {
	return http.Header(hc).Get(key)
}

// Set stores the key-value pair.
func (hc headerCarrier) Set(key string, value string) {
	http.Header(hc).Set(key, value)
}

// Keys lists the keys stored in this carrier.
func (hc headerCarrier) Keys() []string {
	keys := make([]string, 0, len(hc))
	for k := range http.Header(hc) {
		keys = append(keys, k)
	}
	return keys
}
