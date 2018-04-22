package thor

import (
	"context"

	"github.com/luw2007/thor/res"
)

type JobID string
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
	ID     JobID
	Name   string
	TS     int64 // create nanosecond
	Api    string
	Action string
	Params map[string]string
	Metas  []res.Meta
	State  JobState
}
