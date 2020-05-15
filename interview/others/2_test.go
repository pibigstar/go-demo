package others

import (
	"fmt"
	"testing"
)

/*
因为for遍历时，变量stu指针不变，每次遍历仅进行struct值拷贝，
故m[stu.UserName]=&stu实际上是将所有值指向同一个指针
最终该指针的值为遍历的最后一个struct的值拷贝
*/
type student struct {
	Name string
	Age  int
}

func parseStudent() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	//for遍历时，变量stu指针不变
	//最终将所有值都指向同一个指针（最后一个）
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	return m
}

func TestForRange(t *testing.T) {
	students := parseStudent()
	for k, v := range students {
		fmt.Printf("key=%s,value=%+v \n", k, v)
	}
}

// 下面是解决方案
// 遍历时重新赋值
func parseStudentTrue() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	//用中间变量来过渡它
	for _, stu := range stus {
		s := stu
		m[stu.Name] = &s
	}
	return m
}

func TestParseStudentTrue(t *testing.T) {
	students := parseStudentTrue()
	for k, v := range students {
		fmt.Printf("key=%s,value=%+v \n", k, v)
	}
}
