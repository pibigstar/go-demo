# Go面试题及详解

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

### 3. 下面代码输出什么?
```go
func Test3(t *testing.T) {
	i := make([]int, 5)
	i = append(i, 1, 2, 3)
	fmt.Println(i)

	j := make([]int, 0)
	j = append(j, 1, 2, 3, 4)
	fmt.Println(j)
}
```
#### 输出
> [0 0 0 0 0 1 2 3]
>
> [1 2 3 4]

#### 解析
> make如果输入值，会默认给其初始化默认值

### 4. 下面这段代码有什么错误吗？
```go
func funcMui(x,y int)(sum int,error){
    return x+y,nil
}
```
#### 解析
> 第二个返回值没有命名,在函数有多个返回值时，只要有一个返回值有命名，
>其他的也必须命名。如果有多个返回值必须加上括号()；
>如果只有一个返回值且命名也必须加上括号()。
>这里的第一个返回值有命名 sum，第二个没有命名，所以错误。

### 5. new() 与 make() 的区别
> new仅仅只初始化并返回指针，而maker不仅仅要做初始化，他需要设置一些数组的长度、容量等


### 6. 下面几段代码能否通过编译，如果能，输出什么?
```go
func main() {
	list := new([]int)
    // 编译错误
    // new([]int) 之后的 list 是一个 *[]int 类型的指针
    // 不能对指针执行 append 操作。
	list = append(list, 1)
	fmt.Println(list)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
    // 编译错误，s2需要展开
	s1 = append(s1, s2)
	fmt.Println(s1)
}
```

### 7. 下面能否通过编译?
```go
func Test7(t *testing.T) {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
    // true
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
    // 编译错误，map不能进行比较
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}
```

### 8. 通过指针变量 p 访问其成员变量 name，有哪几种方式？
A. p.name

B. (&p).name

C. (*p).name

D. p->name

#### 答案
> AC

### 9. 关于字符串连接，下面语法正确的是？
A. str := 'abc' + '123'

B. str := "abc" + "123"

C. str := '123' + "abc"

D. fmt.Sprintf("abc%d", 123)

### 答案
> BD

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

### 11. 下面赋值正确的是?
A. var x = nil

B. var x interface{} = nil

C. var x string = nil

D. var x error = nil

#### 答案
> BD

### 12. 关于channel，下面语法正确的是?
A. var ch chan int

B. ch := make(chan int)

C. <- ch

D. ch <-
#### 答案
> ABC
>
> 写 chan 时，<- 右端必须要有值

### 13. 下面代码输出什么？

```go
func hello(num ...int) {
	num[0] = 18
}

func Test13(t *testing.T) {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
}
```
A.18

B.5

C.Compilation error  
#### 答案
> A
>
> 可变函数是指针传递

### 14. 下面选择哪个？
```go
func main() {  
    a := 5
    b := 8.1
    fmt.Println(a + b)
}
```
A.13.1  

B.13

C.compilation error  

#### 答案
> C
>
> 整形与浮点形不能相加

### 15. 下面代码输出什么？
```go
func Test15(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[3:4:4]
	fmt.Println(s[0])
}
```
A.3

B.4

C.compilation error  

#### 答案
> B
>
> a[3:4] = 4  a[4:4] = 4

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

### 17. 下列哪个类型可以使用 cap()函数？
A. array

B. slice

C. map

D. channel
#### 答案
> A B D

### 18. 下面代码输出什么？
```go
func Test18(t *testing.T) {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}
```
A. nil

B. not nil

C. compilation error  
#### 答案
> A
>
> 当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil。

### 19. 下面代码输出什么？
```go
func Test19(t *testing.T) {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}
```
A. runtime panic

B. 0

C. compilation error 
#### 答案
> B
>
> 删除 map 不存在的键值对时，不会报错，相当于没有任何作用；
> 获取不存在的减值对时，返回值类型对应的零值，所以返回 0

### 20. 下面代码输出什么？
```go
func Test20(t *testing.T) {
	i := -5
	j := +5
	fmt.Printf("%+d %+d", i, j)
}
```
A. -5 +5

B. +5 +5

C. 0  0
#### 答案
> A
>
> %+d 是带符号输出

### 21. 定义一个全局字符串变量，下列正确的是？
A. var str string

B. str := ""

C. str = ""

D. var str = ""
#### 答案
> A D
>
> B 只支持局部变量声明；C 是赋值，str 必须在这之前已经声明；
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

### 24. 下面代码输出什么？
```go
func inc(p *int) int {
	*p++
	return *p
}

func Test24(t *testing.T) {
	p := 1
	inc(&p)
	fmt.Println(p)
}
```
A. 1

B. 2

C. 3
#### 答案
> B

### 25. 关于可变参数的函数调用正确的是？
```go
func add(args ...int) int {

    sum := 0
    for _, arg := range args {
        sum += arg
    }
    return sum
}
```
A. add(1, 2)

B. add(1, 3, 7)

C. add([]int{1, 2})

D. add([]int{1, 3, 7}…)

#### 答案
> A B D

### 26. 下列代码中下划线处可填入哪些变量？
```go
func Test26(t *testing.T) {
	var s1 []int
	var s2 = []int{}
	if ___ == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}
}
```
A. s1

B. s2

C. s1、s2 都可以
#### 答案
> A
>
> nil 切片和空切片。nil 切片和 nil 相等，一般用来表示一个不存在的切片；
> 空切片和 nil 不相等，表示一个空的集合

### 27. 下面代码输出什么？
```go
func Test27(t *testing.T) {
	i := 65
	fmt.Println(string(i))
}
```
A. A

B. 65

C. compilation error

