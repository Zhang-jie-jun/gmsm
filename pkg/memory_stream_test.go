package pkg

import (
	"encoding/base64"
	"encoding/binary"
	"math"
	"testing"
)

func Test_ReadInt32_min(t *testing.T) {
	stream := MemoryStream{order: binary.BigEndian}
	inValue := int32(math.MinInt32)
	err := stream.WriteInt32(inValue)
	if err != nil {
		t.Error(err)
	}
	outValue, err := stream.ReadInt32()
	if err != nil {
		t.Error("Read Int32 Faild", err)
	}
	if outValue != inValue {
		t.Error("not queal", inValue, outValue)
	}
}

func Test_ReadInt32_normal(t *testing.T) {
	stream := MemoryStream{order: binary.BigEndian}
	inValue := int32(1000)
	err := stream.WriteInt32(inValue)
	if err != nil {
		t.Error(err)
	}
	outValue, err := stream.ReadInt32()
	if err != nil {
		t.Error("Read Int32 Faild", err)
	}
	if outValue != inValue {
		t.Error("not queal", inValue, outValue)
	}
}

func Test_ReadInt32_max(t *testing.T) {
	stream := MemoryStream{order: binary.BigEndian}
	inValue := int32(math.MaxInt32)
	err := stream.WriteInt32(inValue)
	if err != nil {
		t.Error(err)
	}
	outValue, err := stream.ReadInt32()
	if err != nil {
		t.Error("Read Int32 Faild", err)
	}
	if outValue != inValue {
		t.Error("not queal", inValue, outValue)
	}
}

func Test_ReadInt64(t *testing.T) {
	stream := MemoryStream{order: binary.BigEndian}
	inValue := int64(math.MinInt64)
	err := stream.WriteInt64(inValue)
	if err != nil {
		t.Error(err)
	}
	outValue, err := stream.ReadInt64()
	if err != nil {
		t.Error("Read Int64 Faild", err)
	}
	if outValue != inValue {
		t.Error("not queal", inValue, outValue)
	}
}

func Test_ReadInt64_normal(t *testing.T) {
	stream := MemoryStream{order: binary.BigEndian}
	inValue := int64(10000000000)
	err := stream.WriteInt64(inValue)
	if err != nil {
		t.Error(err)
	}
	outValue, err := stream.ReadInt64()
	if err != nil {
		t.Error("Read Int64 Faild", err)
	}
	if outValue != inValue {
		t.Error("not queal", inValue, outValue)
	}
}

func Test_ReadInt64_max(t *testing.T) {
	stream := MemoryStream{order: binary.BigEndian}
	inValue := int64(math.MaxInt64)
	err := stream.WriteInt64(inValue)
	if err != nil {
		t.Error(err)
	}
	outValue, err := stream.ReadInt64()
	if err != nil {
		t.Error("Read Int64 Faild", err)
	}
	if outValue != inValue {
		t.Error("not queal", inValue, outValue)
	}
}

func Test_compare_cpp(t *testing.T) {
	cppBase64V := "AAAAZAAAAAAAAABk"
	stream := MemoryStream{order: binary.BigEndian}
	err := stream.WriteInt32(100)
	if err != nil {
		t.Error(err)
	}
	err = stream.WriteInt64(100)
	if err != nil {
		t.Error(err)
	}
	goBase64V := base64.StdEncoding.EncodeToString(stream.Bytes())
	if cppBase64V != goBase64V {
		t.Error("cppBase64V is equal to goBase64V", cppBase64V, goBase64V)
	}
}

func Test_Read(t *testing.T) {
	stream := MemoryStream{order: binary.BigEndian}
	var1 := int32(math.MinInt32)
	var2 := int32(math.MaxInt32)
	var3 := int64(math.MinInt64)
	var4 := int64(math.MaxInt64)

	err := stream.WriteInt32(var1)
	if err != nil {
		t.Error(err)
	}
	err = stream.WriteInt64(var3)
	if err != nil {
		t.Error(err)
	}
	err = stream.WriteInt32(var2)
	if err != nil {
		t.Error(err)
	}
	err = stream.WriteInt64(var4)
	if err != nil {
		t.Error(err)
	}
	outVal1, err := stream.ReadInt32()
	if err != nil {
		t.Error("Read Va1 faild", err)
	}
	if outVal1 != var1 {
		t.Error("Read Va1 not equal", outVal1, var1)
	}

	outVal3, err := stream.ReadInt64()
	if err != nil {
		t.Error("Read var3 faild", err)
	}
	if outVal3 != var3 {
		t.Error("Read var3 not equal", outVal3, var3)
	}

	outVal2, err := stream.ReadInt32()
	if err != nil {
		t.Error("Read var2 faild", err)
	}
	if outVal2 != var2 {
		t.Error("Read var2 not equal", outVal1, var1)
	}

	outVal4, err := stream.ReadInt64()
	if err != nil {
		t.Error("Read var4 faild", err)
	}
	if outVal4 != var4 {
		t.Error("Read var4 not equal", outVal4, var4)
	}
}

func TestMemoryStream_Write(t *testing.T) {
	stream := MemoryStream{order: binary.BigEndian}
	src := []byte{12, 32, 45, 75, 46}
	err := stream.Write(src)
	if err != nil {
		t.Error("Write faild", err)
	}
	if stream.Len() != uint(len(src)+256) {
		t.Error("Len faild", err)
	}
	if stream.Remains() != uint(len(src)) {
		t.Error("Len faild", err)
	}
}

func TestMemoryStream_WriteString(t *testing.T) {
	src := "dsdadahdh423423h4jhjda"
	memoryStream := NewMemoryStream(binary.BigEndian)
	err := memoryStream.WriteString(src)
	if err != nil {
		t.Error(err)
	}
	dst, err := memoryStream.ReadString()
	if err != nil {
		t.Error(err)
	}
	if dst != src {
		t.Error("fail")
	}

}
