package interview

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

func Test2(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Printf("key: %d, value: %d \n", k, *v)
	}
}

func Test3(t *testing.T) {
	i := make([]int, 5)
	i = append(i, 1, 2, 3)
	fmt.Println(i)

	j := make([]int, 0)
	j = append(j, 1, 2, 3, 4)
	fmt.Println(j)
}

func Test10(t *testing.T) {
	const (
		x = iota
		_
		y
		z = "pi"
		k
		p = iota
		q
	)
	fmt.Println(x, y, z, k, p, q)
}

func hello(num ...int) {
	num[0] = 18
}

func Test13(t *testing.T) {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
}

func Test15(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[3:4:4]
	fmt.Println(s[0])
}

func Test16(t *testing.T) {
	a := [3]int{5, 6}
	b := [3]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

func Test18(t *testing.T) {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

func Test19(t *testing.T) {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}

func Test20(t *testing.T) {
	i := -5
	j := +5
	fmt.Printf("%+d %+d", i, j)
}

func f(i int) {
	fmt.Println(i)
}
func Test22(t *testing.T) {
	i := 5
	defer f(i)
	i = i + 10
}

func Test23(t *testing.T) {
	str := "hello"
	// 编译错误
	// str[0] = 'x'
	fmt.Println(str)
}

func inc(p *int) int {
	*p++
	return *p
}

func Test24(t *testing.T) {
	p := 1
	inc(&p)
	fmt.Println(p)
}

func Test27(t *testing.T) {
	i := 65
	fmt.Println(string(i))
}

func Test28(t *testing.T) {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Println(cap(a))
	fmt.Println(cap(b))
	fmt.Println(cap(c))
}
func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func Test29(t *testing.T) {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func Test30(t *testing.T) {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}

type Person struct {
	age int
}

func Test31(t *testing.T) {
	person := &Person{28}

	defer fmt.Println(person.age)

	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}
