// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/luw2007/thor/manager/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	RegisterEndpoint endpoint.Endpoint
	ResourceEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.ManagerService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		RegisterEndpoint: MakeRegisterEndpoint(s),
		ResourceEndpoint: MakeResourceEndpoint(s),
	}
	for _, m := range mdw["Register"] {
		eps.RegisterEndpoint = m(eps.RegisterEndpoint)
	}
	for _, m := range mdw["Resource"] {
		eps.ResourceEndpoint = m(eps.ResourceEndpoint)
	}
	return eps
}
