// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	cfg       config
	requests  syncint64.Counter
	errors    syncint64.Counter
	duration  syncint64.Histogram
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...Option) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c := &Client{
		cfg:       newConfig(opts...),
		serverURL: u,
	}
	if c.requests, err = c.cfg.Meter.SyncInt64().Counter(otelogen.ClientRequestCount); err != nil {
		return nil, err
	}
	if c.errors, err = c.cfg.Meter.SyncInt64().Counter(otelogen.ClientErrorsCount); err != nil {
		return nil, err
	}
	if c.duration, err = c.cfg.Meter.SyncInt64().Histogram(otelogen.ClientDuration); err != nil {
		return nil, err
	}
	return c, nil
}

// AllRequestBodies invokes allRequestBodies operation.
//
// POST /allRequestBodies
func (c *Client) AllRequestBodies(ctx context.Context, request AllRequestBodiesReq) (res AllRequestBodiesOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("allRequestBodies"),
	}
	// Validate request before sending.
	switch request := request.(type) {
	case *AllRequestBodiesApplicationJSON:
		// Validation is not required for this type.
	case *AllRequestBodiesReqApplicationOctetStream:
		// Validation is not required for this type.
	case *AllRequestBodiesApplicationXWwwFormUrlencoded:
		// Validation is not required for this type.
	case *AllRequestBodiesMultipartFormData:
		// Validation is not required for this type.
	case *AllRequestBodiesReqTextPlain:
		// Validation is not required for this type.
	default:
		return res, errors.Errorf("unexpected request type: %T", request)
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AllRequestBodies",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.serverURL)
	u.Path += "/allRequestBodies"

	stage = "EncodeRequest"
	r := ht.NewRequest(ctx, "POST", u, nil)
	if err := encodeAllRequestBodiesRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAllRequestBodiesResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AllRequestBodiesOptional invokes allRequestBodiesOptional operation.
//
// POST /allRequestBodiesOptional
func (c *Client) AllRequestBodiesOptional(ctx context.Context, request AllRequestBodiesOptionalReq) (res AllRequestBodiesOptionalOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("allRequestBodiesOptional"),
	}
	// Validate request before sending.
	switch request := request.(type) {
	case *AllRequestBodiesOptionalApplicationJSON:
		// Validation is not required for this type.
	case *AllRequestBodiesOptionalReqApplicationOctetStream:
		// Validation is not required for this type.
	case *AllRequestBodiesOptionalApplicationXWwwFormUrlencoded:
		// Validation is not required for this type.
	case *AllRequestBodiesOptionalMultipartFormData:
		// Validation is not required for this type.
	case *AllRequestBodiesOptionalReqTextPlain:
		// Validation is not required for this type.
	default:
		return res, errors.Errorf("unexpected request type: %T", request)
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AllRequestBodiesOptional",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.serverURL)
	u.Path += "/allRequestBodiesOptional"

	stage = "EncodeRequest"
	r := ht.NewRequest(ctx, "POST", u, nil)
	if err := encodeAllRequestBodiesOptionalRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAllRequestBodiesOptionalResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}