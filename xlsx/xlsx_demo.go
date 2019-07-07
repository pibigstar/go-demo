package main

import (
	"fmt"
	"github.com/xuri/excelize"
)

func main() {

	//WriteXlsx()

	ReadXlsx("xlsx/test_write.xlsx")

}

// 读取 xlsx
func ReadXlsx(xlsxfile string) {
	xlsx, err := excelize.OpenFile(xlsxfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	cell := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println("B2:", cell)

	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

// 写内容到xlsx
func WriteXlsx() {

	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet1")
	xlsx.SetCellValue("Sheet1", "A1", "姓名") // 第一行 第一列
	xlsx.SetCellValue("Sheet1", "B1", "年龄") // 第一行 第二列
	xlsx.SetCellValue("Sheet1", "A2", "狗子") // 第二行 第一列
	xlsx.SetCellValue("Sheet1", "B2", "18") // 第二行  第二列
	// Set active sheet of the workbook
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path
	err := xlsx.SaveAs("xlsx/test_write.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
