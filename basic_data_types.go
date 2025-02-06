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

// Vector is a simple structure used to store a 2D vector; coordinates are of the FT_Pos type.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_vector
type Vector struct {
	// The horizontal coordinate.
	X Pos
	// The vertical coordinate.
	Y Pos
}

// BBox is a structure used to hold an outline's bounding box,
// i.e., the coordinates of its extrema in the horizontal and vertical directions.
//
// The bounding box is specified with the coordinates of the lower left and the upper right corner.
// In PostScript, those values are often called (llx,lly) and (urx,ury), respectively.
//
// If yMin is negative, this value gives the glyph's descender.
// Otherwise, the glyph doesn't descend below the baseline. Similarly, if ymax is positive,
// this value gives the glyph's ascender.
//
// xMin gives the horizontal distance from the glyph's origin to the left edge of the glyph's bounding box.
// If xMin is negative, the glyph extends to the left of the origin.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_bbox
type BBox struct {
	// The horizontal minimum (left-most).
	XMin Pos
	// The vertical minimum (bottom-most).
	YMin Pos
	// The horizontal maximum (right-most).
	XMax Pos
	// The vertical maximum (top-most).
	YMax Pos
}

// Matrix is a simple structure used to store a 2x2 matrix.
// Coefficients are in 16.16 fixed-point format.
// The computation performed is:
//
//	x' = x*xx + y*xy
//	y' = x*yx + y*yy
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

// UnitVector is a simple structure used to store a 2D vector unit vector. Uses F2Dot14 types.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_unitvector
type UnitVector struct {
	// Horizontal coordinate.
	X F2Dot14
	// Vertical coordinate.
	Y F2Dot14
}

// F26Dot6 is a signed 26.6 fixed-point type used for vectorial pixel coordinates.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_f26dot6
type F26Dot6 = C.FT_F26Dot6

// Data is read-only binary data represented as a pointer and a length.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_data
type Data struct {
	// The data.
	Pointer *Byte
	// The length of the data in bytes.
	Length UInt
}

// FT_MAKE_TAG
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_make_tag

/*
Client applications often need to associate their own data to a variety of FreeType core objects.
For example, a text layout API might want to associate a glyph cache to a given size object.

Some FreeType object contains a generic field, of type FT_Generic,
which usage is left to client applications and font servers.

It can be used to store a pointer to client-specific data,
as well as the address of a ‘finalizer’ function, which will be called by FreeType when the object is destroyed
(for example, the previous client example would put the address of the glyph cache destructor in the finalizer field).

https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_generic
*/
type Generic struct {
	// A typeless pointer to any client-specified data.
	// This field is completely ignored by the FreeType library.
	Data unsafe.Pointer
	// A pointer to a ‘generic finalizer’ function, which will be called when the object is destroyed.
	// If this field is set to NULL, no code will be called.
	Finalizer GenericFinalizer
}

// Describe a function used to destroy the ‘client’ data of any FreeType object.
// See the description of the FT_Generic type for details of usage.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_generic_finalizer
type GenericFinalizer func(object unsafe.Pointer)

// Bitmap is a structure used to describe a bitmap or pixmap to the raster.
// Note that we now manage pixmaps of various depths through the pixel_mode field.
//
// Width and Rows refer to the physical size of the bitmap, not the logical one.
// For example, if PixelMode is set to PIXEL_MODE_LCD, the logical width is a just a third of the physical one.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_bitmap
type Bitmap struct {
	// The number of bitmap rows.
	Rows UInt
	// The number of pixels in bitmap row.
	Width UInt
	// The pitch's absolute value is the number of bytes taken by one bitmap row,
	// including padding. However, the pitch is positive when the bitmap has a ‘down’ flow,
	// and negative when it has an ‘up’ flow. In all cases, the pitch is an offset to add to
	// a bitmap pointer in order to go down one row.
	//
	// Note that ‘padding’ means the alignment of a bitmap to a byte border,
	// and FreeType functions normally align to the smallest possible integer value.
	//
	// For the B/W rasterizer, pitch is always an even number.
	//
	// To change the pitch of a bitmap (say, to make it a multiple of 4), use FT_Bitmap_Convert.
	// Alternatively, you might use callback functions to directly render to the application's surface;
	// see the file example2.cpp in the tutorial for a demonstration.
	Pitch Int
	// A typeless pointer to the bitmap buffer.
	// This value should be aligned on 32-bit boundaries in most cases.
	buffer *C.uchar
	// This field is only used with FT_PIXEL_MODE_GRAY;
	// it gives the number of gray levels used in the bitmap.
	NumGrays UShort
	// The pixel mode, i.e., how pixel bits are stored. See PixelMode for possible values.
	PixelMode PixelMode
	// This field is intended for paletted pixel modes; it indicates how the palette is stored.
	// Not used currently.
	palette_mode C.uchar
	// A typeless pointer to the bitmap palette; this field is intended for paletted pixel modes.
	// Not used currently.
	palette unsafe.Pointer
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
			byte_ := buffer[rowStart+int(col)]
			densityIndex := density[byte(len(density))-(byte_/byte(len(density)))-1]
			builder.WriteByte(densityIndex)
		}
		builder.WriteRune('\n')
		rowStart += int(bm.Pitch)
	}

	return builder.String()
}

