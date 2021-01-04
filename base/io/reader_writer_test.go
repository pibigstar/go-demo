package io

import (
	"bytes"
	"os"
	"testing"
)

// io.Writer接口，将数据写入到某处
// io.Reader接口，从某处读取数据
func TestWriterReader(t *testing.T) {
	var b bytes.Buffer
	b.WriteString("Hello ")

	// 以只可读的方式打开
	file, err := os.Open("io.txt")
	if err != nil {
		t.Error(err)
	}

	// 从实现了io.Reader接口的对象中读取全部数据
	// ioutil.ReadAll 底层就是用该函数实现的
	_, err = b.ReadFrom(file)
	if err != nil {
		t.Error(err)
	}
	// 返回字节数组
	bs := b.Bytes()
	t.Log(string(bs))

	// 返回字符串
	s := b.String()
	t.Log(s)

	// 重置
	b.Reset()

	b.WriteString("pibigstar")
	// 以可读写的方式打开
	file, err = os.OpenFile("io.txt", os.O_RDWR, 0)
	if err != nil {
		t.Error(err)
	}
	// 将数据写入到实现了io.Writer接口的对象中
	_, err = b.WriteTo(file)
	if err != nil {
		t.Error(err)
	}
}
