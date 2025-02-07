package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"unsafe"
)

// Functions to manage font sizes.

/*
Size is a handle to an object that models a face scaled to a given character size.

A Face has one active Size object that is used by functions like Load_Glyph to determine the
scaling transformation that in turn is used to load and hint glyphs and metrics.

A newly created Size object contains only meaningless zero values.
You must use Set_Char_Size, Set_Pixel_Sizes, Request_Size or even Select_Size to change the content
(i.e., the scaling values) of the active FT_Size. Otherwise, the scaling and hinting will not be performed.

You can use New_Size to create additional size objects for a given Face,
but they won't be used by other functions until you activate it through Activate_Size.
Only one size can be activated at any given time per face.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_size
*/
type Size *SizeRec

func init() {
	assertSameSize(SizeRec{}, C.FT_SizeRec{})
}

// SizeRec is the FreeType root size class structure. A size object models a face object at a given size.
//
// https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_sizerec
type SizeRec struct {
	// Handle to the parent face object.
	Face Face
	// A typeless pointer, unused by the FreeType library or any of its drivers.
	// It can be used by client applications to link their own data to each size object.
	Generic Generic
	// Metrics for this size object. This field is read-only.
	Metrics  SizeMetrics
	internal unsafe.Pointer
}

func init() {
	assertSameSize(SizeMetrics{}, C.FT_Size_Metrics{})
}

/*
SizeMetrics is the size metrics structure gives the metrics of a size object.

The scaling values, if relevant, are determined first during a size changing operation. The remaining fields are then set by the driver. For scalable formats, they are usually set to scaled values of the corresponding fields in FaceRec. Some values like ascender or descender are rounded for historical reasons; more precise values (for outline fonts) can be derived by scaling the corresponding FT_FaceRec values manually, with code similar to the following.

	scaled_ascender = FT_MulFix( face->ascender,
	                             size_metrics->y_scale );

Note that due to glyph hinting and the selected rendering mode these values are usually not exact; consequently, they must be treated as unreliable with an error margin of at least one pixel!

Indeed, the only way to get the exact metrics is to render all glyphs. As this would be a definite performance hit, it is up to client applications to perform such computations.

The SizeMetrics structure is valid for bitmap fonts also.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_size_metrics
*/
type SizeMetrics struct {
	// The width of the scaled EM square in pixels, hence the term ‘ppem’ (pixels per EM).
	// It is also referred to as ‘nominal width’.
	Xppem UShort
	// The height of the scaled EM square in pixels, hence the term ‘ppem’ (pixels per EM).
	// It is also referred to as ‘nominal height’.
	Yppem UShort

	// A 16.16 fractional scaling value to convert horizontal metrics from font units to 26.6 fractional pixels.
	// Only relevant for scalable font formats.
	XScale Fixed
	// A 16.16 fractional scaling value to convert vertical metrics from font units to 26.6 fractional pixels.
	// Only relevant for scalable font formats.
	YScale Fixed

	// The ascender in 26.6 fractional pixels, rounded up to an integer value.
	// See FaceRec for the details.
	Ascender Pos
	// The descender in 26.6 fractional pixels, rounded down to an integer value.
	// See FaceRec for the details.
	Descender Pos
	// The height in 26.6 fractional pixels, rounded to an integer value.
	// See FaceRec for the details.
	Height Pos
	// The maximum advance width in 26.6 fractional pixels, rounded to an integer value.
	// See FaceRec for the details.
	MaxAdvance Pos
}

func init() {
	assertSameSize(BitmapSize{}, C.FT_Bitmap_Size{})
}

/*
BitmapSize structure models the metrics of a bitmap strike (i.e., a set of glyphs for a given point size and resolution) in a bitmap font. It is used for the available_sizes field of Face.

Windows FNT: The nominal size given in a FNT font is not reliable. If the driver finds it incorrect, it sets size to some calculated values, and x_ppem and y_ppem to the pixel width and height given in the font, respectively.

TrueType embedded bitmaps: size, width, and height values are not contained in the bitmap strike itself. They are computed from the global font parameters.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_bitmap_size
*/
type BitmapSize struct {
	// The vertical distance, in pixels, between two consecutive baselines. It is always positive.
	Height Short
	// The average width, in pixels, of all glyphs in the strike.
	Width Short

	// The nominal size of the strike in 26.6 fractional points. This field is not very useful.
	Size Pos

	// The horizontal ppem (nominal width) in 26.6 fractional pixels.
	Xppem Pos
	// The vertical ppem (nominal height) in 26.6 fractional pixels.
	YPpem Pos
}

