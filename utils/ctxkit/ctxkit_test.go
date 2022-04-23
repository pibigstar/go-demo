package ctxkit

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContextValue(t *testing.T) {
	ctx := context.Background()
	testValue := "派大星"
	ctx = SetValue(ctx, "test", testValue)
	value := GetValue(ctx, "test")
	assert.Equal(t, value, testValue)
}

func TestContextHeader(t *testing.T) {
	ctx := context.Background()
	testValue := "派大星"
	ctx = AppendHeader(ctx, "test", testValue)

	assert.Equal(t, GetHeader(ctx, "test"), testValue)
}
