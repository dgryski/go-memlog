// Package memlog implements a light-weight in-memory logger
/*
   http://preshing.com/20120522/lightweight-in-memory-logging/
   http://www.exampler.com/writing/ring-buffer.pdf
*/
package memlog

import (
	"sync/atomic"
	"unsafe"
)

type Event struct {
	ID    int64
	Msg   *string
	Param int64
}

const BufferSize = 1 << 16

var (
	Events [BufferSize]Event
	Pos    int64 = -1
)

func Log(id int64, msg string, param int64) {
	index := atomic.AddInt64(&Pos, 1)
	e := &Events[index&(BufferSize-1)]
	atomic.StoreInt64(&e.ID, id)
	mptr := unsafe.Pointer(&e.Msg)
	atomic.StorePointer((*unsafe.Pointer)(mptr), unsafe.Pointer(&msg))
	atomic.StoreInt64(&e.Param, param)
}
