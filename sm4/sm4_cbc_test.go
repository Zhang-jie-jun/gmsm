package sm4

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSM4CBC(t *testing.T) {
	iv := []byte("123456qo789456xx")
	a := []byte("erl1233312")
	key := []byte("aabbccddaabbccdd")
	sm4, err := NewSM4CBCCipher256(key, iv)
	if err != nil {
		t.Error(err)
	}
	decrypto := sm4.Encrypt(a)

	fmt.Println("sm4加密后：", hex.EncodeToString(decrypto))
	i := sm4.Decrypt(decrypto)

	fmt.Println("sm4解密后：", string(i))
}

func BenchmarkSM4CBC(t *testing.B) {
	iv := []byte("123456qo789456xx")
	key := []byte("1234567890abcdef")
	data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	sm4, err := NewSM4CBCCipher256(key, iv)
	if err != nil {
		t.Error(err)
	}
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		decrypto := sm4.Encrypt(data)
		i := sm4.Decrypt(decrypto)

		if string(i) != string(data) {
			t.Error("fail")
		}
	}
}
