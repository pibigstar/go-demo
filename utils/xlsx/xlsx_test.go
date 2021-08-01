package xlsxkit

import (
	"testing"
)

func TestWriteXlsx(t *testing.T) {
	err := WriteXlsx()
	if err != nil {
		t.Error(err)
	}
}

func TestReadXlsx(t *testing.T) {
	rows, err := ReadXlsx("test_write.xlsx")
	if err != nil {
		t.Error(err)
	}

	for _, row := range rows {
		for _, colCell := range row {
			t.Log(colCell)
		}
	}
}

// 反射写Xlsx
func TestRefactorWrite(t *testing.T) {
	var records []*Record
	records = append(records, &Record{
		Name: "小明",
		Age:  11,
	})
	records = append(records, &Record{
		Name: "小华",
		Age:  12,
	})

	err := RefactorWrite(records)
	if err != nil {
		t.Error(err)
	}
}

func TestRefactorReadXlsx(t *testing.T) {
	var r []Record
	ReadToStruct("test_write.xlsx", "Sheet1", &r)
	t.Log(r)
}
