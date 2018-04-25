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
	endpoint1 "github.com/luw2007/thor/manager/pkg/endpoint"
	http2 "github.com/luw2007/thor/manager/pkg/http"
	service "github.com/luw2007/thor/manager/pkg/service"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options map[string][]http.ClientOption) (service.ManagerService, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	var registerEndpoint endpoint.Endpoint
	{
		registerEndpoint = http.NewClient("POST", copyURL(u, "/register"), encodeHTTPGenericRequest, decodeRegisterResponse, options["Register"]...).Endpoint()
	}

	var resourceEndpoint endpoint.Endpoint
	{
		resourceEndpoint = http.NewClient("POST", copyURL(u, "/resource"), encodeHTTPGenericRequest, decodeResourceResponse, options["Resource"]...).Endpoint()
	}

	var resourceDelEndpoint endpoint.Endpoint
	{
		resourceDelEndpoint = http.NewClient("POST", copyURL(u, "/resource-del"), encodeHTTPGenericRequest, decodeResourceDelResponse, options["ResourceDel"]...).Endpoint()
	}

	var resourceAddEndpoint endpoint.Endpoint
	{
		resourceAddEndpoint = http.NewClient("POST", copyURL(u, "/resource-add"), encodeHTTPGenericRequest, decodeResourceAddResponse, options["ResourceAdd"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		RegisterEndpoint:    registerEndpoint,
		ResourceAddEndpoint: resourceAddEndpoint,
		ResourceDelEndpoint: resourceDelEndpoint,
		ResourceEndpoint:    resourceEndpoint,
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

// decodeRegisterResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeRegisterResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.RegisterResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeResourceResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeResourceResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.ResourceResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeResourceDelResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeResourceDelResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.ResourceDelResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeResourceAddResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeResourceAddResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.ResourceAddResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
