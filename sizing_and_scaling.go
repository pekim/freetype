package freetype

import (
	"unsafe"

	"modernc.org/libfreetype"
)

// Functions to manage font sizes.

/*
Size is a handle to an object that models a face scaled to a given character size.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_size
*/
type Size *SizeRec

func init() {
	assertSameSize(SizeRec{}, libfreetype.TFT_SizeRec{})
}

// SizeRec is the FreeType root size class structure. A size object models a face object at a given size.
//
// https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_sizerec
type SizeRec struct {
	Face    libfreetype.TFT_Face
	Generic Generic
	Metrics SizeMetrics
	_       unsafe.Pointer // internal
}

func init() {
	assertSameSize(SizeMetrics{}, libfreetype.TFT_Size_Metrics{})
}

/*
SizeMetrics is the size metrics structure gives the metrics of a size object.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_size_metrics
*/
type SizeMetrics struct {
	Xppem UShort
	Yppem UShort

	XScale Fixed
	YScale Fixed

	Ascender   Pos
	Descender  Pos
	Height     Pos
	MaxAdvance Pos
}

// func init() {
// 	assertSameSize(BitmapSize{}, C.FT_Bitmap_Size{})
// }

/*
BitmapSize structure models the metrics of a bitmap strike (i.e., a set of glyphs for a given
point size and resolution) in a bitmap font.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_bitmap_size
*/
type BitmapSize struct {
	Height Short
	Width  Short

	Size Pos

	Xppem Pos
	YPpem Pos
}

// /*
// SetCharSize requests the nominal size (in points).

// https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_set_char_size
// */
// func (face Face) SetCharSize(
// 	charWidth F26Dot6, charHeight F26Dot6,
// 	horzResolution UInt, vertResolution UInt,
// ) error {
// 	err := C.FT_Set_Char_Size(face.face, charWidth, charHeight, horzResolution, vertResolution)
// 	return newError(err, "failed to set char size for face")
// }

// /*
// SetPixelSizes requests the nominal size (in pixels).

// https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_set_pixel_sizes
// */
// func (face Face) SetPixelSizes(pixelWidth UInt, pixelHeight UInt) error {
// 	err := C.FT_Set_Pixel_Sizes(face.face, pixelWidth, pixelHeight)
// 	return newError(err, "failed to set pixel sizes for face")
// }

// /*
// RequestSize resizes the scale of the active Size object in a face.

// https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_request_size
// */
// func (face Face) RequestSize(req SizeRequestRec) error {
// 	err := C.FT_Request_Size(face.face, toPointer[C.FT_Size_RequestRec](req))
// 	return newError(err, "failed to request size for face")
// }

// /*
// SelectSize selects a bitmap strike.

// https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_select_size
// */
// func (face Face) SelectSize(strikeIndex Int) error {
// 	err := C.FT_Select_Size(face.face, strikeIndex)
// 	return newError(err, "failed to set select size for face")
// }

// /*
// SizeRequestType is an enumeration type that lists the supported size request types,
// i.e., what input size (in font units) maps to the requested output size (in pixels,
// as computed from the arguments of SizeRequest).

// https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_size_request_type
// */
// type SizeRequestType = C.FT_Size_Request_Type

// const (
// 	SIZE_REQUEST_TYPE_NOMINAL  = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_NOMINAL)
// 	SIZE_REQUEST_TYPE_REAL_DIM = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_REAL_DIM)
// 	SIZE_REQUEST_TYPE_BBOX     = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_BBOX)
// 	SIZE_REQUEST_TYPE_CELL     = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_CELL)
// 	SIZE_REQUEST_TYPE_SCALES   = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_SCALES)
// )

// func init() {
// 	assertSameSize(SizeRequestRec{}, C.FT_Size_RequestRec{})
// }

// /*
// SizeRequestRec is a structure to model a size request.

// https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_size_requestrec
// */
// type SizeRequestRec struct {
// 	Type           SizeRequestType
// 	Width          Long
// 	Height         Long
// 	HoriResolution UInt
// 	VertResolution UInt
// }

// // SizeRequest is a handle to a size request structure.
// //
// // https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_size_request
// type SizeRequest *SizeRequestRec

// /*
// SetTransform sets the transformation that is applied to glyph images when they are loaded into a
// glyph slot through FT_Load_Glyph.

// https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_set_transform
// */
// func (face Face) SetTransform(matrix *Matrix, delta *Vector) {
// 	C.FT_Set_Transform(face.face,
// 		(*C.FT_Matrix)(unsafe.Pointer(matrix)),
// 		(*C.FT_Vector)(unsafe.Pointer(delta)),
// 	)
// }

// /*
// GetTransform returns the transformation that is applied to glyph images when they are loaded into
// a glyph slot through Load_Glyph. See SetTransform for more details.

// https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_get_transform
// */
// func (face Face) GetTransform() (Matrix, Vector) {
// 	var matrix C.FT_Matrix
// 	var vector C.FT_Vector
// 	C.FT_Get_Transform(face.face, &matrix, &vector)
// 	return to[Matrix](matrix), to[Vector](vector)
// }
