package counter

import (
	"fmt"
	"time"
)

type counter struct {
	cnt        []int
	size       int
	lastUpdate int64
}

func CreateCounter(size int) counter {
	return counter{
		cnt:  make([]int, size),
		size: size,
	}
}

func (c *counter) Add() {
	now := time.Now().Unix()
	var offset int
	if now-c.lastUpdate > int64(c.size) {
		offset = c.size
	} else {
		offset = int(now - c.lastUpdate)
	}
	c.lastUpdate = now
	for i := 0; i < c.size-offset; i++ {
		c.cnt[i] = c.cnt[i+offset]
	}
	for i := max(0, c.size-offset); i < c.size; i++ {
		c.cnt[i] = 0
	}
	c.cnt[c.size-1]++
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (c *counter) ToString() string {
	return fmt.Sprintf("%v", c.cnt)
}
