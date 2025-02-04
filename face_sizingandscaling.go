package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"unsafe"
)

type SizeRequestType = C.FT_Size_Request_Type

const (
	SIZE_REQUEST_TYPE_NOMINAL  = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_NOMINAL)
	SIZE_REQUEST_TYPE_REAL_DIM = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_REAL_DIM)
	SIZE_REQUEST_TYPE_BBOX     = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_BBOX)
	SIZE_REQUEST_TYPE_CELL     = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_CELL)
	SIZE_REQUEST_TYPE_SCALES   = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_SCALES)
)

func (face Face) SetCharSize(
	charWidth F26Dot6, charHeight F26Dot6,
	horzResolution UInt, vertResolution UInt,
) error {
	err := C.FT_Set_Char_Size(face.face, charWidth, charHeight, horzResolution, vertResolution)
	return newError(err, "failed to set char size for face")
}

func (face Face) SetPixelSizes(pixelWidth UInt, pixelHeight UInt) error {
	err := C.FT_Set_Pixel_Sizes(face.face, pixelWidth, pixelHeight)
	return newError(err, "failed to set pixel sizes for face")
}

func (face Face) RequestSize(size SizeRequestRec) error {
	err := C.FT_Request_Size(face.face, toPointer[C.FT_Size_RequestRec](size))
	return newError(err, "failed to request size for face")
}

func (face Face) SelectSize(strikeIndex Int) error {
	err := C.FT_Select_Size(face.face, strikeIndex)
	return newError(err, "failed to set select size for face")
}

func (face Face) SetTransform(matrix *Matrix, delta *Vector) {
	C.FT_Set_Transform(face.face,
		(*C.FT_Matrix)(unsafe.Pointer(matrix)),
		(*C.FT_Vector)(unsafe.Pointer(delta)),
	)
}

func (face Face) GetTransform() (Matrix, Vector) {
	var matrix C.FT_Matrix
	var vector C.FT_Vector
	C.FT_Get_Transform(face.face, &matrix, &vector)
	return to[Matrix](matrix), to[Vector](vector)
}
