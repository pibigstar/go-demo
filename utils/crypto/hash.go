package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/spaolacci/murmur3"
)

// 返回64位字符串
func HashStr(data []byte) string {
	bs := sha256.Sum256(data)
	return hex.EncodeToString(bs[:])
}

// 返回数字Hash值
func HashNum(data []byte) uint64 {
	return murmur3.Sum64(data)
}
