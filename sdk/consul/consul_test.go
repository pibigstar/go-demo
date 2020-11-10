package consul

import "testing"

func TestRegister(t *testing.T) {
	Register("10.54.212.69")
}

func TestReadConfig(t *testing.T) {
	result := ReadConfig("test")
	t.Log(result)
}

func TestDeleteService(t *testing.T) {
	err := DeleteService("test")
	if err != nil {
		t.Error(err)
	}
}

func TestListService(t *testing.T) {
	err := ListService()
	if err != nil {
		t.Error(err)
	}
}
