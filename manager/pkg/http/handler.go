package http

import (
	"context"
	"encoding/json"
	http1 "github.com/go-kit/kit/transport/http"
	endpoint "github.com/luw2007/thor/manager/pkg/endpoint"
	"net/http"
)

// makeRegisterHandler creates the handler logic
func makeRegisterHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/register", http1.NewServer(endpoints.RegisterEndpoint, decodeRegisterRequest, encodeRegisterResponse, options...))
}

// decodeRegisterResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeRegisterResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRegisterResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeResourceHandler creates the handler logic
func makeResourceHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/resource", http1.NewServer(endpoints.ResourceEndpoint, decodeResourceRequest, encodeResourceResponse, options...))
}

// decodeResourceResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeResourceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ResourceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeResourceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeResourceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
