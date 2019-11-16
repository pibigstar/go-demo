package utils

import (
	"crypto"
	"encoding/hex"
	"testing"
)

var (
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDh0zfUMoGMfULKQsaWA6GZSb+iOWet/CRGG3oCMxuIWdohmfiG
psyQCROVkRhFC9gAUgvEDciLyX99Gl7FI9SqarzRKm9bG4oTo424632b5gFXgZ5C
OGBCFIm3EAwqzZEr2Xj0ElA1Nc6WHGwBmnw1hA9AbhcXaZh0u8f6+Tj8eQIDAQAB
AoGBAKrExegiVVL++j3nZzLUBiTb3x122Y95J5kYeBgnu79NayWTwJtakUCujG/D
LB4yiaIcaSdV4PzMYCsjgN0FbnAQv5nwDfonOh16l1ZcMoe0kJUDWk4Fwd4Xw1Gs
AE7ltyXnLTWk2U94NIbfDQmm3f1dUdQYi1/dk67iKB7ycWntAkEA956F3kOpHHIX
/4mYaa1MXLQQyzFIFe+0nVl1YIB90yblADRs9CQHFI0jZLp+DllPU4Ke3XxVcA3q
7dWYMD1lxwJBAOl33ITKc9qNFnY4HlxRRkocLgQyXI26mge2rX6WJufsFYa4sR0S
kqd6Q4L9WyyDaUuNy9/BBFOe23PAPkWhi78CQE+Z+lb1UUv/sY9IYGK4fy/eAvgP
I6lJobpjo8QeClTyz/M85zmky1Hj/VjISvW56DJkb0WsTprzHm7Ol1oKoskCQQC6
3IqNZhTQKfh+anAyZ4KgsmlKRpy5e07pOZcnKDq/ib+48n4fzMvAbCU45FtjB1Lx
e+5findSDmWLAaVVyfS1AkAP35gBW68xoLF80ivBvPrPVZzrL3JMukJOrOEreMbE
3sBfm3nkVjxJSNEAj0bsK3fNYWh2vFRaV6ooA1x7VcOq
-----END RSA PRIVATE KEY-----
`
	publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDh0zfUMoGMfULKQsaWA6GZSb+i
OWet/CRGG3oCMxuIWdohmfiGpsyQCROVkRhFC9gAUgvEDciLyX99Gl7FI9SqarzR
Km9bG4oTo424632b5gFXgZ5COGBCFIm3EAwqzZEr2Xj0ElA1Nc6WHGwBmnw1hA9A
bhcXaZh0u8f6+Tj8eQIDAQAB
-----END PUBLIC KEY-----
`
	content = "Hello Pibigstar"
)

func TestRsa(t *testing.T) {
	bs, err := RSAEncrypt([]byte(content), publicKey)
	if err != nil {
		t.Error(err)
	}
	t.Log(hex.EncodeToString(bs))

	bs, err = RASDecrypt(bs, []byte(privateKey))
	if err != nil {
		t.Error(err)
	}
	t.Log(string(bs))
}

func TestRsaWithKey(t *testing.T) {
	pub, err := ParsePublicKey(publicKey)
	if err != nil {
		t.Error(err)
	}

	bs, err := RSAEncryptWithKey([]byte(content), pub)
	t.Log(hex.EncodeToString(bs))

	pri, err := ParsePrivateKey(privateKey)
	if err != nil {
		t.Error(err)
	}

	bs, err = RsaDecryptWithKey(bs, pri)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(bs))
}

func TestSha256WithRsa(t *testing.T) {
	// Sha256WithRSA
	pri, err := ParsePrivateKey(privateKey)
	if err != nil {
		t.Error(err)
	}
	bs, err := RsaSignWithKey([]byte(content), pri, crypto.SHA256)
	if err != nil {
		t.Error(err)
	}
	t.Log(hex.EncodeToString(bs))

	err = RSAVerify([]byte(content), bs, publicKey, crypto.SHA256)
	if err != nil {
		t.Error(err)
	}
}
