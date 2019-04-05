package rw

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const fileName = "file/demo.text"

// 使用ioutil读取文件
func ReadFile() {

	data, err := ioutil.ReadFile(fileName)
	check(err)
	fmt.Println(string(data))
}

// 读取文件夹
func ReadAllDir() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// 这种会覆盖掉原先内容
func WriteFile() {

	data := "this new add data"

	err := ioutil.WriteFile(fileName, []byte(data), 0666)
	check(err)
}

// 追加内容到文件末尾
func AppendToFile() {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	check(err)

	file.Write([]byte("\n hello pibigstar"))

}

// 创建文件并返回文件指针
func CreateFile() {
	file, err := os.Create("test.text")
	defer file.Close()
	file.Write([]byte("hello world"))
	check(err)
}

// 创建文件夹
func MkdirFile() {
	//创建文件
	err := os.Mkdir("mkdir", 0666)
	check(err)

	// 创建多层文件夹
	err = os.MkdirAll("test/demo/user", 0666)
	check(err)
}
// 删除文件
func DeleteFile()  {
	err := os.Remove("test.text")
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
