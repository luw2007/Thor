package res

import (
	"encoding/json"
	"fmt"
)

type Type int

const (
	ResUser Type = iota
	ResCDN
	ResProxy
	ResWorker
)

func (p Type) String() string {
	switch p {
	case ResUser:
		return "user"
	case ResCDN:
		return "cdn"
	case ResProxy:
		return "proxy"
	case ResWorker:
		return "worker"
	default:
		return "unknown"
	}
}

type Meta struct {
	Type Type
	ID   int
	Info []byte
}

func NewMeta(id int, t Type, info []byte) Meta {
	return Meta{
		Type: t,
		ID:   id,
		Info: info,
	}
}

func (m Meta) Key() string {
	return MetaKey(m.Type, m.ID)
}

func MetaKey(t Type, id int) string {
	return fmt.Sprintf("%s:%d", t, id)
}

func (m Meta) String() string {
	return fmt.Sprintf(`{"type":"%s","id":"%s"}`, m.Type, m.ID)
}

func (m Meta) Load() interface{} {
	switch m.Type {
	case ResUser:
		var u User
		json.Unmarshal(m.Info, &u)
		return u
	case ResCDN:
		var c CDN
		json.Unmarshal(m.Info, &c)
		return c
	case ResProxy:
		var p Proxy
		json.Unmarshal(m.Info, &p)
		return p
	case ResWorker:
		var p Worker
		json.Unmarshal(m.Info, &p)
		return p
	default:
		return nil
	}
}
