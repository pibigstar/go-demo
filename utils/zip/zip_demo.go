package main

import (
	"os"
	"log"
	"bytes"
)

// 判断是否是压缩文件
func isZip(filePath string) bool {

	file, err := os.Open(filePath)
	if err!=nil {
		log.Println("file is not exits")
	}

	buf := make([]byte,1024)

	if n, err := file.Read(buf);err != nil || n<4{
		return false
	}

	return bytes.Equal(buf,[]byte("PK\x03\x04"))
}


