package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	thor "github.com/luw2007/thor"
	service "github.com/luw2007/thor/worker/pkg/service"
)

// PostResourceRequest collects the request parameters for the PostResource method.
type PostResourceRequest struct {
	Metas []res.Meta `json:"metas"`
}

// PostResourceResponse collects the response parameters for the PostResource method.
type PostResourceResponse struct {
	Reply thor.Reply `json:"reply"`
}

// MakePostResourceEndpoint returns an endpoint that invokes PostResource on the service.
func MakePostResourceEndpoint(s service.WorkerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostResourceRequest)
		reply := s.PostResource(ctx, req.Metas)
		return PostResourceResponse{Reply: reply}, nil
	}
}

// PostJobRequest collects the request parameters for the PostJob method.
type PostJobRequest struct {
	Job thor.Job `json:"job"`
}

// PostJobResponse collects the response parameters for the PostJob method.
type PostJobResponse struct {
	Reply thor.Reply `json:"reply"`
}

// MakePostJobEndpoint returns an endpoint that invokes PostJob on the service.
func MakePostJobEndpoint(s service.WorkerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostJobRequest)
		reply := s.PostJob(ctx, req.Job)
		return PostJobResponse{Reply: reply}, nil
	}
}

// PostResource implements Service. Primarily useful in a client.
func (e Endpoints) PostResource(ctx context.Context, metas []res.Meta) (reply thor.Reply) {
	request := PostResourceRequest{Metas: metas}
	response, err := e.PostResourceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostResourceResponse).Reply
}

// PostJob implements Service. Primarily useful in a client.
func (e Endpoints) PostJob(ctx context.Context, job thor.Job) (reply thor.Reply) {
	request := PostJobRequest{Job: job}
	response, err := e.PostJobEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostJobResponse).Reply
}
