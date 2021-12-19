package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeConsuming(t *testing.T) {
	defer TimeConsuming()()

	fmt.Println("开始执行....")
	time.Sleep(1 * time.Second)
}

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	t.Log(Format(now, defaultLayout))

	t.Log(Format(now, "Y-M-D h:m:s"))

	t.Log(Format(now, "Y-M-D"))

	t.Log(Format(now, "h:m:s"))

	t.Log(FormatTime(now))

	t.Log(FormatYMD(now))

	t.Log(FormatMD(now))
}

func TestTimeParse(t *testing.T) {
	st, err := ParseTime("2020-09-10 15:22:00")
	if err != nil {
		t.Error(err)
	}
	t.Log(st)

	st, err = ParseYMD("2020-09-10")
	if err != nil {
		t.Error(err)
	}
	t.Log(st)
}

func TestTimeFirstAndLast(t *testing.T) {
	st, err := ParseTime("2020-09-10 15:22:00")
	if err != nil {
		t.Error(err)
	}
	t.Log(FirstMonth(st))
	t.Log(FirstMonthUnix(st))
	t.Log(LastMonth(st))
	t.Log(LastMonthUnix(st))
}
