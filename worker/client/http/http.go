package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	http1 "net/http"
	"net/url"
	"strings"

	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	endpoint1 "github.com/luw2007/thor/worker/pkg/endpoint"
	http2 "github.com/luw2007/thor/worker/pkg/http"
	service "github.com/luw2007/thor/worker/pkg/service"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options map[string][]http.ClientOption) (service.WorkerService, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	var postResourceEndpoint endpoint.Endpoint
	{
		postResourceEndpoint = http.NewClient("POST", copyURL(u, "/post-resource"), encodeHTTPGenericRequest, decodePostResourceResponse, options["PostResource"]...).Endpoint()
	}

	var postJobEndpoint endpoint.Endpoint
	{
		postJobEndpoint = http.NewClient("POST", copyURL(u, "/post-job"), encodeHTTPGenericRequest, decodePostJobResponse, options["PostJob"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		PostJobEndpoint:      postJobEndpoint,
		PostResourceEndpoint: postResourceEndpoint,
	}, nil
}

// EncodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// SON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http1.Request, request interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodePostResourceResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodePostResourceResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.PostResourceResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodePostJobResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodePostJobResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.PostJobResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
