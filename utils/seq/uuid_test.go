package seq

import (
	"testing"
)

func TestUUID(t *testing.T) {
	uuid := UUID()
	t.Log(uuid)

	shortUUID := UUIDShort()
	t.Log(shortUUID)
}
