package main

import "testing"

func TestCalculate(t *testing.T) {
	result := calculate("1+5*2-1+2")
	t.Log(result)
}
