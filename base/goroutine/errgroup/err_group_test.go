package errgroup

import "testing"

func TestErrGroup(t *testing.T) {
	work := []int{1, 2, 3}
	err := Run(work)
	if err != nil {
		t.Log(err)
	}
}
