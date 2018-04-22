package res

import "encoding/json"

type proxy struct {
	Id   int    `json:"id"`   // 代理编号
	Host string `json:"host"` // 代理 ip:port
}

func (u proxy) GetID() int {
	return u.Id
}

func (u proxy) Info() []byte {
	info, _ := json.Marshal(u)
	return info
}

func (u *proxy) Type() Type {
	return Proxy
}
