package res

import (
	"encoding/json"
	"fmt"
)

type Type int

const (
	User Type = iota
	CDN
	Proxy
)

func (p Type) String() string {
	switch p {
	case User:
		return "user"
	case CDN:
		return "cdn"
	case Proxy:
		return "proxy"
	default:
		return "unknown"
	}
}

type Meta struct {
	Type Type
	ID   int
	Info []byte
}

func (m Meta) String() string {
	return fmt.Sprintf(`{"type":"%s","id":"%s"}`, m.Type, m.ID)
}
func (m Meta) Load() interface{} {
	switch m.Type {
	case User:
		var u user
		json.Unmarshal(m.Info, &u)
		return u
	case CDN:
		var c cdn
		json.Unmarshal(m.Info, &c)
		return c
	case Proxy:
		var p proxy
		json.Unmarshal(m.Info, &p)
		return p
	default:
		return nil
	}
}
