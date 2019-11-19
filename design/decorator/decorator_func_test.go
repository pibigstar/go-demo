package decorator

import (
	"go-demo/utils/env"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	if env.IsCI() {
		return
	}

	http.HandleFunc("/", Auth(f))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		t.Error(err)
	}
}

func TestReflectDecorator(t *testing.T) {
	type MyFoo func(int, int, int) int
	var myfoo MyFoo
	Decorator(&myfoo, foo)
	myfoo(1, 2, 3)

	mybar := bar
	Decorator(&mybar, bar)
	mybar("hello", "world!")
}
