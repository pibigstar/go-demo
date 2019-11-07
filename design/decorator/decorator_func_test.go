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
