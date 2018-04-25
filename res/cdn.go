package res

import (
	"encoding/json"
	"net/url"
)

type cdn struct {
	ID     int     `json:"id"`
	Source url.URL `json:"source"` // 源地址
	Host   string  `json:"host"`   // cdn地址 ip 或者ip:port
}

func NewCDN(id int, source url.URL, host string) *cdn {
	return &cdn{
		ID:     id,
		Source: source,
		Host:   host,
	}
}

func (u cdn) GetID() int {
	return u.ID
}

func (u cdn) Info() []byte {
	info, _ := json.Marshal(u)
	return info
}

func (u *cdn) Type() Type {
	return CDN
}
