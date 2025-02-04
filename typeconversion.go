package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"unsafe"
)

func cBoolToGo(value C.FT_Bool) bool {
	return value != 0
}

// to converts a value to another type in an unsafe manner, without regard to
// whether the conversion is reasonable.
func to[T any, U any](value U) T {
	return *(toPointer[T](value))
}

// toPointer accecpts a value and converts it to a pointer to the value as a different type.
// It does so in an unsafe manner, without regard to whether the conversion is reasonable.
func toPointer[T any, U any](value U) *T {
	return (*T)(unsafe.Pointer(&value))
}
