# 精选

### 1. 下面这段代码的输出什么?
```go
func Test1(t *testing.T) {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}
```
#### 输出
>打印后 
>
>打印中 
>
>打印前 
>
>panic: 触发异常
 
#### 解析
> defer 的执行顺序是先进后出,发生panic后，会先执行defer

### 2. 下段代码输出什么?
```go
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
```
#### 输出
> key: 0, value: 3 
>
> key: 1, value: 3 
>
> key: 2, value: 3 
>
> key: 3, value: 3 
>
#### 解析:
>for range 循环的时候会创建每个元素的副本，而不是元素的引用，
>所以 m[key] = &val 取的都是变量 val 的地址，所以最后 map 中的所有元素的值都是变量 val 的地址，
>因为最后 val 被赋值为3，所有输出都是3.

### 10. 关于iota，下面代码输出什么?
```go
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
```
#### 输出
> 0 2 pi pi 5 6

### 16. 下面代码输出什么？
```go
func Test16(t *testing.T) {
	a := [2]int{5, 6}
	b := [3]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
```
A. compilation error  

B. equal  

C. not equal  

#### 答案
> A 编译错误
>
> 对于数组而言，一个数组是由数组中的值和数组的长度两部分组成的，如果两个数组长度不同，那么两个数组是属于
> 不同类型的，是不能进行比较的

### 22. 下列代码输出什么？
```go
func f(i int) {
	fmt.Println(i)
}
func Test22(t *testing.T) {
	i := 5
	defer f(i)
	i = i + 10
}
```
#### 输出
> 5
>
> f() 函数的参数在执行 defer 语句的时候会保存一份副本，
> 在实际调用 f() 函数时用，所以是 5.

### 23. 下列代码输出什么？
```go
func Test23(t *testing.T) {
	str := "hello"
	str[0] = 'x'
	fmt.Println(str)
}
```
A. hello

B. xello

C. compilation error
#### 答案
> C 编译错误
>
> Go 语言中的字符串是只读的

### 30. 函数 f1(),f2(),f3()分别返回什么？
```go
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
```
#### 输出
> 1 5 1

### 31. 下面代码输出什么？
```go
type Person struct {
	age int
}

func Test31(t *testing.T) {
	person := &Person{28}
	// 1
	defer fmt.Println(person.age)
        // 2
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)
	// 3
	defer func() {
		fmt.Println(person.age)
	}()
	
	person.age = 29
}
```
#### 输出
> 29 29 28
>
> 1.person.age 此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28；
>
> 2.defer 缓存的是结构体 Person{28} 的地址，最终 Person{28} 的 age 被重新赋值为 29，所以 defer 语句最后执行的时候，依靠缓存的地址取出的 age 便是 29，即输出 29；
>
> 3.闭包引用，输出 29；

### 37. 下列代码输出什么？
```go
func Test37(t *testing.T) {
	m := map[int]string{0: "zero", 1: "one"}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
```
#### 答案
> 0 zero
>
> 1 one
>
> 或者
>
> 1 one
>
> 0 zero
>
> map输出是无序的

### 39. 下面代码输出什么？
```go
const (
	a = iota
	b = iota
)
const (
	name = "name"
	c    = iota
	d    = iota
)

func Test39(t *testing.T) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
```
#### 答案
> 0 1 1 2
>
> iota 在 const 关键字出现时将被重置为0，const中每新增一行常量声明将使 iota 计数一次。

### 42. 下列代码是否可以编译通过？
```go
type Square struct {
	x, y int
}

var m = map[string]Square{
	"foo": Square{2, 3},
}

func Test42(t *testing.T) {
	m["foo"].x = 1
	fmt.Println(m["foo"].x)
}
```
#### 答案
> 编译失败， m["foo"].x = 4 报错
>
> 对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X，
> 但 go 中的 map 的 value 本身是不可寻址的
#### 正确写法
有两种解决方法：

第一种：
```go
square := m["foo"]
square.x = 1
```
第二种：
```go
var m = map[string]*Math{
    "foo": &Math{2, 3},
}
m["foo"].x = 1
```

### 46. 下面代码输出什么？
```go
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
```
#### 答案
> 7


### 50. 关于协程，下列说法正确的有？

A. 协程和线程都可以实现程序的并发执行；

B. 线程比协程更轻量级；

C. 协程不存在死锁问题；

D. 通过 channel 来进行协程间的通信；

#### 答案
> A D

### 53. .关于switch语句，下面说法正确的有?
        
A. 条件表达式必须为常量或者整数；

B. 单个case中，可以出现多个结果选项；

C. 需要用break来明确退出一个case；

D. 只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case；

#### 答案
> B D

### 60. 关于channel的特性，下面说法正确的是？
        
A. 给一个 nil channel 发送数据，造成永远阻塞

B. 从一个 nil channel 接收数据，造成永远阻塞

C. 给一个已经关闭的 channel 发送数据，引起 panic

D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值

#### 答案
> A B C D

### 65. 关于select机制，下面说法正确的是?
        
A. select机制用来处理异步IO问题；

B. select机制最大的一条限制就是每个case语句里必须是一个IO操作；

C. golang在语言级别支持select关键字；

D. select关键字的用法与switch语句非常类似，后面要带判断条件；

#### 答案
> A B C 

### 66. 下列代码有什么问题？
```go
func Stop(stop <-chan bool) {
	    close(stop)
}
```
#### 答案
> 有方向的 channel 不可以被关闭

