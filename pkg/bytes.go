/**
 * @Author: hiram
 * @Date: 2020/12/2 10:32
 */

package pkg

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"math/big"
)

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func Uint16ToBytes(n uint16, order binary.ByteOrder) []byte {
	if order == binary.LittleEndian {
		return []byte{
			byte(n),
			byte(n >> 8),
		}
	} else {
		return []byte{
			byte(n >> 8),
			byte(n),
		}
	}
}
func Uint32ToBytes(n uint32, order binary.ByteOrder) []byte {
	if order == binary.LittleEndian {
		return []byte{
			byte(n),
			byte(n >> 8),
			byte(n >> 16),
			byte(n >> 24),
		}
	} else {
		return []byte{
			byte(n >> 24),
			byte(n >> 16),
			byte(n >> 8),
			byte(n),
		}
	}
}

func Uint64ToBytes(n uint64, order binary.ByteOrder) []byte {
	if order == binary.LittleEndian {
		return []byte{
			byte(n),
			byte(n >> 8),
			byte(n >> 16),
			byte(n >> 24),
			byte(n >> 32),
		}
	} else {
		return []byte{
			byte(n >> 32),
			byte(n >> 24),
			byte(n >> 16),
			byte(n >> 8),
			byte(n),
		}
	}
}

func PutNByte(order binary.ByteOrder, b []byte, v uint64) {
	if order == binary.LittleEndian {
		_ = b[v-1]
		for i := 0; i < int(v); i++ {
			if i == 0 {
				b[0] = byte(v)
				continue
			}
			b[i] = byte(v >> i * 8)
		}
	} else {
		_ = b[v-1]
		for i := int(v - 1); i >= 0; i-- {
			if i == int(v-1) {
				b[v-1] = byte(v)
				continue
			}
			b[i] = byte(v >> (v - 1 - uint64(i)) * 8)
		}
	}

}

//安全随机数
func RandomBytes(len int, dst []byte) []byte {
	//var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		dst = append(dst, str[randomInt.Int64()])
	}
	return dst
}
