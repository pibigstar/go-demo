package main

import "testing"

func TestMinWindow(t *testing.T) {
	s := "abcdefgslfjslkj"
	tt := "cfg"
	result := minWindow(s, tt)
	t.Log(result)
}
