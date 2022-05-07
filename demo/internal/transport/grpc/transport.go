package grpc

import (
	"github.com/walrusyu/gocamp007/demo/internal/transport"
	"google.golang.org/grpc/metadata"
)

var _ transport.Transporter = &Transport{}

type Transport struct {
	endpoint    string
	operation   string
	reqHeader   HeaderCarrier
	replyHeader HeaderCarrier
}

func NewTransport(endpoint string, operation string, reqHeader HeaderCarrier, replyHeader HeaderCarrier) *Transport {
	return &Transport{
		endpoint:    endpoint,
		operation:   operation,
		reqHeader:   reqHeader,
		replyHeader: replyHeader,
	}
}

func (t *Transport) Kind() string {
	return transport.KindGRPC
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

type HeaderCarrier metadata.MD

// Get returns the value associated with the passed key.
func (mc HeaderCarrier) Get(key string) string {
	vals := metadata.MD(mc).Get(key)
	if len(vals) > 0 {
		return vals[0]
	}
	return ""
}

// Set stores the key-value pair.
func (mc HeaderCarrier) Set(key string, value string) {
	metadata.MD(mc).Set(key, value)
}

// Keys lists the keys stored in this carrier.
func (mc HeaderCarrier) Keys() []string {
	keys := make([]string, 0, len(mc))
	for k := range metadata.MD(mc) {
		keys = append(keys, k)
	}
	return keys
}
