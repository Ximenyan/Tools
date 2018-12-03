package tools

import (
	"fmt"
	"sync"
	"time"
)

var id_once sync.Once = sync.Once{}
var ch chan string

func init() {
	id_once.Do(func() {
		ch = make(chan string)
		go func() {
			for {
				ch <- fmt.Sprintf("%d", time.Now().UnixNano())
			}
		}()
	})
}

func GetId() string {
	return <-ch
}
