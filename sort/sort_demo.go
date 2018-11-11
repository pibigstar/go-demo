package main

import (
	"log"
	"sort"
)

func main() {

	ids := []int{1, 5, 6, 11, 19, 2, 7}

	//递增排序
	//sort.Ints(ids)
	//sort.Sort(sort.IntSlice(ids))

	//递减排序
	reverse := sort.Reverse(sort.IntSlice(ids))
	sort.Sort(reverse)

	for _, id := range ids {
		log.Println(id)
	}

}
