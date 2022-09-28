package tl

import (
	"sort"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/dgrr/tl"
)

func consume(rb *tl.Ring[func()], sum *int32, max int32) {
	for atomic.LoadInt32(sum) < max {
		fn, _ := rb.Pop()
		if fn == nil {
			continue
		}

		fn()

		atomic.AddInt32(sum, 1)
	}
}

func TestRingBuffer(t *testing.T) {
	var wg sync.WaitGroup
	rb := tl.NewRing[func()](8)

	max := 1000000
	sum := int32(0)

	for i := 0; i < 256; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			consume(rb, &sum, int32(max))
		}()
	}

	var execs []int
	var lck sync.Mutex

	n := max / 8

	for i := 0; i < 8; i++ {
		go func(start, end int) {
			for i := start; i < end; i++ {
				if !rb.Push(justAppend(&lck, &execs, i)) {
					i--
				}
			}
		}(i*n, i*n+n)
	}

	wg.Wait()

	sort.Ints(execs)

	get := func(i int) []int {
		min := i - 10
		max := i + 10

		if min < 0 {
			min = 0
		}

		if i >= len(execs) {
			max = len(execs)
		}

		return execs[min:max]
	}

	for i := 0; i < max; i++ {
		if execs[i] != i {
			t.Fatalf("Missing %d, got %d:\n%v", i, execs[i], get(i))
		}
	}
}

func justAppend(lck *sync.Mutex, execs *[]int, i int) func() {
	return func() {
		lck.Lock()
		*execs = append(*execs, i)
		lck.Unlock()
	}
}
