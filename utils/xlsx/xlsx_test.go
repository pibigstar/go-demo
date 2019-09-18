package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/xuri/excelize"
)

type Record struct {
	Name string `xlsx:"A-姓名"`
	Age  int32  `xlsx:"B-年齡"`
}

func TestXlsx(t *testing.T) {
	// 写xlsx
	WriteXlsx()
	// 读xlsx
	ReadXlsx("test_write.xlsx")

	var records []*Record
	records = append(records, &Record{
		Name: "小明",
		Age:  11,
	})
	records = append(records, &Record{
		Name: "小华",
		Age:  12,
	})
	// 反射写
	RefactorWrite(records)
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
	err := xlsx.SaveAs("test_write.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func RefactorWrite(records []*Record) {
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet1")

	for i, t := range records {
		d := reflect.TypeOf(t).Elem()
		for j := 0; j < d.NumField(); j++ {
			// 设置表头
			if i == 0 {
				column := strings.Split(d.Field(j).Tag.Get("xlsx"), "-")[0]
				name := strings.Split(d.Field(j).Tag.Get("xlsx"), "-")[1]
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+1), name)
			}
			// 设置内容
			column := strings.Split(d.Field(j).Tag.Get("xlsx"), "-")[0]
			switch d.Field(j).Type.String() {
			case "string":
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+2), reflect.ValueOf(t).Elem().Field(j).String())
			case "int32":
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+2), reflect.ValueOf(t).Elem().Field(j).Int())
			case "int64":
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+2), reflect.ValueOf(t).Elem().Field(j).Int())
			case "bool":
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+2), reflect.ValueOf(t).Elem().Field(j).Bool())
			case "float32":
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+2), reflect.ValueOf(t).Elem().Field(j).Float())
			case "float64":
				xlsx.SetCellValue("Sheet1", fmt.Sprintf("%s%d", column, i+2), reflect.ValueOf(t).Elem().Field(j).Float())
			}

		}
	}

	xlsx.SetActiveSheet(index)
	// 保存到xlsx中
	err := xlsx.SaveAs("test_write.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
