package aesCbc

import (
	"testing"
)

const (
	CRYPT_KEY_256 = "1~$c31kjtR^@@c2#9&iy"
	CRYPT_KEY_128 = "c31kjtR^@@c2#9&"
)

type testInfo struct {
	key      string //秘钥
	iv       string //iv
	origData string //原文
}

func TestAesCbc256(t *testing.T) {
	tests := []*testInfo{
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "704114497615264",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "es0414497615272",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "es704144222222297615406",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertysdddddduiopzxcvbnmgqo",
			origData: "df041449712619896",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwssaaertyuiopzxcvbnmgqo",
			origData: "fd041449768759744",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "ds41449722659772",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "ff4144977fe304674",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "vfr414497615388",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "1234567890qwertyu",
			origData: "bth14497615418",
		},
		{
			key:      CRYPT_KEY_256,
			iv:       "123456qo",
			origData: "4rb2auut300790842620672",
		},
	}

	for index, test := range tests {
		encrData := AesEncrypt([]byte(test.key), []byte(test.iv), []byte(test.origData))
		origData := AesDecrypt([]byte(test.key), []byte(test.iv), encrData)
		if string(origData) != test.origData {
			t.Log(index, " fail")
		}
	}
}

func BenchmarkAes256(b *testing.B) {
	iv := []byte("123456qo789456xx")
	key := []byte("1234567890abcdef")
	src := []byte{79, 219, 222, 154, 162, 109, 22, 27, 13, 169, 90, 53, 188, 76, 156, 224, 252, 23, 128, 94, 161, 208, 222, 144, 170, 109, 83, 14, 9, 118, 101, 177} //[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	aesCipher := NewAesCipher256(key, iv)
	aesCipher2 := NewAesCipher256(key, iv)
	b.ReportAllocs()

	//b.ReportMetric(11111,"test")
	for i := 0; i < b.N; i++ {
		encrData := aesCipher.Encrypt(src)
		origData := aesCipher2.Decrypt(encrData)
		if string(origData) != string(src) {
			b.Error("fail")
		}
	}
}

func TestAesCbc128(t *testing.T) {
	tests := []*testInfo{
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwerty",
			origData: "704114497615264",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890iopzxc",
			origData: "es0414497615272",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "es704144222222297615406",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwertysdddddduiopzxcvbnmgqo",
			origData: "df041449712619896",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwssaaertyuiopzxcvbnmgqo",
			origData: "fd041449768759744",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "ds41449722659772",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "ff4144977fe304674",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwertyuiopzxcvbnmgqo",
			origData: "vfr414497615388",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "1234567890qwertyu",
			origData: "bth14497615418",
		},
		{
			key:      CRYPT_KEY_128,
			iv:       "123456qo",
			origData: "4rb2auut300790842620672",
		},
	}

	for index, test := range tests {
		encrData := AesEncrypt([]byte(test.key), []byte(test.iv), []byte(test.origData))
		origData := AesDecrypt([]byte(test.key), []byte(test.iv), encrData)
		if string(origData) != test.origData {
			t.Error(index, " fail")
		}
	}
}
