package service

import (
	"context"

	"github.com/luw2007/thor"
	"github.com/luw2007/thor/res"
)

// ManagerService 管理服务
type ManagerService interface {
	Register(ctx context.Context, workerId int, addr string) (reply thor.Reply)
	Resource(ctx context.Context, t res.Type, id int) (reply thor.Reply)
	ResourceAdd(ctx context.Context, meta res.Meta) (reply thor.Reply)
}

type basicManagerService struct{}

func (b *basicManagerService) Register(ctx context.Context, workerId int, addr string) (reply thor.Reply) {
	// TODO implement the business logic of Register
	return reply
}
func (b *basicManagerService) Resource(ctx context.Context, t res.Type, id int) (reply thor.Reply) {
	// TODO implement the business logic of Resource
	return reply
}
func (b *basicManagerService) ResourceAdd(ctx context.Context, meta res.Meta) (reply thor.Reply) {
	// TODO implement the business logic of ResourceAdd
	return reply
}

// NewBasicManagerService returns a naive, stateless implementation of ManagerService.
func NewBasicManagerService() ManagerService {
	return &basicManagerService{}
}

// New returns a ManagerService with all of the expected middleware wired in.
func New(middleware []Middleware) ManagerService {
	var svc ManagerService = NewBasicManagerService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
