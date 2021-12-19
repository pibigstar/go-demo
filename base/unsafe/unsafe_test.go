package unsafe

import (
	"reflect"
	"testing"
	"unsafe"
)

// unsafe.Pointer是一种特殊意义的指针，它可以包含任意类型的地址
// 任何指针都可以转换为unsafe.Pointer
// unsafe.Pointer可以转换为任何指针
func TestPointer(t *testing.T) {
	i := 10  // 声明一个地址和数据
	ip := &i // 取得这个数据的地址
	// 将 *int 转为 *float64
	var fp = (*float64)(unsafe.Pointer(ip)) // 把地址值给fp
	t.Logf("%p", &i)
	*fp = *fp * 3    // 把地址值的数据乘以3
	t.Logf("%p", fp) // float64
}

// uintptr可以转换为unsafe.Pointer
// unsafe.Pointer可以转换为uintptr
func TestUintptr(t *testing.T) {
	u := &user{}

	// 因为name是第一个字段，所以不用偏移，
	// 我们获取user的指针，然后通过unsafe.Pointer转为*string进行赋值操作即可
	pName := (*string)(unsafe.Pointer(u))
	*pName = "派大星"

	// 获取age的指针地址，需要偏移
	// 内存偏移牵涉到的计算只能通过uintptr
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	*pAge = 20

	t.Log(*u)
}

type user struct {
	name string
	age  int
}

// Sizeof函数可以返回一个类型所占用的内存大小，这个大小只有类型有关。
// 单位为字节 = 8bit
func TestSizeOf(t *testing.T) {
	t.Log(unsafe.Sizeof(true))
	t.Log(unsafe.Sizeof(int8(0)))
	t.Log(unsafe.Sizeof(int16(10)))
	t.Log(unsafe.Sizeof(int32(10000000)))
	t.Log(unsafe.Sizeof(int64(10000000000000)))
	t.Log(unsafe.Sizeof(int(10000000000000)))
}

// Alignof返回一个类型的对齐值，也可以叫做对齐系数或者对齐倍数。
// 对齐值是一个和内存对齐有关的值，合理的内存对齐可以提高内存读写的性能
func TestAlignOf(t *testing.T) {
	var (
		b   bool
		i8  int8
		i16 int16
		i64 int64
		f32 float32
		s   string
		m   map[string]string
		p   *int32
	)
	t.Log(unsafe.Alignof(b))
	t.Log(unsafe.Alignof(i8))
	t.Log(unsafe.Alignof(i16))
	t.Log(unsafe.Alignof(i64))
	t.Log(unsafe.Alignof(f32))
	t.Log(unsafe.Alignof(s))
	t.Log(unsafe.Alignof(m))
	t.Log(unsafe.Alignof(p))
}

// Offsetof函数只适用于struct结构体中的字段相对于结构体的内存位置偏移量。
// 结构体的第一个字段的偏移量都是0.
// 根据字段的偏移量，我们可以定位结构体的字段，进而可以读写该结构体的字段，哪怕他们是私有的
func TestOffSetOf(t *testing.T) {
	var u user
	t.Log(unsafe.Offsetof(u.name))
	t.Log(unsafe.Offsetof(u.age))
}

// 结构体字段的顺序会影响内存对齐，合理的字段顺序可以减少内存的开销
func TestStructSize(t *testing.T) {
	var u1 user1
	var u2 user2
	t.Log(unsafe.Sizeof(u1))
	t.Log(unsafe.Sizeof(u2))
}

type user1 struct {
	a byte
	b int32
	c int64
}

type user2 struct {
	a byte
	c int64
	b int32
}

// 通过构造 slice header 和 string header，来完成 string 和 byte slice 之间的转换
func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

func TestSlice2String(t *testing.T) {
	bs := string2bytes("派大星")
	t.Log(string(bs))

	s := bytes2string(bs)
	t.Log(s)
}
