package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"fmt"
	"unsafe"
)

func init() {
	assertSameSize(SizeRequestRec{}, C.FT_Size_RequestRec{})
}

func assertSameSize[A any, B any](a A, b B) {
	if unsafe.Sizeof(a) != unsafe.Sizeof(b) {
		panic(fmt.Sprintf("size of %T (%d) != size of %T (%d)", a, unsafe.Sizeof(a), b, unsafe.Sizeof(b)))
	}
}

type SizeRequestRec struct {
	Type           SizeRequestType
	Width          Long
	Height         Long
	HoriResolution UInt
	VertResolution UInt
}

func (src SizeRequestRec) toFT() C.FT_Size_RequestRec {
	return *(*C.FT_Size_RequestRec)(unsafe.Pointer(&src))
}
