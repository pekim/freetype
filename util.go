package freetype

import (
	"fmt"
	"unsafe"

	"modernc.org/libc"
)

func assertSameSize[A any, B any](a A, b B) {
	if unsafe.Sizeof(a) != unsafe.Sizeof(b) {
		panic(fmt.Sprintf("size of %T (%d) != size of %T (%d)", a, unsafe.Sizeof(a), b, unsafe.Sizeof(b)))
	}
}

// formatTag returns a formatted representation of the 4 bytes of a tag.
func formatTag(tag uint32) string {
	return fmt.Sprintf("'%s', '%s', '%s', '%s'",
		string(rune(tag>>24&0x000000ff)),
		string(rune(tag>>16&0x000000ff)),
		string(rune(tag>>8&0x000000ff)),
		string(rune(tag>>0&0x000000ff)),
	)
}

func toUintptr[T any](ptr *T) uintptr {
	return uintptr(unsafe.Pointer(ptr))
}

func fromUintptr[T any](ptr uintptr) *T {
	// Jump through some hoops to avoid "possible misuse of unsafe.Pointer" warning.
	return (*T)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr)))
}

func alloc[T any](tls *libc.TLS, exampleObject T) (*T, func()) {
	size := int(unsafe.Sizeof(exampleObject))
	object := fromUintptr[T](tls.Alloc(size))
	return object, func() { tls.Free(size) }
}

func goStringForNotNullTerminatedCString(str *Byte, strLen UInt) string {
	// The string is not guaranteed to be null-terminated, so create buffer with one more
	// byte than string_len
	buffer := make([]byte, strLen+1)
	// Copy the string to the buffer.
	copy(buffer[:strLen], unsafe.Slice((*byte)(unsafe.Pointer(str)), strLen))
	// Add a null to ensure that the string is terminated.
	buffer[strLen] = 0

	return libc.GoString(toUintptr(&buffer[0]))
}
