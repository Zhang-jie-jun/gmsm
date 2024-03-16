package pkg

import (
	"encoding/base64"
	"fmt"
	"math"
	"testing"
)

func Test_ReadInt32_minv2(t *testing.T) {
	stream := MemoryStreamv2{}
	inValue := int32(math.MinInt32)
	stream.WriteInt32(int32(inValue))
	outValue, err := stream.ReadInt32()
	if err != nil {
		t.Error("Read Int32 Faild", err)
	}
	if outValue != inValue {
		t.Error("not queal", inValue, outValue)
	}
}

func Test_ReadInt32_normalv2(t *testing.T) {
	stream := MemoryStreamv2{}
	inValue := int32(1000)
	stream.WriteInt32(int32(inValue))
	outValue, err := stream.ReadInt32()
	if err != nil {
		t.Error("Read Int32 Faild", err)
	}
	if outValue != inValue {
		t.Error("not queal", inValue, outValue)
	}
}

func Test_ReadInt32_maxv2(t *testing.T) {
	stream := MemoryStreamv2{}
	inValue := int32(math.MaxInt32)
	stream.WriteInt32(int32(inValue))
	outValue, err := stream.ReadInt32()
	if err != nil {
		t.Error("Read Int32 Faild", err)
	}
	if outValue != inValue {
		t.Error("not queal", inValue, outValue)
	}
}

func Test_ReadInt64v2(t *testing.T) {
	stream := MemoryStreamv2{}
	inValue := int64(math.MinInt64)
	stream.WriteInt64(int64(inValue))
	outValue, err := stream.ReadInt64()
	if err != nil {
		t.Error("Read Int64 Faild", err)
	}
	if outValue != inValue {
		t.Error("not queal", inValue, outValue)
	}
}

func Test_ReadInt64_normalv2(t *testing.T) {
	stream := MemoryStreamv2{}
	inValue := int64(10000000000)
	stream.WriteInt64(int64(inValue))
	outValue, err := stream.ReadInt64()
	if err != nil {
		t.Error("Read Int64 Faild", err)
	}
	if outValue != inValue {
		t.Error("not queal", inValue, outValue)
	}
}

func Test_ReadInt64_maxv2(t *testing.T) {
	stream := MemoryStreamv2{}
	inValue := int64(math.MaxInt64)
	stream.WriteInt64(int64(inValue))
	outValue, err := stream.ReadInt64()
	if err != nil {
		t.Error("Read Int64 Faild", err)
	}
	if outValue != inValue {
		t.Error("not queal", inValue, outValue)
	}
}

func Test_compare_cppv2(t *testing.T) {
	stream := MemoryStreamv2{}
	stream.WriteInt32(100)
	stream.WriteInt64(100)
	fmt.Println("cpp_value:", base64.StdEncoding.EncodeToString(stream.Bytes()))
}

func Test_Readv2(t *testing.T) {
	stream := MemoryStreamv2{}
	var1 := int32(math.MinInt32)
	var2 := int32(math.MaxInt32)
	var3 := int64(math.MinInt64)
	var4 := int64(math.MaxInt64)

	stream.WriteInt32(var1)
	stream.WriteInt64(var3)
	stream.WriteInt32(var2)
	stream.WriteInt64(var4)

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

func TestMemoryStreamv2_WriteString(t *testing.T) {
	stream := NewMemoryStreamv2()
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

func TestMemoryStreamv2_ReadString(t *testing.T) {
	src := "dsdadahdh423423h4jhjda"
	memoryStream := NewMemoryStreamv2()
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

func TestMemoryStreamv2_WriteUInt32(t *testing.T) {
	memoryStream := NewMemoryStreamv2()
	memoryStream.WriteInt32(int32(32))
	t.Log(memoryStream.Bytes())
}
