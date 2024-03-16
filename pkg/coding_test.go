package pkg

import (
	"math"
	"testing"
)

func Test_DecodeZigzag32(t *testing.T) {
	v := int32(math.MinInt32)
	v_encode := EncodeZigzag32(v)
	v_decode := DecodeZigzag32(uint32(v_encode))
	if v_decode != v {
		t.Error("DecodeZigzag32 faild", v_decode, v)
	}
}

func Test_DecodeZigzag32_Max(t *testing.T) {
	v := int32(math.MaxInt32)
	v_encode := EncodeZigzag32(v)
	v_decode := DecodeZigzag32(uint32(v_encode))
	if v_decode != v {
		t.Error("DecodeZigzag32 faild", v_decode, v)
	}
}

func Test_DecodeZigzag64(t *testing.T) {
	v := int64(math.MinInt64)
	v_encode := EncodeZigzag64(v)
	v_decode := DecodeZigzag64(uint64(v_encode))
	if v_decode != v {
		t.Error("DecodeZigzag64 faild", v, v_decode)
	}
}

func Test_DecodeZigzag64_max(t *testing.T) {
	v := int64(math.MaxInt64)
	v_encode := EncodeZigzag64(v)
	v_decode := DecodeZigzag64(uint64(v_encode))
	if v_decode != v {
		t.Error("DecodeZigzag64 faild")
	}
}
