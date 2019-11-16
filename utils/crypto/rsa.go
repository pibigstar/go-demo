package utils

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"strings"
)

const (
	publicKeyPrefix = "-----BEGIN PUBLIC KEY-----"
	publicKeySuffix = "-----END PUBLIC KEY-----"

	pkcs1Prefix = "-----BEGIN RSA PRIVATE KEY-----"
	pkcs1Suffix = "-----END RSA PRIVATE KEY-----"

	pkcs8Prefix = "-----BEGIN PRIVATE KEY-----"
	pkcs8Suffix = "-----END PRIVATE KEY-----"
)

var (
	ErrPrivateKey = errors.New("private key is error")
	ErrPublicKey  = errors.New("public key is error")
)

func ParsePrivateKey(privateKey string) (key *rsa.PrivateKey, err error) {
	privateKeyStr := privateKey
	privateKey = strings.Replace(privateKey, pkcs8Prefix, "", 1)
	privateKey = strings.Replace(privateKey, pkcs8Suffix, "", 1)
	data := formatKey(privateKey, pkcs1Prefix, pkcs1Suffix, 64)
	pri, err := parsePKCS1PrivateKey(data)
	if err != nil {
		privateKeyStr = strings.Replace(privateKeyStr, pkcs1Prefix, "", 1)
		privateKeyStr = strings.Replace(privateKeyStr, pkcs1Suffix, "", 1)
		data = formatKey(privateKeyStr, pkcs8Prefix, pkcs8Suffix, 64)
		pri, err = parsePKCS8PrivateKey(data)
		if err != nil {
			return nil, err
		}
	}
	return pri, nil
}

func parsePKCS1PrivateKey(data []byte) (key *rsa.PrivateKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, ErrPrivateKey
	}

	key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, err
}

func parsePKCS8PrivateKey(data []byte) (key *rsa.PrivateKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, ErrPrivateKey
	}

	rawKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	key, ok := rawKey.(*rsa.PrivateKey)
	if ok == false {
		return nil, ErrPrivateKey
	}

	return key, err
}

func ParsePublicKey(publicKey string) (key *rsa.PublicKey, err error) {
	data := formatKey(publicKey, publicKeyPrefix, publicKeySuffix, 64)
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, ErrPublicKey
	}

	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, ErrPublicKey
	}

	return key, err
}

func packageData(originalData []byte, packageSize int) (r [][]byte) {
	var src = make([]byte, len(originalData))
	copy(src, originalData)

	r = make([][]byte, 0)
	if len(src) <= packageSize {
		return append(r, src)
	}
	for len(src) > 0 {
		var p = src[:packageSize]
		r = append(r, p)
		src = src[packageSize:]
		if len(src) <= packageSize {
			r = append(r, src)
			break
		}
	}
	return r
}

// 使用公钥 key 对数据 src 进行 RSA 加密
func RSAEncrypt(src []byte, publicKey string) ([]byte, error) {
	pub, err := ParsePublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	return RSAEncryptWithKey(src, pub)
}

// 使用公钥 key 对数据 src 进行 RSA 加密
func RSAEncryptWithKey(src []byte, key *rsa.PublicKey) ([]byte, error) {
	var data = packageData(src, key.N.BitLen()/8-11)
	var cipher = make([]byte, 0, 0)

	for _, d := range data {
		var c, e = rsa.EncryptPKCS1v15(rand.Reader, key, d)
		if e != nil {
			return nil, e
		}
		cipher = append(cipher, c...)
	}

	return cipher, nil
}

// 使用私钥 key 对数据 cipher 进行 RSA 解密,
func RASDecrypt(cipher, privateKey []byte) ([]byte, error) {
	data, err := RSADecryptWithPKCS1(cipher, privateKey)
	if err != nil {
		data, err = RSADecryptWithPKCS8(cipher, privateKey)
		if err != nil {
			return nil, err
		}
	}
	return data, err
}

// 使用私钥 key 对数据 cipher 进行 RSA 解密, key 的格式为 pkcs1
func RSADecryptWithPKCS1(cipher, privateKey []byte) ([]byte, error) {
	pri, err := parsePKCS1PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	return RsaDecryptWithKey(cipher, pri)
}

// 使用私钥 key 对数据 cipher 进行 RSA 解密，key 的格式为 pkcs8
func RSADecryptWithPKCS8(cipher, key []byte) ([]byte, error) {
	pri, err := parsePKCS8PrivateKey(key)
	if err != nil {
		return nil, err
	}

	return RsaDecryptWithKey(cipher, pri)
}

// 使用私钥 key 对数据 cipher 进行 RSA 解密
func RsaDecryptWithKey(cipher []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	var data = packageData(cipher, privateKey.PublicKey.N.BitLen()/8)
	var plain = make([]byte, 0, 0)

	for _, d := range data {
		var p, e = rsa.DecryptPKCS1v15(rand.Reader, privateKey, d)
		if e != nil {
			return nil, e
		}
		plain = append(plain, p...)
	}
	return plain, nil
}

// 使用私钥签名
func RsaSign(src, privateKey []byte, hash crypto.Hash) ([]byte, error) {
	pri, err := parsePKCS1PrivateKey(privateKey)
	if err != nil {
		pri, err = parsePKCS8PrivateKey(privateKey)
		if err != nil {
			return nil, err
		}
	}
	return RsaSignWithKey(src, pri, hash)
}

// 使用私钥签名
func RsaSignWithKey(src []byte, privateKey *rsa.PrivateKey, hash crypto.Hash) ([]byte, error) {
	var h = hash.New()
	h.Write(src)
	var hashed = h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, privateKey, hash, hashed)
}

// 使用公钥对签名内容进行校验
func RSAVerify(src, sign []byte, publicKey string, hash crypto.Hash) error {
	pub, err := ParsePublicKey(publicKey)
	if err != nil {
		return err
	}
	return RSAVerifyWithKey(src, sign, pub, hash)
}

func RSAVerifyWithKey(src, sign []byte, key *rsa.PublicKey, hash crypto.Hash) error {
	var h = hash.New()
	h.Write(src)
	var hashed = h.Sum(nil)
	return rsa.VerifyPKCS1v15(key, hash, hashed, sign)
}

func formatKey(raw, prefix, suffix string, lineCount int) []byte {
	if raw == "" {
		return nil
	}
	raw = strings.Replace(raw, prefix, "", 1)
	raw = strings.Replace(raw, suffix, "", 1)
	raw = strings.Replace(raw, " ", "", -1)
	raw = strings.Replace(raw, "\n", "", -1)
	raw = strings.Replace(raw, "\r", "", -1)
	raw = strings.Replace(raw, "\t", "", -1)

	var sl = len(raw)
	var c = sl / lineCount
	if sl%lineCount > 0 {
		c = c + 1
	}

	var buf bytes.Buffer
	buf.WriteString(prefix + "\n")
	for i := 0; i < c; i++ {
		var b = i * lineCount
		var e = b + lineCount
		if e > sl {
			buf.WriteString(raw[b:])
		} else {
			buf.WriteString(raw[b:e])
		}
		buf.WriteString("\n")
	}
	buf.WriteString(suffix)
	return buf.Bytes()
}
