package file

import (
	"go-demo/utils/env"
	"testing"
)

func TestMain(m *testing.M) {
	if env.IsCI() {
		return
	}
	m.Run()
}

const fileName = "demo.text"

func TestCreateFile(t *testing.T) {
	CreateFile(fileName)
}

func TestWriteFile(t *testing.T) {
	WriteFile(fileName, "Hello,World")
}

func TestAppendToFile(t *testing.T) {
	AppendToFile(fileName, "Hello,Pibigstar")
}

func TestReadFile(t *testing.T) {
	ReadFile(fileName)
}

func TestMkdirFile(t *testing.T) {
	MkOneDir("demo")
	MkAllDir("test/user/one")
}

func TestReadAllDir(t *testing.T) {
	ReadAllDir(".")
}

func TestDeleteFile(t *testing.T) {
	DeleteFile(fileName)
}

func TestFileAbs(t *testing.T) {
	t.Log(FileAbs("file"))
}