#### 答案
> A
> 
> UTF-8 编码中，十进制数字 65 对应的符号是 A

### 28. 切片a,b,c的容量分别是多少？
```go
func Test28(t *testing.T) {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Println(cap(a))
	fmt.Println(cap(b))
	fmt.Println(cap(c))
}
```
#### 输出
> 3 3 2
>
> 操作符 [i:j:k]，k 主要是用来限制切片的容量，
> 但是不能大于数组的长度 ，截取得到的切片长度和容量计算方法是 j-i、k-i

### 29. 下面代码输出什么？
```go
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
```
#### 输出
> 0 1

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

### 32. 下面的两个切片声明中有什么区别？哪个更可取？
A. var a []int

B. a := []int{}
#### 答案
> A
>
> A 声明的是 nil 切片；B 声明的是长度和容量都为 0 的空切片。
> A的声明不会分配内存，优先选择

### 33. A,B，C，D那个有语法错误？
```go
type S struct {
}

func m(x interface{}) {
}

func g(x *interface{}) {
}

func Test33(t *testing.T) {
	s := S{}
	p := &s
	m(s) //A
	g(s) //B
	m(p) //C
	g(p) //D
}
```
#### 答案
> B D 会编译错误
>
> 函数参数为 interface{} 时可以接收任何类型的参数，包括用户自定义类型等，
> 即使是接收指针类型也用 interface{}，而不是使用 *interface{}
> 永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。

### 34. 下面代码输出什么？
```go
func Test34(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1)
	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1)
}
```
#### 答案
> [1 2 4]
>
> [1 2 4]
>
> golang 中切片底层的数据结构是数组。当使用 s1[1:] 获得切片 s2，和 s1 共享同一个底层数组
> 这会导致 s2[1] = 4 语句影响 s1。 
> 而 append 操作会导致底层数组扩容，生成新的数组，因此追加数据后的 s2 不会影响 s1

### 35. 下列代码输出什么？
```go
func Test35(t *testing.T) {
	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}
}
```
#### 答案
> 1 2

### 36. 下列代码输出什么？
```go
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
```
#### 答案
> 10 1 2 3
>
> 20 0 2 2
>  
> B 0 2 2
>
> A 1 3 4


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

### 38. 下面代码是否可以编译通过？
```go
type People interface {
    Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) {
    fmt.Println(think)
}

func main() {
    var peo People = Student{}
    think := "speak"
    fmt.Println(peo.Speak(think))
}
```

#### 答案
> 不能编译通过，因为是 *Student 实现了Speak，并不是 值类型的Student，
但是如果是 Student 类型实现了Speak方法，那么用 值类型的`Student{}` 或是指针类型的`&Student{}`都可以访问到该方法

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

### 39. 下面代码输出什么？
```go
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
```
#### 答案
> s is nil
>
> p is not nil
>
> 当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。上面的代码，给变量 p 赋值之后，
> p 的动态值是 nil，但是动态类型却是 *Student，是一个 nil 指针，所以相等条件不成立。

### 41. 下列代码输出什么？
```go
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
```
#### 答案
> South

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

### 43. 下面代码输出什么？
```go
var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p)
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
```
#### 答案
> bar 函数会发生panic，空指针异常
>
> 因为 err 前面没有声明，所以 p, err := foo() 中的 p 是重新声明的局部变量，而不是我们在前面声明的全局变量 p

### 44. 下面代码输出什么？
```go
func Test44(t *testing.T) {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
		fmt.Println(v)
	}
}
```

#### 答案
> [1 2 3 0]
>
> [1 2 3 0 1]
>
> [1 2 3 0 1 2]

### 45. 下面代码输出什么？
```go
func Test45(t *testing.T) {
	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 1)
}
```
#### 答案
> 2 3
>
> 2 3
>
> 2 3
> 
> for range 使用短变量声明(:=)的形式迭代变量，
> 需要注意的是，变量 i、v 在每次循环体中都会被重用，而不是重新声明。

#### 解决方案
> 有两种解决方式

第一种(推荐)
```go
for i, v := range m {
    go func(i,v int) {
        fmt.Println(i, v)
    }(i,v)
}
```
第二种
```go
for i, v := range m {
    i := i           // 这里的 := 会重新声明变量，而不是重用
    v := v
    go func() {
        fmt.Println(i, v)
    }()
}
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

### 47. 下列代码输出什么？
```go
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
```
#### 答案
> r =  [1 2 3 4 5]
>
> a =  [1 12 13 4 5]

### 48. 下面代码输出什么？
```go
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
```
#### 答案
> [1 2 0 0 0]
>
> [1 2 3 0 0]

### 49. 下列代码输出什么？
```go
func Test49(t *testing.T) {
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
```

#### 答案
> counter is 2 或者 counter is 3
>
> for range map 是无序的

### 50. 关于协程，下列说法正确的有？

A. 协程和线程都可以实现程序的并发执行；

B. 线程比协程更轻量级；

C. 协程不存在死锁问题；

D. 通过 channel 来进行协程间的通信；

#### 答案
> A D

### 51.关于循环语句，下面说法正确的有？
    
A. 循环语句既支持 for 关键字，也支持 while 和 do-while；

B. 关键字 for 的基本使用方法与 C/C++ 中没有任何差异；

C. for 循环支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break，可以选择中断哪一个循环；

D. for 循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量；

#### 答案
> C D

### 52. 下列代码输出什么？
```go
func Test52(t *testing.T) {
	i := 1
	s := []string{"A", "B", "C"}
	i, s[i-1] = 2, "Z"
	fmt.Printf("s: %v \n", s)
}
```

#### 答案
> s: [Z B C] 