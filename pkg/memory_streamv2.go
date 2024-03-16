package pkg

import (
	"bytes"
	"errors"
	"unicode/utf8"
)

type MemoryStreamv2 struct {
	buffer   []byte
	writePos uint
	readPos  uint
}

func NewMemoryStreamv2() MemoryStreamv2 {
	return MemoryStreamv2{}
}

func (p *MemoryStreamv2) Remains() uint {
	return p.writePos - p.readPos
}

func (p *MemoryStreamv2) Len() uint {
	return uint(len(p.buffer))
}

func (p *MemoryStreamv2) Bytes() []byte {
	return p.buffer[:p.writePos]
}

func (p *MemoryStreamv2) checkExpand(length uint) {
	if p.buffer == nil {
		p.buffer = make([]byte, length+256)
	}
	remains := p.Len() - p.writePos
	if length > remains {
		bufferLen := p.Len()
		length = bufferLen + length + 256
		newBuffer := make([]byte, length)
		copy(newBuffer[:bufferLen], p.buffer)
		p.buffer = newBuffer
	}
}

func (p *MemoryStreamv2) Write(b []byte) error {
	length := uint(len(b))
	if length == 0 {
		return errors.New("Write Buffer is Empty")
	}
	p.checkExpand(uint(length))
	copy(p.buffer[p.writePos:p.writePos+length], b)
	p.writePos += length
	return nil
}

func (p *MemoryStreamv2) Read(length uint) ([]byte, error) {
	if length == 0 {
		return nil, errors.New("Read length is Zero")
	}

	avail := p.writePos - p.readPos
	cnt := length
	if avail < length {
		cnt = avail
	}

	if cnt == 0 {
		return nil, errors.New("Buffer is Zero")
	}

	bytesBuffer := bytes.NewBuffer([]byte{})
	bytesBuffer.Write(p.buffer[p.readPos : p.readPos+cnt])
	return bytesBuffer.Bytes(), nil
}

func (p *MemoryStreamv2) writeVarInt32(value uint32) {
	p.checkExpand(8)
	encodeBuffer := EncodeVarint32(value)
	length := uint(len(encodeBuffer))
	copy(p.buffer[p.writePos:p.writePos+length], encodeBuffer)
	p.writePos += length
}

func (p *MemoryStreamv2) readVarInt32() (uint32, error) {
	cnt := p.writePos - p.readPos
	if cnt > 5 {
		cnt = 5
	}
	result, index, err := GetVarint32Ptr(p.buffer[p.readPos : p.readPos+cnt])
	if err != nil {
		return 0, err
	}
	p.readPos += index

	return result, nil
}

func (p *MemoryStreamv2) readVarInt64() (uint64, error) {
	cnt := p.writePos - p.readPos
	if cnt > 10 {
		cnt = 10
	}
	result, index, err := GetVarint64Ptr(p.buffer[p.readPos : p.readPos+cnt])
	if err != nil {
		return 0, err
	}
	p.readPos += index

	return result, nil
}

func (p *MemoryStreamv2) ReadInt32() (int32, error) {
	value, err := p.readVarInt32()
	if err != nil {
		return 0, err
	}
	return DecodeZigzag32(value), nil
}

func (p *MemoryStreamv2) ReadInt64() (int64, error) {
	value, err := p.readVarInt64()
	if err != nil {
		return 0, err
	}
	return DecodeZigzag64(value), nil
}

func (p *MemoryStreamv2) WriteInt32(value int32) {
	encodeValue := EncodeZigzag32(value)
	encodeBuffer := EncodeVarint32(uint32(encodeValue))
	p.Write(encodeBuffer)
}

func (p *MemoryStreamv2) WriteInt64(value int64) {
	encodeValue := EncodeZigzag64(value)
	encodeBuffer := EncodeVarint64(uint64(encodeValue))
	p.Write(encodeBuffer)
}

func (p *MemoryStreamv2) WriteUInt32(value uint32) {
	v := EncodeZigzag32(int32(value))
	p.writeVarInt32(uint32(v))
}

func (p *MemoryStreamv2) ReadUInt32() (uint32, error) {
	v, err := p.readVarInt32()
	if err != nil {
		return 0, err
	}
	return uint32(DecodeZigzag32(v)), nil
}

func (p *MemoryStreamv2) WriteString(value string) error {
	length := utf8.RuneCountInString(value)
	if length > 0 {
		p.WriteUInt32(uint32(length))
		return p.Write([]byte(value))
	}
	return nil
}

func (p *MemoryStreamv2) ReadString() (string, error) {
	length, err := p.ReadUInt32()
	if err != nil {
		return "", err
	}
	if length > 0 {
		res, err := p.Read(uint(length))
		if err != nil {
			return "", err
		}
		return string(res), nil
	}
	return "", nil
}
