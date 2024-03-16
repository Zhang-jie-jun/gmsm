package aes

import (
	"testing"
)

const (
	CRYPT_KEY_256 = "1~$c31kjtR^@@c2#9&iy"
	CRYPT_KEY_128 = "c31kjtR^@@c2#9&"
)

var (
	key2    = []byte{221, 95, 179, 0, 186, 114, 83, 205, 8, 211, 58, 146, 246, 222, 18, 189, 194, 225, 252, 92, 91, 20, 226, 82, 166, 74, 110, 64, 197, 173, 167, 57}
	IV2     = []byte{48, 48, 65, 79, 110, 65, 52, 55, 105, 78, 113, 87, 115, 115, 56, 52}
	encdata = []byte{24, 111, 56, 62, 23, 30, 249, 81, 82, 145, 232, 178, 135, 25, 15, 171, 95, 188, 106, 200, 138, 114, 247, 32, 151, 47, 114, 8, 95, 205, 159, 211}
	src2    = "1234566423243242"
)

func TestAES256(t *testing.T) {
	crypt := NewAesCipher256(key2, IV2)
	dest := crypt.Encrypt([]byte(src2))
	data := crypt.Decrypt(dest)
	t.Log(string(data))
	if src2 != string(data) {
		t.Error("fail")
	}
}

func BenchmarkAes256(b *testing.B) {
	iv := []byte("123456qo789456xx")
	key := []byte("1234567890abcdef")
	origData := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	crypt := NewAesCipher256(key, iv)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		encrypted := crypt.Encrypt(origData)
		decrypted := crypt.Decrypt(encrypted)
		if string(origData) != string(decrypted) {
			b.Error("fail")
		}
	}
}

func TestEquals(t *testing.T) {
	//iv := []byte("123456qo789456xx")
	//key := []byte("1234567890abcdef")
	//src := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	//fmt.Println(src)
	//aesCipher := aesCbc.NewAesCipher256(key, iv)
	//aesCipher2 :=aesCbc. NewAesCipher256(key, iv)
	//encrData := aesCipher.Encrypt(src)
	//origData := aesCipher2.Decrypt(encrData)
	//fmt.Println(origData)
	//
	//
	//encrypted := AesEncryptCBC(src, key, iv)
	//decrypted := AesDecryptCBC(encrypted, key,iv)
	//fmt.Println(decrypted)
}
