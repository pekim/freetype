package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"fmt"
)

type BBox = C.FT_BBox
type Bool = C.FT_Bool
type Byte = C.FT_Byte
type Bytes = C.FT_Bytes
type Char = C.FT_Char
type F2Dot14 = C.FT_F2Dot14
type F26Dot6 = C.FT_F26Dot6
type Fixed = C.FT_Fixed
type FWord = C.FT_FWord
type Int = C.FT_Int
type Int16 = C.FT_Int16
type Int32 = C.FT_Int32
type Int64 = C.FT_Int64
type Long = C.FT_Long
type Offset = C.FT_Offset
type Pointer = C.FT_Pointer
type Pos = C.FT_Pos
type PtrDist = C.FT_PtrDist
type Short = C.FT_Short
type String = C.FT_String
type Tag = C.FT_Tag
type UFWord = C.FT_UFWord
type UInt = C.FT_UInt
type UInt16 = C.FT_UInt16
type UInt32 = C.FT_UInt32
type UInt64 = C.FT_UInt64
type ULong = C.FT_ULong
type UShort = C.FT_UShort

type GlyphFormat = C.FT_Glyph_Format

const (
	GLYPH_FORMAT_NONE = GlyphFormat(C.FT_GLYPH_FORMAT_NONE)

	GLYPH_FORMAT_COMPOSITE = GlyphFormat(C.FT_GLYPH_FORMAT_COMPOSITE)
	GLYPH_FORMAT_BITMAP    = GlyphFormat(C.FT_GLYPH_FORMAT_BITMAP)
	GLYPH_FORMAT_OUTLINE   = GlyphFormat(C.FT_GLYPH_FORMAT_OUTLINE)
	GLYPH_FORMAT_PLOTTER   = GlyphFormat(C.FT_GLYPH_FORMAT_PLOTTER)
	GLYPH_FORMAT_SVG       = GlyphFormat(C.FT_GLYPH_FORMAT_SVG)
)

func (glyphFormat GlyphFormat) String() string {
	return formatTag(uint32(glyphFormat))
}

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

func (pixelMode PixelMode) String() string {
	return formatTag(uint32(pixelMode))
}

func formatTag(tag uint32) string {
	return fmt.Sprintf("'%s', '%s', '%s', '%s'",
		string(rune(tag>>24&0x000000ff)),
		string(rune(tag>>16&0x000000ff)),
		string(rune(tag>>8&0x000000ff)),
		string(rune(tag>>0&0x000000ff)),
	)
}
