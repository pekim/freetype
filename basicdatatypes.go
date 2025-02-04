package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"fmt"
)

type Byte = C.FT_Byte
type Bytes = C.FT_Bytes
type Char = C.FT_Char
type Int = C.FT_Int
type UInt = C.FT_UInt
type Int16 = C.FT_Int16
type UInt16 = C.FT_UInt16
type Int32 = C.FT_Int32
type UInt32 = C.FT_UInt32
type Int64 = C.FT_Int64
type UInt64 = C.FT_UInt64
type Short = C.FT_Short
type UShort = C.FT_UShort
type Long = C.FT_Long
type ULong = C.FT_ULong
type Bool = C.FT_Bool
type Offset = C.FT_Offset
type PtrDist = C.FT_PtrDist
type String = C.FT_String
type Tag = C.FT_Tag
type Fixed = C.FT_Fixed
type Pointer = C.FT_Pointer
type Pos = C.FT_Pos
type BBox = C.FT_BBox
type FWord = C.FT_FWord
type UFWord = C.FT_UFWord
type F2Dot14 = C.FT_F2Dot14
type UnitVector = C.FT_UnitVector
type F26Dot6 = C.FT_F26Dot6
type Data = C.FT_Data
type Generic_Finalizer = C.FT_Generic_Finalizer
type Bitmap = C.FT_Bitmap
type Pixel_Mode = C.FT_Pixel_Mode

type GlyphFormat C.FT_Glyph_Format

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

func formatTag(tag uint32) string {
	return fmt.Sprintf("'%s', '%s', '%s', '%s'",
		string(rune(tag>>24&0x000000ff)),
		string(rune(tag>>16&0x000000ff)),
		string(rune(tag>>8&0x000000ff)),
		string(rune(tag>>0&0x000000ff)),
	)
}
