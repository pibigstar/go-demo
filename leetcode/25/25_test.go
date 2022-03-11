package main

import "testing"

func TestBacktrack(t *testing.T) {
	path = nil
	result = nil
	backtrack(3, 1)
	t.Log(result)
}