// PixelMode is an enumeration type used to describe the format of pixels in a given bitmap.
// Note that additional formats may be added in the future.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_pixel_mode
type PixelMode = C.uchar

const (
	// Value 0 is reserved.
	PIXEL_MODE_NONE = PixelMode(C.FT_PIXEL_MODE_NONE)
	// A monochrome bitmap, using 1 bit per pixel.
	// Note that pixels are stored in most-significant order (MSB),
	// which means that the left-most pixel in a byte has value 128.
	PIXEL_MODE_MONO = PixelMode(C.FT_PIXEL_MODE_MONO)
	// An 8-bit bitmap, generally used to represent anti-aliased glyph images.
	// Each pixel is stored in one byte.
	// Note that the number of ‘gray’ levels is stored in the num_grays field of the
	// Bitmap structure (it generally is 256).
	PIXEL_MODE_GRAY = PixelMode(C.FT_PIXEL_MODE_GRAY)
	// A 2-bit per pixel bitmap, used to represent embedded anti-aliased bitmaps in font files
	// according to the OpenType specification.
	// We haven't found a single font using this format, however.
	PIXEL_MODE_GRAY2 = PixelMode(C.FT_PIXEL_MODE_GRAY2)
	// A 4-bit per pixel bitmap, representing embedded anti-aliased bitmaps in font files
	// according to the OpenType specification.
	// We haven't found a single font using this format, however.
	PIXEL_MODE_GRAY4 = PixelMode(C.FT_PIXEL_MODE_GRAY4)
	// An 8-bit bitmap, representing RGB or BGR decimated glyph images used for display on LCD displays;
	// the bitmap is three times wider than the original glyph image.
	// See also RENDER_MODE_LCD.
	PIXEL_MODE_LCD = PixelMode(C.FT_PIXEL_MODE_LCD)
	// An 8-bit bitmap, representing RGB or BGR decimated glyph images used for display on
	// rotated LCD displays; the bitmap is three times taller than the original glyph image.
	// See also RENDER_MODE_LCD_V.
	PIXEL_MODE_LCD_V = PixelMode(C.FT_PIXEL_MODE_LCD_V)
	// [Since 2.5] An image with four 8-bit channels per pixel, representing a color image
	// (such as emoticons) with alpha channel. For each pixel, the format is BGRA, which means,
	// the blue channel comes first in memory. The color channels are pre-multiplied and in the
	// sRGB colorspace. For example, full red at half-translucent opacity will be represented
	// as ‘00,00,80,80’, not ‘00,00,FF,80’.
	// See also LOAD_COLOR.
	PIXEL_MODE_BGRA = PixelMode(C.FT_PIXEL_MODE_BGRA)
)

// String returns a formatted representation of the 4 bytes of the PixelMode tag.
func (pixelMode PixelMode) String() string {
	return formatTag(uint32(pixelMode))
}

// GlyphFormat is an enumeration type used to describe the format of a given glyph image.
// Note that this version of FreeType only supports two image formats,
// even though future font drivers will be able to register their own format.
//
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_glyph_format
type GlyphFormat = C.FT_Glyph_Format

const (
	// The value 0 is reserved.
	GLYPH_FORMAT_NONE = GlyphFormat(C.FT_GLYPH_FORMAT_NONE)
	// The glyph image is a composite of several other images.
	// This format is only used with LOAD_NO_RECURSE, and is used to report
	// compound glyphs (like accented characters).
	GLYPH_FORMAT_COMPOSITE = GlyphFormat(C.FT_GLYPH_FORMAT_COMPOSITE)
	// The glyph image is a bitmap, and can be described as an FT_Bitmap.
	// You generally need to access the bitmap field of the FT_GlyphSlotRec structure to read it.
	GLYPH_FORMAT_BITMAP = GlyphFormat(C.FT_GLYPH_FORMAT_BITMAP)
	// The glyph image is a vectorial outline made of line segments and Bezier arcs;
	// it can be described as an FT_Outline; you generally want to access the outline field
	// of the FT_GlyphSlotRec structure to read it.
	GLYPH_FORMAT_OUTLINE = GlyphFormat(C.FT_GLYPH_FORMAT_OUTLINE)
	// The glyph image is a vectorial path with no inside and outside contours.
	// Some Type 1 fonts, like those in the Hershey family, contain glyphs in this format.
	// These are described as FT_Outline, but FreeType isn't currently capable of rendering them correctly.
	GLYPH_FORMAT_PLOTTER = GlyphFormat(C.FT_GLYPH_FORMAT_PLOTTER)
	// [Since 2.12] The glyph is represented by an SVG document in the ‘SVG ’ table.
	GLYPH_FORMAT_SVG = GlyphFormat(C.FT_GLYPH_FORMAT_SVG)
)

// String returns a formatted representation of the 4 bytes of the GlyphFormat tag.
func (glyphFormat GlyphFormat) String() string {
	return formatTag(uint32(glyphFormat))
}

// FT_IMAGE_TAG
// https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_image_tag
