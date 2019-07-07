package goroutinue

import (
	"testing"
	"time"
)

// 当任意有一个完成之后即可返回
func TestGetOne(t *testing.T) {
	t.Log(GetOne())
}

func TestGetAll(t *testing.T) {
	t.Log(GetAll())
}

func TestGetAllWithGroup(t *testing.T) {
	t.Log(GetAllWithGroup())
}

func TestGetWithTimeout(t *testing.T) {
	// no time out
	result, err := GetWithTimeout(time.Millisecond * 50)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(result)
	// time  out
	result, err = GetWithTimeout(time.Millisecond * 150)
	if err != nil {
		t.Log(err.Error())
	}
}

func TestCancelTask(t *testing.T) {
	CancelTask()
}

func TestCancelAllTask(t *testing.T) {
	CancelAllTask()
}
