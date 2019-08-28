package cmp

import "testing"

type User struct {
	Name     string
	Password string
}

func TestDiff(t *testing.T) {
	u1 := &User{
		Name:     "派大星",
		Password: "123456",
	}

	u2 := &User{
		Name:     "海绵宝宝",
		Password: "123456",
	}

	diff := Diff(u1, u2)
	t.Log(diff)
}
