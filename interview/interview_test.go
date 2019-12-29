package interview

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	//panic("触发异常")
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

func Test34(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1)
	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1)
}

func Test35(t *testing.T) {
	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}
}

func Test36(t *testing.T) {
	a := 1
	b := 2
	defer calc("A", a, calc("10", a, b))
	a = 0
	defer calc("B", a, calc("20", a, b))
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func Test37(t *testing.T) {
	m := map[int]string{0: "zero", 1: "one"}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

const (
	a = iota
	b = iota
)
const (
	name  = "name"
	name2 = "name"
	c     = iota
	d     = iota
)

func Test39(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func Test40(t *testing.T) {
	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p People = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func Test41(t *testing.T) {
	fmt.Println(South)
}

type Square struct {
	x, y int
}

var m = map[string]Square{
	"foo": Square{2, 3},
}

func Test42(t *testing.T) {
	// error
	// m["foo"].x = 1
	square := m["foo"]
	square.x = 1
	fmt.Println(m["foo"].x)
}

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//panic, p nil
	//fmt.Println(*p)
}

func Test43(t *testing.T) {
	p, err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}

func Test44(t *testing.T) {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
		fmt.Println(v)
	}
}

func Test45(t *testing.T) {
	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Millisecond * 10)
}

func f46(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()

	defer f()
	f = func() {
		r += 2
	}
	return n + 1
}

func Test46(t *testing.T) {
	fmt.Println(f46(3))
}

func Test47(t *testing.T) {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

func change(s ...int) {
	s = append(s, 3)
}

func Test48(t *testing.T) {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)
}

func Benchmark49(t *testing.B) {
	for i := 0; i < t.N; i++ {
		var m = map[string]int{
			"A": 21,
			"B": 22,
			"C": 23,
		}
		counter := 0
		for k, v := range m {
			if counter == 0 {
				delete(m, "A")
			}
			counter++
			fmt.Println(k, v)
		}
		fmt.Println("counter is ", counter)
	}
}

func Test52(t *testing.T) {
	i := 1
	s := []string{"A", "B", "C"}
	i, s[i-1] = 2, "Z"
	fmt.Printf("s: %v \n", s)
}

type Integer int

func (a Integer) Add(b Integer) Integer {
	return a + b
}

func Test54(t *testing.T) {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b)
	fmt.Println(sum)
}

func Test59(t *testing.T) {
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	intChan <- 1
	stringChan <- "hello"
	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-stringChan:
		fmt.Printf("panic: %s", value)
	}
}

func Test62(t *testing.T) {
	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Print(v)
	}
}

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func Test64(t *testing.T) {
	var x *int = nil
	Foo(x)
}

func Test67(t *testing.T) {
	var x = []int{2: 2, 3, 0: 1}
	fmt.Println(x)
}

func incr(p *int) int {
	*p++
	return *p
}
func Test68(t *testing.T) {
	v := 1
	incr(&v)
	fmt.Println(v)
}

func Test69(t *testing.T) {
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)

	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}
		r = append(r, v)
	}
	fmt.Println(r)
}

func Test75(t *testing.T) {
	s := make([]int, 3, 9)
	fmt.Println(len(s))
	s2 := s[4:8]
	fmt.Println(len(s2))
}

func Test76(t *testing.T) {
	var x interface{}
	var y interface{} = []int{3, 5}
	_ = x == x
	_ = x == y
	// _ = y == y panic
}

func Test77(t *testing.T) {
	x := make([]int, 2, 10)
	_ = x[6:10]
	// _ = x[6:] panic
	_ = x[2:]
}

type data struct {
	sync.Mutex
}

func (d data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func Test78(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}

func Test79(t *testing.T) {
	var k = 1
	var s = []int{1, 2}
	k, s[k] = 0, 3
	fmt.Println(s[0] + s[1])
}

func Test80(t *testing.T) {
	//nil := 123
	//fmt.Println(nil)
	//var _ map[string]int = nil
}

func Test81(t *testing.T) {
	var x int8 = -128
	var y = x / -1
	fmt.Println(y)
}

func Test82(t *testing.T) {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		defer fmt.Println(recover())
		panic(1)
	}()
	defer recover()
	panic(2)
}

func printI(num ...int) {
	num[0] = 18
}

func Test87(t *testing.T) {
	i := []int{5, 6, 7}
	printI(i...)
	fmt.Println(i[0])
}

func alwaysFalse() bool {
	return false
}

func Test88(t *testing.T) {
	switch alwaysFalse(); {
	case true:
		println(true)
	case false:
		println(false)
	}
}

func Test92(t *testing.T) {
	fmt.Println(strings.TrimRight("ABBA", "BA"))
}

func Test93(t *testing.T) {
	var src, dst []int
	src = []int{1, 2, 3}
	copy(dst, src)
	fmt.Println(dst)
}
