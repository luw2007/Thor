package res

import (
	"encoding/json"
)

type CDN struct {
	ID     int     `json:"id"`
	Source string `json:"source"` // 源地址
	Host   string  `json:"host"`   // cdn地址 ip 或者ip:port
}

func NewCDN(id int, source string, host string) *CDN {
	return &CDN{
		ID:     id,
		Source: source,
		Host:   host,
	}
}

func (u CDN) GetID() int {
	return u.ID
}

func (u CDN) Info() []byte {
	info, _ := json.Marshal(u)
	return info
}

func (u *CDN) Type() Type {
	return ResCDN
}
