package thor

import (
	"context"

	"encoding/json"

	"github.com/luw2007/thor/res"
)

type JobState int

const (
	Wait JobState = iota
	Ready
	Send
	Success
	Fail
)

type JobFunc func(ctx context.Context, params map[string]string, metas ...res.Meta) interface{}

type Job struct {
	ID     string            `json:"id"`
	Name   string            `json:"name"`
	TS     int64             `json:"ts"` // create nanosecond
	Api    string            `json:"api"`
	Action string            `json:"action"`
	Params map[string]string `json:"params"`
	Metas  []res.Meta        `json:"metas"`
	State  JobState          `json:"state"`
}

type Reply struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Info    json.RawMessage `json:"Info"`
	Delay   int             `json:"delay"`
}
