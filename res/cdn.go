package res

import (
	"encoding/json"
	"net/url"
)

type cdn struct {
	Id     int     `json:"id"`
	Source url.URL `json:"source"` // 源地址
	Host   string  `json:"host"`   // cdn地址 ip 或者ip:port
}

func (u cdn) GetID() int {
	return u.Id
}

func (u cdn) Info() []byte {
	info, _ := json.Marshal(u)
	return info
}

func (u *cdn) Type() Type {
	return CDN
}
