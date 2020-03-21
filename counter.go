package types

import "sync/atomic"

type Counter struct {
	count int64
}

func (c *Counter) Next() int64 {
	atomic.AddInt64(&c.count, 1)
	return c.count
}
