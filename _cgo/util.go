package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"fmt"
	"unsafe"
)

// formatTag returns a formatted representation of the 4 bytes of a tag.
func formatTag(tag uint32) string {
	return fmt.Sprintf("'%s', '%s', '%s', '%s'",
		string(rune(tag>>24&0x000000ff)),
		string(rune(tag>>16&0x000000ff)),
		string(rune(tag>>8&0x000000ff)),
		string(rune(tag>>0&0x000000ff)),
	)
}

func assertSameSize[A any, B any](a A, b B) {
	if unsafe.Sizeof(a) != unsafe.Sizeof(b) {
		panic(fmt.Sprintf("size of %T (%d) != size of %T (%d)", a, unsafe.Sizeof(a), b, unsafe.Sizeof(b)))
	}
}

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

func goStringForNotNullTerminatedCString(str *Byte, strLen UInt) string {
	// The string is not guaranteed to be null-terminated, so create buffer with one more
	// byte than string_len
	buffer := make([]byte, strLen+1)
	// Copy the string to the buffer.
	copy(buffer[:strLen], unsafe.Slice((*byte)(unsafe.Pointer(str)), strLen))
	// Add a null to ensure that the string is terminated.
	buffer[strLen] = 0

	return C.GoString((*C.char)(unsafe.Pointer(&buffer[0])))
}
