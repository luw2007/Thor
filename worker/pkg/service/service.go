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

// WorkerService 工作服务
type WorkerService interface {
	PostResource(ctx context.Context, metas []res.Meta) (reply thor.Reply)
	PostJob(ctx context.Context, job thor.Job) (reply thor.Reply)
}

type basicWorkerService struct{}

func (b *basicWorkerService) PostResource(ctx context.Context, metas []res.Meta) (reply thor.Reply) {
	// TODO implement the business logic of PostResource
	return reply
}
func (b *basicWorkerService) PostJob(ctx context.Context, job thor.Job) (reply thor.Reply) {
	// TODO implement the business logic of PostJob
	return reply
}

// NewBasicWorkerService returns a naive, stateless implementation of WorkerService.
func NewBasicWorkerService() WorkerService {
	return &basicWorkerService{}
}

// New returns a WorkerService with all of the expected middleware wired in.
func New(middleware []Middleware) WorkerService {
	var svc WorkerService = NewBasicWorkerService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
