package main

import (
	"testing"
)

func TestGetTokenFromContext(t *testing.T) {
	tokenContext := mockTokenContext(ContextMDTokenKey)

	token := GetTokenFromContext(tokenContext)
	t.Log(token)
}
