package pkg

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func VarintLength(v uint64) int {
	len := 1
	if v >= 128 {
		v >>= 7
		len++
	}
	return len
}

func EncodeFixed16(value uint16) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, value)
	return bytesBuffer.Bytes()
}

func EncodeFixed32(value uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, value)
	return bytesBuffer.Bytes()
}

func EncodeFixed64(value uint64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, value)
	return bytesBuffer.Bytes()
}

func EncodeVarint32(n uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	B := byte(128)
	if n < (1 << 7) {
		bytesBuffer.WriteByte(byte(n))
	} else if n < (1 << 14) {
		bytesBuffer.WriteByte(byte(n) | B)
		bytesBuffer.WriteByte(byte(n >> 7))
	} else if n < (1 << 21) {
		bytesBuffer.WriteByte(byte(n) | B)
		bytesBuffer.WriteByte(byte(n>>7) | B)
		bytesBuffer.WriteByte(byte(n >> 14))
	} else if n < (1 << 28) {
		bytesBuffer.WriteByte(byte(n) | B)
		bytesBuffer.WriteByte(byte(n>>7) | B)
		bytesBuffer.WriteByte(byte(n>>14) | B)
		bytesBuffer.WriteByte(byte(n >> 21))
	} else {
		bytesBuffer.WriteByte(byte(n) | B)
		bytesBuffer.WriteByte(byte(n>>7) | B)
		bytesBuffer.WriteByte(byte(n>>14) | B)
		bytesBuffer.WriteByte(byte(n>>21) | B)
		bytesBuffer.WriteByte(byte(n >> 28))
	}

	return bytesBuffer.Bytes()
}

func EncodeVarint64(n uint64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	B := byte(128)
	for n > 128 {
		bytesBuffer.WriteByte(((byte(n)) & (B - 1)) | B)
		n >>= 7
	}
	bytesBuffer.WriteByte(byte(n))
	return bytesBuffer.Bytes()
}

func DecodeFixed16(b []byte) uint16 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint16
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return x
}

func DecodeFixed32(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return x
}

func DecodeFixed64(b []byte) uint64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint64
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return x
}

func EncodeFixed64BE(value uint64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, value)
	return bytesBuffer.Bytes()
}

func DecodeFixed64BE(b []byte) uint64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint64
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func EncodeFixed32BE(value uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, value)
	return bytesBuffer.Bytes()
}

func DecodeFixed32BE(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func EncodeZigzag64(value int64) int64 {
	return int64(value<<1) ^ (value >> 63)
}

func DecodeZigzag64(value uint64) int64 {
	return int64((value >> 1) ^ -(value & 1))
}

func EncodeZigzag32(value int32) int32 {
	return (value << 1) ^ (value >> 31)
}

func DecodeZigzag32(value uint32) int32 {
	return int32((value >> 1) ^ -(value & 1))
}

func getVarint32PtrFallback(buffer []byte) (uint32, uint, error) {
	index := uint(0)
	B := byte(128)
	var result uint32
	for shift := 0; shift <= 28; shift += 7 {
		byteV := uint32(buffer[index])
		index += 1
		if (byte(byteV) & B) != 0 {
			result |= uint32(byte(byteV)&127) << shift
		} else {
			result |= byteV << shift
			return result, index, nil
		}
	}

	return 0, 0, errors.New("Return is Null")
}

func GetVarint32Ptr(buffer []byte) (uint32, uint, error) {
	B := byte(128)
	if (buffer[0] & B) == 0 {
		return uint32(buffer[0]), 1, nil
	}
	return getVarint32PtrFallback(buffer)
}

func GetVarint64Ptr(buffer []byte) (uint64, uint, error) {
	index := uint(0)
	B := byte(128)
	var result uint64
	for shift := 0; shift <= 63; shift += 7 {
		byteV := uint64(buffer[index])
		index++
		if (byte(byteV) & B) != 0 {
			result |= uint64(byte(byteV)&127) << shift
		} else {
			result |= byteV << shift
			return result, index, nil
		}
	}
	return 0, 0, errors.New("Return is Null")
}
