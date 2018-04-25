package res

import "encoding/json"

type Proxy struct {
	ID   int    `json:"id"`   // 代理编号
	Host string `json:"host"` // 代理 ip:port
}

func NewProxy(id int, host string) *Proxy {
	return &Proxy{
		ID:   id,
		Host: host,
	}
}
func (u Proxy) GetID() int {
	return u.ID
}

func (u Proxy) Info() []byte {
	info, _ := json.Marshal(u)
	return info
}

func (u *Proxy) Type() Type {
	return ResProxy
}
