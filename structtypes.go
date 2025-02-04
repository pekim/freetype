package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"fmt"
	"unsafe"
)

func init() {
	assertSameSize(Matrix{}, C.FT_Matrix{})
	assertSameSize(SizeRequestRec{}, C.FT_Size_RequestRec{})
	assertSameSize(Vector{}, C.FT_Vector{})
}

func assertSameSize[A any, B any](a A, b B) {
	if unsafe.Sizeof(a) != unsafe.Sizeof(b) {
		panic(fmt.Sprintf("size of %T (%d) != size of %T (%d)", a, unsafe.Sizeof(a), b, unsafe.Sizeof(b)))
	}
}

type Matrix struct {
	xx, xy Fixed
	yx, yy Fixed
}

type SizeRequestRec struct {
	Type           SizeRequestType
	Width          Long
	Height         Long
	HoriResolution UInt
	VertResolution UInt
}

type Vector struct {
	x Pos
	y Pos
}
