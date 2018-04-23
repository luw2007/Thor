package http

import (
	"context"
	"encoding/json"
	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "github.com/luw2007/thor/worker/pkg/endpoint"
	http1 "net/http"
)

// makePostResourceHandler creates the handler logic
func makePostResourceHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/post-resource").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.PostResourceEndpoint, decodePostResourceRequest, encodePostResourceResponse, options...)))
}

// decodePostResourceResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostResourceRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.PostResourceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePostResourceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostResourceResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePostJobHandler creates the handler logic
func makePostJobHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/post-job").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.PostJobEndpoint, decodePostJobRequest, encodePostJobResponse, options...)))
}

// decodePostJobResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostJobRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.PostJobRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePostJobResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostJobResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
