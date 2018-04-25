package res

import "encoding/json"

type Worker struct {
	ID   int    `json:"id"`   // 代理编号
	Host string `json:"host"` // 代理 ip:port
}

func NewWorker(id int, host string) *Worker {
	return &Worker{
		ID:   id,
		Host: host,
	}
}

func (u Worker) GetID() int {
	return u.ID
}

func (u Worker) Info() []byte {
	info, _ := json.Marshal(u)
	return info
}

func (u *Worker) Type() Type {
	return ResProxy
}
