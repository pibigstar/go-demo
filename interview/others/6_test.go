package others

import (
	"fmt"
	"testing"
)

/**
当执行make([]int,5)时返回的是一个含义默认值(int的默认值为0)的数组:[0,0,0,0,0]。
而append函数是便是在一个数组或slice后面追加新的元素，并返回一个新的数组或slice。

输出: [0 0 0 0 0 1 2 3]
*/
func TestMakeSlice(t *testing.T) {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
