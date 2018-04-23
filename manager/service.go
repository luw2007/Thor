package service

import (
	"context"

	"github.com/luw2007/thor"
	"github.com/luw2007/thor/res"
)

// ManagerService 管理服务
type ManagerService interface {
	Register(ctx context.Context, workerId int, addr string) (reply thor.Reply)
	Resource(ctx context.Context, t res.Type, id int) (meta res.Meta)
}
