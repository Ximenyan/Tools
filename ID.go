package tools

import (
	"sync"
)

var ID IDCreater

type IDCreater struct {
	ch chan uint64
}

var init_once = sync.Once{}

func (this IDCreater) GetUint64() uint64 {
	return <-this.ch
}

func init() {
	init_once.Do(func() {
		ch := make(chan uint64)
		ID = IDCreater{ch}
		go func() {
			i := uint64(0)
			for {
				ch <- i
				i++
			}
		}()
	})
}
