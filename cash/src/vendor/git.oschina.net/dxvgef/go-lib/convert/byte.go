package convert

import (
	"strconv"
	"unsafe"
)

//[]byte转string
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//[]byte转换为字符串
func BytesMustStr(value []byte, defaults ...string) string {
	if len(value) > 0 {
		return BytesToStr(value)
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	return ""
}

//[]byte转换为Int
func BytesMustInt(value []byte, defaults ...int) int {
	str := BytesToStr(value)
	v, err := strconv.Atoi(str)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return v
}

//[]byte转换为int32
func BytesMustInt32(value []byte, defaults ...int32) int32 {
	str := BytesToStr(value)
	v, err := strconv.ParseInt(str, 10, 32)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return int32(v)
}

//[]byte转换为int64
func BytesMustInt64(value []byte, defaults ...int64) int64 {
	str := BytesToStr(value)
	v, err := strconv.ParseInt(str, 10, 64)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return v
}

//[]byte转换为uint
func BytesMustUint(value []byte, defaults ...uint) uint {
	str := BytesToStr(value)
	v, err := strconv.ParseUint(str, 10, 64)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return uint(v)
}

//[]byte转换为uint32
func BytesMustUint32(value []byte, defaults ...uint32) uint32 {
	str := BytesToStr(value)
	v, err := strconv.ParseUint(str, 10, 32)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return uint32(v)
}

//[]byte转换为uint64
func BytesMustUint64(value []byte, defaults ...uint64) uint64 {
	str := BytesToStr(value)
	v, err := strconv.ParseUint(str, 10, 64)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return v
}

//[]byte转换为float32
func BytesMustFloat32(value []byte, defaults ...float32) float32 {
	str := BytesToStr(value)
	v, err := strconv.ParseFloat(str, 32)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return float32(v)
}

//[]byte转换为float64
func BytesMustFloat64(value []byte, defaults ...float64) float64 {
	str := BytesToStr(value)
	v, err := strconv.ParseFloat(str, 64)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return v
}

//[]byte转换为bool
func BytesMustBool(value []byte, defaults ...bool) bool {
	str := BytesToStr(value)
	v, err := strconv.ParseBool(str)
	if len(defaults) > 0 && err != nil {
		return defaults[0]
	}
	return v
}
