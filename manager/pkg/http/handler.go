package http

import (
	"context"
	"encoding/json"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "github.com/luw2007/thor/manager/pkg/endpoint"
)

// makeRegisterHandler creates the handler logic
func makeRegisterHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/register").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.RegisterEndpoint, decodeRegisterRequest, encodeRegisterResponse, options...)))
}

// decodeRegisterResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeRegisterRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeRegisterResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeRegisterResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeResourceHandler creates the handler logic
func makeResourceHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/resource").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ResourceEndpoint, decodeResourceRequest, encodeResourceResponse, options...)))
}

// decodeResourceResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeResourceRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ResourceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeResourceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeResourceResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeResourceDelHandler creates the handler logic
func makeResourceDelHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/resource-del").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ResourceDelEndpoint, decodeResourceDelRequest, encodeResourceDelResponse, options...)))
}

// decodeResourceDelResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeResourceDelRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ResourceDelRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeResourceDelResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeResourceDelResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeResourceAddHandler creates the handler logic
func makeResourceAddHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/resource-add").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.ResourceAddEndpoint, decodeResourceAddRequest, encodeResourceAddResponse, options...)))
}

// decodeResourceAddResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeResourceAddRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.ResourceAddRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeResourceAddResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeResourceAddResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
