package tiny_aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"github.com/Zhang-jie-jun/gmsm/pkg"
)

const (
	MAX_IV_SIZE  = 16
	IV_HEAD_SIZE = 16
	IV_HEAD      = "ase_iv_head:____"
)

func padding(plainText []byte, blockSize int) []byte {
	n := blockSize - len(plainText)%blockSize
	temp := bytes.Repeat([]byte{byte(n)}, n)
	plainText = append(plainText, temp...)
	return plainText
}

func unPadding(cipherText []byte) []byte {
	end := cipherText[len(cipherText)-1]
	cipherText = cipherText[:len(cipherText)-int(end)]
	return cipherText
}

func paddingKey(key []byte) []byte {
	// make the length multiple of 16 bytes
	length := len(key)
	if (length % 32) != 0 {
		length += 32 - (length % 32)
		tempKey := make([]byte, length)
		copy(tempKey[:len(key)], key)
		return tempKey
	}

	return key
}

func aesCBCEncrypt(plainText []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(paddingKey(key))
	if err != nil {
		return nil, err
	}

	// encrypt string
	plainText = padding(plainText, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	return cipherText, nil
}

func aesCBCDecrypt(cipherText []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(paddingKey(key))
	if err != nil {
		return nil, err
	}

	// decrypt string
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	return plainText, nil
}

func Encrypt(key string, value string) (string, error) {
	iv := make([]byte, 0, MAX_IV_SIZE)
	iv = pkg.RandomBytes(MAX_IV_SIZE, iv)

	encryptBuffer, err := aesCBCEncrypt([]byte(value), []byte(key), iv)
	if err != nil {
		return "", err
	}

	// memory stream
	memoryStream := pkg.NewMemoryStream(binary.BigEndian)
	err = memoryStream.Write([]byte(IV_HEAD))
	if err != nil {
		return "", err
	}
	err = memoryStream.Write(iv)
	if err != nil {
		return "", err
	}
	err = memoryStream.WriteInt32(int32(len(value)))
	if err != nil {
		return "", err
	}
	err = memoryStream.Write(encryptBuffer)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(memoryStream.Bytes()), nil
}

func Decrypt(key string, encryptValue string) (string, error) {
	resBuf, err := base64.StdEncoding.DecodeString(encryptValue)
	if err != nil {
		return "", err
	}

	// decode base64 string to stream
	memoryStream := pkg.NewMemoryStream(binary.BigEndian)
	err = memoryStream.Write(resBuf)
	if err != nil {
		return "", err
	}
	// read magic from buffer
	header, _ := memoryStream.Read(IV_HEAD_SIZE)
	if string(header) != IV_HEAD {
		return "", errors.New("encrypt string invalid")
	}

	iv, err := memoryStream.Read(MAX_IV_SIZE)
	if err != nil {
		return "", err
	}

	length, err := memoryStream.ReadInt32()
	if err != nil {
		return "", err

	}

	remains := int32(memoryStream.Remains())
	encryptBuf, err := memoryStream.Read(uint(remains))
	if err != nil {
		return "", err
	}

	decryptBuffer, err := aesCBCDecrypt(encryptBuf, []byte(key), iv)
	if err != nil {
		return "", err
	}
	return string(decryptBuffer[:length]), nil
}
