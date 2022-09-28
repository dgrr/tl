package iter

import (
	"math/rand"
	"testing"
	"time"

	"github.com/dgrr/tl"
)

func TestSum(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	data := func() (r []int) {
		for i := 0; i < 1_000; i++ {
			r = append(r, rand.Intn(10_000))
		}

		return r
	}()

	res := Get(
		Sum(
			Slice(data),
		),
	)

	r := tl.Sum(data...)

	if res != r {
		t.Fatalf("differs: %v <> %v", res, r)
	}
}
