# Go面试题及详解
> 每个文件包含5个左右面试题，下面是面试题汇总

## 汇总

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


