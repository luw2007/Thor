package res

import "encoding/json"

type worker struct {
	ID   int    `json:"id"`   // 代理编号
	Host string `json:"host"` // 代理 ip:port
}

func NewWorker(id int, host string) *worker {
	return &worker{
		ID:   id,
		Host: host,
	}
}

func (u worker) GetID() int {
	return u.ID
}

func (u worker) Info() []byte {
	info, _ := json.Marshal(u)
	return info
}

func (u *worker) Type() Type {
	return Proxy
}