/*
SetCharSize requests the nominal size (in points).

  - charWidth - The nominal width, in 26.6 fractional points.
  - charHeight	- The nominal height, in 26.6 fractional points.
  - horzResolution	- The horizontal resolution in dpi.
  - vertResolution	- The vertical resolution in dpi.

While this function allows fractional points as input values, the resulting ppem value for the given resolution is always rounded to the nearest integer.

If either the character width or height is zero, it is set equal to the other value.

If either the horizontal or vertical resolution is zero, it is set equal to the other value.

A character width or height smaller than 1pt is set to 1pt; if both resolution values are zero, they are set to 72dpi.

Don't use this function if you are using the FreeType cache API.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_set_char_size
*/
func (face Face) SetCharSize(
	charWidth F26Dot6, charHeight F26Dot6,
	horzResolution UInt, vertResolution UInt,
) error {
	err := C.FT_Set_Char_Size(face.face, charWidth, charHeight, horzResolution, vertResolution)
	return newError(err, "failed to set char size for face")
}

/*
SetPixelSizes requests the nominal size (in pixels).

  - pixelWidth - The nominal width, in pixels.
  - pixelHeight	- The nominal height, in pixels.

You should not rely on the resulting glyphs matching or being constrained to this pixel size. Refer to Request_Size to understand how requested sizes relate to actual sizes.

Don't use this function if you are using the FreeType cache API.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_set_pixel_sizes
*/
func (face Face) SetPixelSizes(pixelWidth UInt, pixelHeight UInt) error {
	err := C.FT_Set_Pixel_Sizes(face.face, pixelWidth, pixelHeight)
	return newError(err, "failed to set pixel sizes for face")
}

/*
RequestSize resizes the scale of the active Size object in a face.

  - req	- A pointer to a FT_Size_RequestRec.

Although drivers may select the bitmap strike matching the request, you should not rely on this if you intend to select a particular bitmap strike. Use FT_Select_Size instead in that case.

The relation between the requested size and the resulting glyph size is dependent entirely on how the size is defined in the source face. The font designer chooses the final size of each glyph relative to this size. For more information refer to ‘https://www.freetype.org/freetype2/docs/glyphs/glyphs-2.html’.

Contrary to Set_Char_Size, this function doesn't have special code to normalize zero-valued widths, heights, or resolutions, which are treated as LOAD_NO_SCALE.

Don't use this function if you are using the FreeType cache API.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_request_size
*/
func (face Face) RequestSize(req SizeRequestRec) error {
	err := C.FT_Request_Size(face.face, toPointer[C.FT_Size_RequestRec](req))
	return newError(err, "failed to request size for face")
}

/*
SelectSize selects a bitmap strike. To be more precise, this function sets the scaling factors of the active Size object in a face so that bitmaps from this particular strike are taken by Load_Glyph and friends.

  - strike_index - The index of the bitmap strike in the available_sizes field of FT_FaceRec structure.

For bitmaps embedded in outline fonts it is common that only a subset of the available glyphs at a given ppem value is available. FreeType silently uses outlines if there is no bitmap for a given glyph index.

For GX and OpenType variation fonts, a bitmap strike makes sense only if the default instance is active (that is, no glyph variation takes place); otherwise, FreeType simply ignores bitmap strikes. The same is true for all named instances that are different from the default instance.

Don't use this function if you are using the FreeType cache API.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_select_size
*/
func (face Face) SelectSize(strikeIndex Int) error {
	err := C.FT_Select_Size(face.face, strikeIndex)
	return newError(err, "failed to set select size for face")
}

/*
SizeRequestType is an enumeration type that lists the supported size request types, i.e., what input size (in font units) maps to the requested output size (in pixels, as computed from the arguments of SizeRequest).

The const descriptions only apply to scalable formats. For bitmap formats, the behavior is up to the driver.

See the note section of SizeMetrics if you wonder how size requesting relates to scaling values.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_size_request_type
*/
type SizeRequestType = C.FT_Size_Request_Type

