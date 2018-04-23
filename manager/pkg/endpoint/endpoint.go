package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	thor "github.com/luw2007/thor"
	service "github.com/luw2007/thor/manager/pkg/service"
	res "github.com/luw2007/thor/res"
)

// RegisterRequest collects the request parameters for the Register method.
type RegisterRequest struct {
	WorkerId int    `json:"worker_id"`
	Addr     string `json:"addr"`
}

// RegisterResponse collects the response parameters for the Register method.
type RegisterResponse struct {
	Reply thor.Reply `json:"reply"`
}

// MakeRegisterEndpoint returns an endpoint that invokes Register on the service.
func MakeRegisterEndpoint(s service.ManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		reply := s.Register(ctx, req.WorkerId, req.Addr)
		return RegisterResponse{Reply: reply}, nil
	}
}

// ResourceRequest collects the request parameters for the Resource method.
type ResourceRequest struct {
	T  res.Type `json:"t"`
	Id int      `json:"id"`
}

// ResourceResponse collects the response parameters for the Resource method.
type ResourceResponse struct {
	Reply thor.Reply `json:"reply"`
}

// MakeResourceEndpoint returns an endpoint that invokes Resource on the service.
func MakeResourceEndpoint(s service.ManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ResourceRequest)
		reply := s.Resource(ctx, req.T, req.Id)
		return ResourceResponse{Reply: reply}, nil
	}
}

// ResourceAddRequest collects the request parameters for the ResourceAdd method.
type ResourceAddRequest struct {
	Meta res.Meta `json:"meta"`
}

// ResourceAddResponse collects the response parameters for the ResourceAdd method.
type ResourceAddResponse struct {
	Reply thor.Reply `json:"reply"`
}

// MakeResourceAddEndpoint returns an endpoint that invokes ResourceAdd on the service.
func MakeResourceAddEndpoint(s service.ManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ResourceAddRequest)
		reply := s.ResourceAdd(ctx, req.Meta)
		return ResourceAddResponse{Reply: reply}, nil
	}
}

// Register implements Service. Primarily useful in a client.
func (e Endpoints) Register(ctx context.Context, workerId int, addr string) (reply thor.Reply) {
	request := RegisterRequest{
		Addr:     addr,
		WorkerId: workerId,
	}
	response, err := e.RegisterEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RegisterResponse).Reply
}

// Resource implements Service. Primarily useful in a client.
func (e Endpoints) Resource(ctx context.Context, t res.Type, id int) (reply thor.Reply) {
	request := ResourceRequest{
		Id: id,
		T:  t,
	}
	response, err := e.ResourceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ResourceResponse).Reply
}

// ResourceAdd implements Service. Primarily useful in a client.
func (e Endpoints) ResourceAdd(ctx context.Context, meta res.Meta) (reply thor.Reply) {
	request := ResourceAddRequest{Meta: meta}
	response, err := e.ResourceAddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ResourceAddResponse).Reply
}
