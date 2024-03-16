package pkg

import (
	"bytes"
	"encoding/binary"
	"errors"
	"unicode/utf8"
)

type MemoryStream struct {
	buffer   []byte
	writePos uint
	readPos  uint
	order    binary.ByteOrder
}

func NewMemoryStream(order binary.ByteOrder) MemoryStream {
	return MemoryStream{order: order}
}

func (p *MemoryStream) Len() uint {
	return uint(len(p.buffer))
}

func (p *MemoryStream) Remains() uint {
	return p.writePos - p.readPos
}

func (p *MemoryStream) Bytes() []byte {
	return p.buffer[:p.writePos]
}

func (p *MemoryStream) checkExpand(length uint) {
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

func (p *MemoryStream) Write(b []byte) error {
	length := uint(len(b))
	if length == 0 {
		return errors.New("b is nil")
	}
	p.checkExpand(length)
	copy(p.buffer[p.writePos:p.writePos+length], b)
	p.writePos += length
	return nil
}

func (p *MemoryStream) Read(len uint) ([]byte, error) {
	remains := p.writePos - p.readPos
	if remains < len {
		return nil, errors.New("EOF")
	}

	buffer := p.buffer[p.readPos : p.readPos+len]
	p.readPos += len
	return buffer, nil
}

func (p *MemoryStream) ReadInt32() (int32, error) {
	buffer, err := p.Read(uint(4))
	if err != nil {
		return 0, err
	}
	bytesBuffer := bytes.NewBuffer(buffer)
	var value int32
	err = binary.Read(bytesBuffer, p.order, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (p *MemoryStream) ReadInt64() (int64, error) {
	buffer, err := p.Read(uint(8))
	if err != nil {
		return 0, err
	}
	bytesBuffer := bytes.NewBuffer(buffer)
	var value int64
	err = binary.Read(bytesBuffer, p.order, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (p *MemoryStream) WriteInt32(value int32) error {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, p.order, value)
	if err != nil {
		return err
	}
	return p.Write(bytesBuffer.Bytes())
}

func (p *MemoryStream) WriteInt64(value int64) error {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, p.order, value)
	if err != nil {
		return err
	}
	return p.Write(bytesBuffer.Bytes())
}

func (p *MemoryStream) WriteString(value string) error {
	length := utf8.RuneCountInString(value)
	if length > 0 {
		err := p.WriteInt32(int32(length))
		if err != nil {
			return err
		}
		return p.Write([]byte(value))
	}
	return nil
}

func (p *MemoryStream) ReadString() (string, error) {
	length, err := p.ReadInt32()
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
