/**
 * @Author: hiram
 * @Date: 2021/6/25 10:55
 */

package pkg

import (
	"encoding/binary"
	"testing"
	"unsafe"
)

func TestIntToBytes(t *testing.T) {
	src := 213
	if src != BytesToInt(IntToBytes(src)) {
		t.Error("error")
	}
}

func TestUint16ToBytes(t *testing.T) {
	src := uint16(323)
	if src != binary.BigEndian.Uint16(Uint16ToBytes(src, binary.BigEndian)) {
		t.Error("error")
	}
	if src != binary.LittleEndian.Uint16(Uint16ToBytes(src, binary.LittleEndian)) {
		t.Error("error")
	}
}

func TestUint32ToBytes(t *testing.T) {
	src := uint32(323)
	if src != binary.BigEndian.Uint32(Uint32ToBytes(src, binary.BigEndian)) {
		t.Error("error")
	}
	if src != binary.LittleEndian.Uint32(Uint32ToBytes(src, binary.LittleEndian)) {
		t.Error("error")
	}
}

func TestUint64ToBytes(t *testing.T) {
	src := uint64(1624590870000)
	t.Log(Uint64ToBytes(src, binary.BigEndian))
	dst := make([]byte, 8)
	binary.BigEndian.PutUint64(dst, src)
	if src != binary.BigEndian.Uint64(dst) {
		t.Error("error")
	}
}

func TestPutNByte(t *testing.T) {
	src := uint64(154)
	srcLen := unsafe.Sizeof(src)
	dst := make([]byte, 11)
	PutNByte(binary.BigEndian, dst[:srcLen], 8)

}
