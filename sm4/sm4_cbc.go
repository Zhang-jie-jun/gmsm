package sm4

import (
	"bytes"
	"crypto/cipher"
	"errors"
	"fmt"
	"github.com/Zhang-jie-jun/gmsm/pkg"
)

type SM4CBCCipher256 struct {
	block      cipher.Block
	iv         []byte
	BufferPool *pkg.LimitedPool
}

func NewSM4CBCCipher256(key, iv []byte) (*SM4CBCCipher256, error) {
	block, e := NewCipher(key)
	if e != nil {
		return nil, errors.New(fmt.Sprintf("newcipher faild,%v! ", e))
	}
	return &SM4CBCCipher256{block: block, iv: iv, BufferPool: pkg.NewLimitedPool(1, 4096)}, nil
}

func (sm4 *SM4CBCCipher256) Encrypt(src []byte) []byte {

	a := sm4.block.BlockSize() - len(src)%sm4.block.BlockSize()
	repeat := bytes.Repeat([]byte{byte(a)}, a)
	newSrc := append(src, repeat...)

	dst := *sm4.BufferPool.Get(len(newSrc)) // make([]byte, len(newSrc))
	defer sm4.BufferPool.Put(&dst)
	blockMode := cipher.NewCBCEncrypter(sm4.block, sm4.iv) //key[:block.BlockSize()]
	blockMode.CryptBlocks(dst, newSrc)
	return dst
}

func (sm4 *SM4CBCCipher256) Decrypt(src []byte) []byte {
	blockMode := cipher.NewCBCDecrypter(sm4.block, sm4.iv) //key[:block.BlockSize()]

	dst := *sm4.BufferPool.Get(len(src)) //make([]byte, len(src))
	defer sm4.BufferPool.Put(&dst)
	blockMode.CryptBlocks(dst, src)

	num := int(dst[len(dst)-1])
	return dst[:len(dst)-num]
}
