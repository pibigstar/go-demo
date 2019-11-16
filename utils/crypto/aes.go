package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/pbkdf2"
	"hash"
)

const (
	Pkcs5SaltLength   = 8
	Pkcs5DefaultIter  = 2048
	EvpMaxIvLength    = 16
	Pkcs5DefaultMagic = "Salted__"
)

var (
	ErrKeyLength = errors.New("the key length is illegal")
)

func AesEncryptWithSalt(plaintext, key []byte, iterCount int, magic string, h func() hash.Hash) (dst []byte, err error) {
	if iterCount <= 0 {
		iterCount = Pkcs5DefaultIter
	}

	if h == nil {
		h = md5.New
	}

	salt := make([]byte, Pkcs5SaltLength)
	_, err = rand.Read(salt)

	var sKey = pbkdf2.Key(key, salt, iterCount, len(key), h)
	var sIV = pbkdf2.Key(sKey, salt, iterCount, EvpMaxIvLength, h)

	dst, err = AesCbcDecrypt(plaintext, sKey, sIV)

	dst = append(salt, dst...)
	dst = append([]byte(magic), dst...)

	return dst, err
}

func AesDecryptWithSalt(cipherText, key []byte, iterCount int, magic string, h func() hash.Hash) (dst []byte, err error) {
	if iterCount <= 0 {
		iterCount = Pkcs5DefaultIter
	}
	if h == nil {
		h = md5.New
	}
	salt := cipherText[len(magic) : len(magic)+Pkcs5SaltLength]
	sKey := pbkdf2.Key(key, salt, iterCount, len(key), h)
	sIV := pbkdf2.Key(sKey, salt, iterCount, EvpMaxIvLength, h)

	dst, err = AesCbcDecrypt(cipherText[len(magic)+Pkcs5SaltLength:], sKey, sIV)

	return dst, err
}

// CBC模式
func AesCbcEncrypt(plaintext, key, iv []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrKeyLength
	}

	var block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var blockSize = block.BlockSize()
	iv = iv[:blockSize]

	var src = PKCS7Padding(plaintext, blockSize)
	var dst = make([]byte, len(src))

	var mode = cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(dst, src)
	return dst, nil
}

func AesCbcDecrypt(cipherText, key, iv []byte) ([]byte, error) {
	var block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var blockSize = block.BlockSize()
	iv = iv[:blockSize]

	var dst = make([]byte, len(cipherText))

	var mode = cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(dst, cipherText)
	dst = PKCS7UnPadding(dst)
	return dst, nil
}

// CFB模式
func AesCfbEncrypt(plaintext, key, iv []byte) ([]byte, error) {
	var block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var blockSize = block.BlockSize()
	iv = iv[:blockSize]

	var dst = make([]byte, len(plaintext))

	var mode = cipher.NewCFBEncrypter(block, iv)
	mode.XORKeyStream(dst, plaintext)
	return dst, nil
}

func AesCfbDecrypt(cipherText, key, iv []byte) ([]byte, error) {
	var block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var blockSize = block.BlockSize()
	iv = iv[:blockSize]

	var dst = make([]byte, len(cipherText))

	var mode = cipher.NewCFBDecrypter(block, iv)
	mode.XORKeyStream(dst, cipherText)
	return dst, nil
}

func PKCS7Padding(text []byte, blockSize int) []byte {
	var diff = blockSize - len(text)%blockSize
	var paddingText = bytes.Repeat([]byte{byte(diff)}, diff)
	return append(text, paddingText...)
}

func PKCS7UnPadding(text []byte) []byte {
	var length = len(text)
	var unPadding = int(text[length-1])
	return text[:(length - unPadding)]
}
