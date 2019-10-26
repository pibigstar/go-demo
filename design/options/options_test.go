package options

import (
	"testing"
)

func TestConnect(t *testing.T) {
	connection, err := Connect("127.0.0.1", WithCaching(false), WithTimeout(20))
	if err != nil {
		panic(err)
	}
	t.Logf("%+v", connection)
}
