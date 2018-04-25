package service

import (
	"context"

	"errors"

	"github.com/luw2007/thor"
	"github.com/luw2007/thor/res"
)

var (
	ErrorDecoder = errors.New("decode response error")
)

// ManagerService 管理服务
type ManagerService interface {
	// worker 管理接口
	Register(ctx context.Context, workerId int, addr string) (reply thor.Reply)
	// resource 管理接口
	Resource(ctx context.Context, t res.Type, id int) (reply thor.Reply)
	ResourceDel(ctx context.Context, meta res.Meta) (reply thor.Reply)
	ResourceAdd(ctx context.Context, meta res.Meta) (reply thor.Reply)
}

type basicManagerService struct {
	Res    map[string]res.Meta
	Worker map[string]res.Meta
}

func (b *basicManagerService) Register(ctx context.Context, workerId int, addr string) (reply thor.Reply) {
	reply.Code = 0
	reply.Message = "ok"
	return reply
}

func (b *basicManagerService) Resource(ctx context.Context, t res.Type, id int) (reply thor.Reply) {
	// TODO implement the business logic of Resource
	return reply
}
func (b *basicManagerService) ResourceDel(ctx context.Context, meta res.Meta) (reply thor.Reply) {
	// TODO implement the business logic of ResourceDel
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
