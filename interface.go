package thor

import (
	"context"
)

type API interface {
	Login(ctx context.Context) bool
	Index(ctx context.Context) bool
	List(ctx context.Context) interface{}
	Detail(ctx context.Context, id int) (hit, pay bool)
	Order(ctx context.Context, id int) bool
}

type Transporter interface {
	Delay() (ok bool, delay float64)
	Post(url, ua string, params map[string]string) (int, []byte)
	Close()
}

type Resource interface {
	GetID() int
	Info() []byte
}
