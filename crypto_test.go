package gmsm

import (
	"testing"
)

var (
	iv  = []byte("123456qo789456xx")
	key = []byte("1234567890abcdef")
	src = []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
)

func TestICrypt(t *testing.T) {
	crypt, err := NewCrypt(ALGORITHM_AES_CBC, key, iv)
	if err != nil {
		t.Error(err)
	}
	dest, err := crypt.Encrypt(src)
	if err != nil {
		t.Error(err)
	}

	data, err := crypt.Decrypt(dest)
	if string(src) != string(data) {
		t.Error("fail")
	}

	crypt, err = NewCrypt(ALGORITHM_SM4_CBC, key, iv)
	if err != nil {
		t.Error(err)
	}
	dest, err = crypt.Encrypt(src)
	if err != nil {
		t.Error(err)
	}
	data, err = crypt.Decrypt(dest)
	if string(src) != string(data) {
		t.Error("fail")
	}
}

func BenchmarkICrypt_AES(b *testing.B) {
	crypt, err := NewCrypt(ALGORITHM_AES_CBC, key, iv)
	if err != nil {
		b.Error(err)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		dest, err := crypt.Encrypt(src)
		if err != nil {
			b.Error(err)
		}
		data, err := crypt.Decrypt(dest)
		if string(src) != string(data) {
			b.Error("fail")
		}
	}
}
func BenchmarkICrypt_SM4(b *testing.B) {
	crypt, err := NewCrypt(ALGORITHM_SM4_CBC, key, iv)
	if err != nil {
		b.Error(err)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		dest, err := crypt.Encrypt(src)
		if err != nil {
			b.Error(err)
		}
		data, err := crypt.Decrypt(dest)
		if string(src) != string(data) {
			b.Error("fail")
		}
	}
}
