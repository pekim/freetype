package freetype

import (
	"math"
	"strings"
	"unsafe"

	"modernc.org/libfreetype"
)

// The basic data types defined by the library.

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_byte
type Byte = libfreetype.TFT_Byte

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_bytes
type Bytes = libfreetype.TFT_Bytes

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_char
type Char = libfreetype.TFT_Char

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_int
type Int = libfreetype.TFT_Int

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_uint
type UInt = libfreetype.TFT_UInt

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_int16
type Int16 = libfreetype.TFT_Int16

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_uint16
type UInt16 = libfreetype.TFT_UInt16

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_int32
type Int32 = libfreetype.TFT_Int32

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_uint32
type UInt32 = libfreetype.TFT_UInt32

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_int64
type Int64 = libfreetype.TFT_Int64

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_uint64
type UInt64 = libfreetype.TFT_UInt64

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_short
type Short = libfreetype.TFT_Short

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_ushort
type UShort = libfreetype.TFT_UShort

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_long
type Long = libfreetype.TFT_Long

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_ulong
type ULong = libfreetype.TFT_ULong

// Bool is a type of unsigned char, used for simple booleans.
// As usual, values 1 and 0 represent true and false, respectively.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_bool
type Bool = libfreetype.TFT_Bool

// Offset is equivalent to the ANSI C size_t type,
// i.e., the largest unsigned integer type used to express a file size or position, or a memory block size.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_offset
type Offset = libfreetype.TFT_Offset

// PtrDist is equivalent to the ANSI C ptrdiff_t type,
// i.e., the largest signed integer type used to express the distance between two pointers.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_ptrdist
type PtrDist = libfreetype.TFT_PtrDist

// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_string
type String = libfreetype.TFT_String

// Tag is type for 32-bit tags (as used in the SFNT format).
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_tag
type Tag = libfreetype.TFT_Tag

// For FT_Error see error.go file.

// Fixed is used to store 16.16 fixed-point values, like scaling values or matrix coefficients.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_fixed
type Fixed = libfreetype.TFT_Fixed

// Pointer is a simple type for a typeless pointer.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_pointer
type Pointer = libfreetype.TFT_Pointer

// Pos is used to store vectorial coordinates.
// Depending on the context, these can represent distances in integer font units,
// or 16.16, or 26.6 fixed-point pixel coordinates.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_pos
type Pos = libfreetype.TFT_Pos

func init() {
	assertSameSize(Vector{}, libfreetype.TFT_Vector{})
}

// Vector is a simple structure used to store a 2D vector; coordinates are of the FT_Pos type.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_vector
type Vector struct {
	X Pos
	Y Pos
}

func init() {
	assertSameSize(BBox{}, libfreetype.TFT_BBox{})
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
	assertSameSize(Matrix{}, libfreetype.TFT_Matrix{})
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
type FWord = libfreetype.TFT_FWord

// UWord is an unsigned 16-bit integer used to store a distance in original font units.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_uword
type UFWord = libfreetype.TFT_UFWord

// F2Dot14 is a signed 2.14 fixed-point type used for unit vectors.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_f2dot14
type F2Dot14 = libfreetype.TFT_F2Dot14

func init() {
	assertSameSize(UnitVector{}, libfreetype.TFT_UnitVector{})
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
type F26Dot6 = libfreetype.TFT_F26Dot6

func init() {
	assertSameSize(Data{}, libfreetype.TFT_Data{})
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
	assertSameSize(Generic{}, libfreetype.TFT_Generic{})
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
	assertSameSize(Bitmap{}, libfreetype.TFT_Bitmap{})
}

// Bitmap is a structure used to describe a bitmap or pixmap to the raster.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_bitmap
type Bitmap struct {
	Rows      uint32
	Width     uint32
	Pitch     int32
	buffer    unsafe.Pointer
	NumGrays  uint16
	PixelMode PixelMode
	_         uint8          // palette_mode
	_         unsafe.Pointer //palette
}

// Buffer returns the Bitmap's buffer as byte slice.
func (bm Bitmap) Buffer() []byte {
	return unsafe.Slice((*byte)(bm.buffer), bm.Rows*UInt(math.Abs(float64(bm.Pitch))))
}

// BufferVisualization returns a string containing an ascii representation of a grayscale bitmap.
// It can be useful for demonstration or debug purposes.
func (bm Bitmap) BufferVisualization() string {
	const density = " .'`^\",:;Il!i><~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
	var builder strings.Builder
	buffer := bm.Buffer()

	rowStart := 0
	for range bm.Rows {
		for col := range bm.Width {
			pixel := buffer[rowStart+int(col)]
			index := int(float64(pixel) / 0xff * float64(len(density)-1))
			builder.WriteByte(density[index])
		}
		builder.WriteRune('\n')
		rowStart += int(bm.Pitch)
	}

	return builder.String()
}

// PixelMode is an enumeration type used to describe the format of pixels in a given bitmap.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_pixel_mode
type PixelMode uint8

const (
	PIXEL_MODE_NONE  = PixelMode(0)
	PIXEL_MODE_MONO  = PixelMode(1)
	PIXEL_MODE_GRAY  = PixelMode(2)
	PIXEL_MODE_GRAY2 = PixelMode(3)
	PIXEL_MODE_GRAY4 = PixelMode(4)
	PIXEL_MODE_LCD   = PixelMode(5)
	PIXEL_MODE_LCD_V = PixelMode(6)
	PIXEL_MODE_BGRA  = PixelMode(7)
)

// String returns a formatted representation of the 4 bytes of the PixelMode tag.
func (pixelMode PixelMode) String() string {
	return formatTag(uint32(pixelMode))
}

// GlyphFormat is an enumeration type used to describe the format of a given glyph image.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_glyph_format
type GlyphFormat libfreetype.TFT_Glyph_Format

var (
	GLYPH_FORMAT_NONE      = GlyphFormat(imageTag(0, 0, 0, 0))
	GLYPH_FORMAT_COMPOSITE = GlyphFormat(imageTag('c', 'o', 'm', 'p'))
	GLYPH_FORMAT_BITMAP    = GlyphFormat(imageTag('b', 'i', 't', 's'))
	GLYPH_FORMAT_OUTLINE   = GlyphFormat(imageTag('o', 'u', 't', 'l'))
	GLYPH_FORMAT_PLOTTER   = GlyphFormat(imageTag('p', 'l', 'o', 't'))
	GLYPH_FORMAT_SVG       = GlyphFormat(imageTag('S', 'V', 'G', ' '))
)

// String returns a formatted representation of the 4 bytes of the GlyphFormat tag.
func (glyphFormat GlyphFormat) String() string {
	return formatTag(uint32(glyphFormat))
}

// imageTag converts four-letter tags to an unsigned long type.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_image_tag
func imageTag(x1, x2, x3, x4 byte) uint32 {
	return 0 |
		uint32(x1)<<24 |
		uint32(x2)<<16 |
		uint32(x3)<<8 |
		uint32(x4)<<0
}
