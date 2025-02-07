package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"math"
	"strings"
	"unsafe"
)

// The basic data types defined by the library.

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_byte
type Byte = C.FT_Byte

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_bytes
type Bytes = C.FT_Bytes

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_char
type Char = C.FT_Char

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_int
type Int = C.FT_Int

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_uint
type UInt = C.FT_UInt

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_int16
type Int16 = C.FT_Int16

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_uint16
type UInt16 = C.FT_UInt16

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_int32
type Int32 = C.FT_Int32

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_uint32
type UInt32 = C.FT_UInt32

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_int64
type Int64 = C.FT_Int64

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_uint64
type UInt64 = C.FT_UInt64

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_short
type Short = C.FT_Short

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_ushort
type UShort = C.FT_UShort

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_long
type Long = C.FT_Long

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_ulong
type ULong = C.FT_ULong

// Bool is a type of unsigned char, used for simple booleans.
// As usual, values 1 and 0 represent true and false, respectively.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_bool
type Bool = C.FT_Bool

// Offset is equivalent to the ANSI C size_t type,
// i.e., the largest unsigned integer type used to express a file size or position, or a memory block size.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_offset
type Offset = C.FT_Offset

// PtrDist is equivalent to the ANSI C ptrdiff_t type,
// i.e., the largest signed integer type used to express the distance between two pointers.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_ptrdist
type PtrDist = C.FT_PtrDist

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_string
type String = C.FT_String

// Tag is type for 32-bit tags (as used in the SFNT format).
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_tag
type Tag = C.FT_Tag

// For FT_Error see error.go file.

// Fixed is used to store 16.16 fixed-point values, like scaling values or matrix coefficients.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_fixed
type Fixed = C.FT_Fixed

// Pointer is a simple type for a typeless pointer.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_pointer
type Pointer = C.FT_Pointer

// Pos is used to store vectorial coordinates.
// Depending on the context, these can represent distances in integer font units,
// or 16.16, or 26.6 fixed-point pixel coordinates.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_pos
type Pos = C.FT_Pos

func init() {
	assertSameSize(Vector{}, C.FT_Vector{})
}

// Vector is a simple structure used to store a 2D vector; coordinates are of the FT_Pos type.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_vector
type Vector struct {
	X Pos
	Y Pos
}

func init() {
	assertSameSize(BBox{}, C.FT_BBox{})
}

// BBox is a structure used to hold an outline's bounding box,
// i.e., the coordinates of its extrema in the horizontal and vertical directions.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_bbox
type BBox struct {
	XMin Pos
	YMin Pos
	XMax Pos
	YMax Pos
}

func init() {
	assertSameSize(Matrix{}, C.FT_Matrix{})
}

// Matrix is a simple structure used to store a 2x2 matrix.
// Coefficients are in 16.16 fixed-point format.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_matrix
type Matrix struct {
	XX, XY Fixed
	YX, YY Fixed
}

// Word is a signed 16-bit integer used to store a distance in original font units.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_word
type FWord = C.FT_FWord

// UWord is an unsigned 16-bit integer used to store a distance in original font units.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_uword
type UFWord = C.FT_UFWord

// F2Dot14 is a signed 2.14 fixed-point type used for unit vectors.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_f2dot14
type F2Dot14 = C.FT_F2Dot14

func init() {
	assertSameSize(UnitVector{}, C.FT_UnitVector{})
}

// UnitVector is a simple structure used to store a 2D vector unit vector. Uses F2Dot14 types.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_unitvector
type UnitVector struct {
	X F2Dot14
	Y F2Dot14
}

// F26Dot6 is a signed 26.6 fixed-point type used for vectorial pixel coordinates.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_f26dot6
type F26Dot6 = C.FT_F26Dot6

func init() {
	assertSameSize(Data{}, C.FT_Data{})
}

// Data is read-only binary data represented as a pointer and a length.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_data
type Data struct {
	Pointer *Byte
	Length  UInt
}

// FT_MAKE_TAG
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_make_tag

func init() {
	assertSameSize(Generic{}, C.FT_Generic{})
}

