package others

import (
	"fmt"
	"testing"
)

/**
即使空对象，长度为0，可对象依然是合法存在的，拥有合法的内存地址，这与nil的定义不符合

输出:
false
false
*/
type Student struct {
}

func TestNil(t *testing.T) {
	var stu1, stu2 Student
	// 如果把注释放开，那么结果为false，true 不清楚为啥，有知道希望能给我提个issue
	// fmt.Println(&stu1,&stu2)
	fmt.Println(&stu1 == nil)
	fmt.Println(&stu1 == &stu2)
}
