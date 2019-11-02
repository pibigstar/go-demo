package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func writeCsv(file *os.File) {
	w := csv.NewWriter(file)
	w.Write([]string{"123", "456", "789", "666"})
	w.Flush()
	file.Close()
}

func readCsv(file *os.File) {
	r := csv.NewReader(file)
	strs, _ := r.Read()
	for _, str := range strs {
		fmt.Println(str)
	}
}
