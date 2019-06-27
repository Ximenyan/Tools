package tools

import (
	"log"

	"github.com/shirou/gopsutil/mem"
)

//最小剩余内存
var MinFree uint64
var ChanGetFree chan uint64

func GetFree() uint64 {
	v, _ := mem.VirtualMemory()
	log.Println("free : ", v.Free/(1024*1024), " min_free : ", MinFree)
	return v.Free / (1024 * 1024)
}
