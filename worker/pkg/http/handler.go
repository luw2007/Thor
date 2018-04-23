package http

import (
	"context"
	"encoding/json"
	http1 "github.com/go-kit/kit/transport/http"
	endpoint "github.com/luw2007/thor/worker/pkg/endpoint"
	"net/http"
)

// makePostResourceHandler creates the handler logic
func makePostResourceHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/post-resource", http1.NewServer(endpoints.PostResourceEndpoint, decodePostResourceRequest, encodePostResourceResponse, options...))
}

// decodePostResourceResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostResourceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.PostResourceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePostResourceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostResourceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePostJobHandler creates the handler logic
func makePostJobHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/post-job", http1.NewServer(endpoints.PostJobEndpoint, decodePostJobRequest, encodePostJobResponse, options...))
}

// decodePostJobResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostJobRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.PostJobRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePostJobResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostJobResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