/*
Client applications often need to associate their own data to a variety of FreeType core objects.
For example, a text layout API might want to associate a glyph cache to a given size object.

https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_generic
*/
type Generic struct {
	Data      unsafe.Pointer
	Finalizer GenericFinalizer
}

// GenericFinalizer describes a function used to destroy the ‘client’ data of any FreeType object.
// See the description of the Generic type for details of usage.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_generic_finalizer
type GenericFinalizer func(object unsafe.Pointer)

func init() {
	assertSameSize(Bitmap{}, C.FT_Bitmap{})
}

// Bitmap is a structure used to describe a bitmap or pixmap to the raster.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_bitmap
type Bitmap struct {
	Rows         UInt
	Width        UInt
	Pitch        Int
	buffer       *C.uchar
	NumGrays     UShort
	PixelMode    PixelMode
	palette_mode C.uchar
	palette      unsafe.Pointer
}

// Buffer returns the Bitmap's buffer as byte slice.
func (bm Bitmap) Buffer() []byte {
	return unsafe.Slice((*byte)(bm.buffer), bm.Rows*UInt(math.Abs(float64(bm.Pitch))))
}

// BufferVisualization returns a string containing an ascii representation of a grayscale buffer.
func (bm Bitmap) BufferVisualization() string {
	const density = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "
	var builder strings.Builder
	buffer := bm.Buffer()

	rowStart := 0
	for range bm.Rows {
		for col := range bm.Width {
			pixel := buffer[rowStart+int(col)]
			index := int(float64(pixel) / 0xff * float64(len(density)-1))
			index = len(density) - 1 - index // reverse the index, as density const runs from most to least dense
			densityIndex := density[index]
			builder.WriteByte(densityIndex)
		}
		builder.WriteRune('\n')
		rowStart += int(bm.Pitch)
	}

	return builder.String()
}

// PixelMode is an enumeration type used to describe the format of pixels in a given bitmap.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_pixel_mode
type PixelMode = C.uchar

const (
	PIXEL_MODE_NONE  = PixelMode(C.FT_PIXEL_MODE_NONE)
	PIXEL_MODE_MONO  = PixelMode(C.FT_PIXEL_MODE_MONO)
	PIXEL_MODE_GRAY  = PixelMode(C.FT_PIXEL_MODE_GRAY)
	PIXEL_MODE_GRAY2 = PixelMode(C.FT_PIXEL_MODE_GRAY2)
	PIXEL_MODE_GRAY4 = PixelMode(C.FT_PIXEL_MODE_GRAY4)
	PIXEL_MODE_LCD   = PixelMode(C.FT_PIXEL_MODE_LCD)
	PIXEL_MODE_LCD_V = PixelMode(C.FT_PIXEL_MODE_LCD_V)
	PIXEL_MODE_BGRA  = PixelMode(C.FT_PIXEL_MODE_BGRA)
)

// String returns a formatted representation of the 4 bytes of the PixelMode tag.
func (pixelMode PixelMode) String() string {
	return formatTag(uint32(pixelMode))
}

// GlyphFormat is an enumeration type used to describe the format of a given glyph image.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_glyph_format
type GlyphFormat = C.FT_Glyph_Format

const (
	GLYPH_FORMAT_NONE      = GlyphFormat(C.FT_GLYPH_FORMAT_NONE)
	GLYPH_FORMAT_COMPOSITE = GlyphFormat(C.FT_GLYPH_FORMAT_COMPOSITE)
	GLYPH_FORMAT_BITMAP    = GlyphFormat(C.FT_GLYPH_FORMAT_BITMAP)
	GLYPH_FORMAT_OUTLINE   = GlyphFormat(C.FT_GLYPH_FORMAT_OUTLINE)
	GLYPH_FORMAT_PLOTTER   = GlyphFormat(C.FT_GLYPH_FORMAT_PLOTTER)
	GLYPH_FORMAT_SVG       = GlyphFormat(C.FT_GLYPH_FORMAT_SVG)
)

// String returns a formatted representation of the 4 bytes of the GlyphFormat tag.
func (glyphFormat GlyphFormat) String() string {
	return formatTag(uint32(glyphFormat))
}

// FT_IMAGE_TAG
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_image_tag
