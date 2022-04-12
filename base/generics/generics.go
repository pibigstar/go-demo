//go:build go1.18
// +build go1.18

package main

import (
	"fmt"
)

func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v \n", v)
	}
}

func main() {
	printSlice([]int{1, 2, 3})
}
