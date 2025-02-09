package freetype

import (
	"unsafe"

	"modernc.org/libc"
)

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
