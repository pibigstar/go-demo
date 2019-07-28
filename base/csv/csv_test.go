package main

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestCsv(t *testing.T) {
	file, _ := os.OpenFile("test.csv", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	w := csv.NewWriter(file)
	w.Write([]string{"123", "456", "789", "666"})
	w.Flush()
	file.Close()

	rfile, _ := os.Open("test.csv")
	r := csv.NewReader(rfile)
	strs, _ := r.Read()
	for _, str := range strs {
		t.Log(str)
	}
}