const (
	// The nominal size. The units_per_EM field of FaceRec is used to determine both scaling values.
	//
	// This is the standard scaling found in most applications. In particular, use this size request type
	// for TrueType fonts if they provide optical scaling or something similar. Note, however, that
	// units_per_EM is a rather abstract value which bears no relation to the actual size of the glyphs in a font.
	SIZE_REQUEST_TYPE_NOMINAL = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_NOMINAL)
	// The real dimension. The sum of the ascender and (minus of) the descender fields of FaceRec is
	// used to determine both scaling values.
	SIZE_REQUEST_TYPE_REAL_DIM = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_REAL_DIM)
	// The font bounding box. The width and height of the bbox field of FaceRec are used to determine
	// the horizontal and vertical scaling value, respectively.
	SIZE_REQUEST_TYPE_BBOX = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_BBOX)
	// The max_advance_width field of FaceRec is used to determine the horizontal scaling value; the vertical
	// scaling value is determined the same way as SIZE_REQUEST_TYPE_REAL_DIM does. Finally, both scaling
	// values are set to the smaller one. This type is useful if you want to specify the font size for, say,
	// a window of a given dimension and 80x24 cells.
	SIZE_REQUEST_TYPE_CELL = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_CELL)
	// Specify the scaling values directly.
	SIZE_REQUEST_TYPE_SCALES = SizeRequestType(C.FT_SIZE_REQUEST_TYPE_SCALES)
)

func init() {
	assertSameSize(SizeRequestRec{}, C.FT_Size_RequestRec{})
}

/*
SizeRequestRec is a structure to model a size request.

If width is zero, the horizontal scaling value is set equal to the vertical scaling value, and vice versa.

If type is SIZE_REQUEST_TYPE_SCALES, width and height are interpreted directly as 16.16 fractional scaling values, without any further modification, and both horiResolution and vertResolution are ignored.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_size_requestrec
*/
type SizeRequestRec struct {
	// See SizeRequestType.
	Type SizeRequestType
	// The desired width, given as a 26.6 fractional point value (with 72pt = 1in).
	Width Long
	// The desired height, given as a 26.6 fractional point value (with 72pt = 1in).
	Height Long
	// The horizontal resolution (dpi, i.e., pixels per inch). If set to zero, width is treated as a
	// 26.6 fractional pixel value, which gets internally rounded to an integer.
	HoriResolution UInt
	// The vertical resolution (dpi, i.e., pixels per inch). If set to zero, height is treated as a
	// 26.6 fractional pixel value, which gets internally rounded to an integer.
	VertResolution UInt
}

// SizeRequest is a handle to a size request structure.
//
// https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_size_request
type SizeRequest *SizeRequestRec

/*
SetTransform sets the transformation that is applied to glyph images when they are loaded into a glyph slot through FT_Load_Glyph.

  - matrix - A pointer to the transformation's 2x2 matrix. Use NULL for the identity matrix.
  - delta - A pointer to the translation vector. Use NULL for the null vector.

This function is provided as a convenience, but keep in mind that _Matrix coefficients are only 16.16 fixed-point values, which can limit the accuracy of the results. Using floating-point computations to perform the transform directly in client code instead will always yield better numbers.

The transformation is only applied to scalable image formats after the glyph has been loaded. It means that hinting is unaltered by the transformation and is performed on the character size given in the last call to SetCharSize or SetPixelSizes.

Note that this also transforms the face.glyph.advance field, but not the values in face.glyph.metrics.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_set_transform
*/
func (face Face) SetTransform(matrix *Matrix, delta *Vector) {
	C.FT_Set_Transform(face.face,
		(*C.FT_Matrix)(unsafe.Pointer(matrix)),
		(*C.FT_Vector)(unsafe.Pointer(delta)),
	)
}

/*
GetTransform returns the transformation that is applied to glyph images when they are loaded into a glyph slot through Load_Glyph. See SetTransform for more details.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_get_transform
*/
func (face Face) GetTransform() (Matrix, Vector) {
	var matrix C.FT_Matrix
	var vector C.FT_Vector
	C.FT_Get_Transform(face.face, &matrix, &vector)
	return to[Matrix](matrix), to[Vector](vector)
}
