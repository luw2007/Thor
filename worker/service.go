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
