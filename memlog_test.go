package memlog

import (
	"sync"
	"testing"
)

func TestRace(t *testing.T) {
	var wg sync.WaitGroup

	var msgs = [4]string{"foo", "bar", "baz", "qux"}

	for id := int64(0); id < 4; id++ {
		wg.Add(1)
		go func(id int64) {
			for i := int64(0); i < 1e6; i++ {
				Log(id, msgs[id], i)
			}
			wg.Done()
		}(id)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		e := &Events[i]
		t.Log(e.ID, *e.Msg, e.Param)
	}
}
