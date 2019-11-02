package zip

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
)

// 判断是否是压缩文件
func isZip(filePath string) bool {

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("file is not exits")
	}

	buf := make([]byte, 1024)

	if n, err := file.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

// 解压zip文件
func unzip(zipfile, target string) error {

	reader, err := zip.OpenReader(zipfile)
	if err != nil {
		return err
	}
	defer reader.Close()

	if err := os.MkdirAll(target, 0765); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targerReader, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targerReader.Close()

		if _, err := io.Copy(targerReader, fileReader); err != nil {
			return err
		}
	}

	return nil
}

type zipFile struct {
	Name string
	Body string
}

// 压缩zip
func zipfile(fileName string, files []zipFile) {
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	buf.WriteTo(f)
}
