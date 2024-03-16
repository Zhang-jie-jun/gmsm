package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/Zhang-jie-jun/gmsm/pkg"
	"io"
)

type AesCipher256 struct {
	encBlockMode cipher.BlockMode
	blockSize    int
	BufferPool   *pkg.LimitedPool
	decBlockMode cipher.BlockMode
}

func NewAesCipher256(key, iv []byte) *AesCipher256 {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	return &AesCipher256{encBlockMode: cipher.NewCBCEncrypter(block, iv), decBlockMode: cipher.NewCBCDecrypter(block, iv), blockSize: block.BlockSize(), BufferPool: pkg.NewLimitedPool(1, 4096)}
}

// =================== CBC ======================
func (aesCipher *AesCipher256) Encrypt(origData []byte) []byte {
	origData = pkcs5Padding(origData, aesCipher.blockSize) // 补全码
	//encrypted = make([]byte, len(origData))        // 创建数组
	encrypted := *aesCipher.BufferPool.Get(len(origData))
	defer aesCipher.BufferPool.Put(&encrypted)
	aesCipher.encBlockMode.CryptBlocks(encrypted, origData) // 加密
	return encrypted
}
func (aesCipher *AesCipher256) Decrypt(origData []byte) []byte {
	//	decrypted = make([]byte, len(encrypted))       // 创建数组
	decrypted := *aesCipher.BufferPool.Get(len(origData))
	defer aesCipher.BufferPool.Put(&decrypted)
	aesCipher.decBlockMode.CryptBlocks(decrypted, origData) // 解密
	decrypted = pkcs5UnPadding(decrypted)                   // 去除补全码
	return decrypted
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// =================== ECB ======================
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// =================== CFB ======================
func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}
func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}
