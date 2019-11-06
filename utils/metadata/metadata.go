package main

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"go-demo/utils/token"
	"google.golang.org/grpc/metadata"
	"time"
)

// 为Context 设值 和取值

const (
	ContextMDTokenKey = "token"
	ContextMDReqIDKey = "req-id"
)

// 将token放到上下文中
func mockTokenContext(tokenKey string) context.Context {

	// 生成token
	claims := make(jwt.MapClaims)
	claims[tokenKey] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims[tokenKey] = "this is user id"
	token, err := utils.GenJwtToken(claims)
	if err != nil {
		return nil
	}

	md := metadata.New(map[string]string{
		tokenKey: token,
	})

	return metadata.NewOutgoingContext(context.Background(), md)
}

// 从上下文中取出token
func GetTokenFromContext(ctx context.Context) string {
	// incoming
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if md[ContextMDTokenKey] != nil && len(md[ContextMDTokenKey]) > 0 {
			return md[ContextMDTokenKey][0]
		}
	}

	// outcoming
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		if md[ContextMDTokenKey] != nil && len(md[ContextMDTokenKey]) > 0 {
			return md[ContextMDTokenKey][0]
		}
	}
	return ""
}
