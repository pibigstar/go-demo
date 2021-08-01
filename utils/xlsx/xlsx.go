package xlsxkit

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"reflect"
	"strings"

	"github.com/xuri/excelize"
)

type Record struct {
	Name string `xlsx:"A-姓名"`
	Age  int32  `xlsx:"B-年齡"`
}

func RefactorWrite(records []*Record) error {
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
	return err
}

// 写内容到xlsx
func WriteXlsx() error {

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
	return err
}

// 读取 xlsx
func ReadXlsx(xlsxfile string) ([][]string, error) {
	xlsx, err := excelize.OpenFile(xlsxfile)
	if err != nil {
		return nil, err
	}
	cell := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println("B2:", cell)

	rows := xlsx.GetRows("Sheet1")
	return rows, nil
}

func ReadToStruct(fileName, sheetName string, result interface{}) error {
	xlsx, err := excelize.OpenFile(fileName)
	if err != nil {
		return err
	}

	rows := xlsx.GetRows(sheetName)

	var records []map[string]interface{}
	if len(rows) == 0 {
		return nil
	}

	columns := rows[0]
	columnJson := getColumnJson(result)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		record := make(map[string]interface{})
		for f, c := range row {
			column := columns[f]
			if j, ok := columnJson[column]; ok {
				record[j] = c
			}
		}
		records = append(records, record)
	}

	err = CopyStruct(records, &result)
	fmt.Println(records)
	return err
}

var json2 = func() jsoniter.API {
	// 开启模糊模式
	extra.RegisterFuzzyDecoders()
	return jsoniter.ConfigCompatibleWithStandardLibrary
}()

func CopyStruct(src interface{}, dst interface{}) error {
	r := reflect.ValueOf(src)
	if r.Kind() == reflect.Ptr {
		if !r.IsNil() {
			bs, _ := json2.Marshal(src)
			return json2.Unmarshal(bs, &dst)
		}
		return nil
	}
	bs, _ := json2.Marshal(src)
	return json2.Unmarshal(bs, &dst)
}

func getColumnJson(model interface{}) map[string]string {
	columnJson := make(map[string]string)
	d := reflect.TypeOf(model).Elem().Elem()
	for j := 0; j < d.NumField(); j++ {
		var (
			columnName string
		)
		columns := strings.Split(d.Field(j).Tag.Get("xlsx"), "-")
		if len(columns) == 2 {
			columnName = columns[1]
		}
		columnJson[columnName] = d.Field(j).Name
	}
	return columnJson
}
