package gmsm

import (
	"errors"
	"github.com/Zhang-jie-jun/gmsm/aes"
	"github.com/Zhang-jie-jun/gmsm/sm4"
)

type ALGORITHM uint8

const (
	ALGORITHM_AES_CBC ALGORITHM = 1
	ALGORITHM_SM4_CBC ALGORITHM = 2
)

type CryptFunc interface {
	Encrypt(origData []byte) []byte
	Decrypt(origData []byte) []byte
}
type ICrypt struct {
	encodeFunc CryptFunc
	decodeFunc CryptFunc
}

//创建加解密工具
func NewCrypt(algorithm ALGORITHM, key, iv []byte) (ICrypt, error) {
	if algorithm == ALGORITHM_AES_CBC {
		return ICrypt{encodeFunc: aes.NewAesCipher256(key, iv), decodeFunc: aes.NewAesCipher256(key, iv)}, nil
	} else if algorithm == ALGORITHM_SM4_CBC {
		key, err := sm4.NewSM4CBCCipher256(key, iv)
		if err != nil {
			return ICrypt{}, err
		}
		return ICrypt{encodeFunc: key, decodeFunc: key}, nil
	}
	return ICrypt{}, errors.New("the algorithm not support")
}

//加密
func (crypto *ICrypt) Encrypt(src []byte) ([]byte, error) {
	return crypto.encodeFunc.Encrypt(src), nil
}

//解密
func (crypto *ICrypt) Decrypt(src []byte) ([]byte, error) {
	return crypto.decodeFunc.Decrypt(src), nil
}
