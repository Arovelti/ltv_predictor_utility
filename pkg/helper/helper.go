package helper

import "unsafe"

func SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
