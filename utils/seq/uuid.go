package seq

import (
	"github.com/google/uuid"
	"strings"
)

func UUID() string {
	return uuid.New().String()
}

func UUIDShort() string {
	return strings.Replace(UUID(), "-", "", -1)
}
