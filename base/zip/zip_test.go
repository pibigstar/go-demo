package zip

import "testing"

func TestZipFile(t *testing.T) {
	var files = []zipFile{
		{"1.txt", "first"},
		{"2.txt", "second"},
		{"3.txt", "third"},
	}
	zipfile("file.zip", files)

	unzip("file.zip", "test")

}
