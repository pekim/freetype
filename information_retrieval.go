package freetype

import (
	"modernc.org/libc"
	"modernc.org/libfreetype"
)

// Functions to retrieve font and glyph information.

/*
GetNameIndex returns the glyph index of a given glyph name.
This only works for those faces where HasGlyphNames returns true.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_name_index
*/
func (face Face) GetNameIndex(glyphName string) (UInt, error) {
	cName, err := libc.CString(glyphName)
	if err != nil {
		return 0, err
	}
	defer libc.Xfree(nil, cName)

	return libfreetype.XFT_Get_Name_Index(face.tls, face.face, cName), nil
}

/*
GetGlyphName retrieves the ASCII name of a given glyph in a face.
This only works for those faces where HasGlyphNames returns true.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_glyph_name
*/
func (face Face) GetGlyphName(glyphIndex UInt) (string, error) {
	buffer := make([]byte, 128)
	err := libfreetype.XFT_Get_Glyph_Name(face.tls, face.face, glyphIndex,
		toUintptr(&buffer[0]), UInt(len(buffer)))
	name := libc.GoString(toUintptr(&buffer[0]))
	return name, newError(err, "failed to get glyph name for glyph index %d", glyphIndex)
}

/*
GetPostscriptName retrieves the ASCII PostScript name of a given face, if available.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_postscript_name
*/
func (face Face) GetPostscriptName() string {
	cName := libfreetype.XFT_Get_Postscript_Name(face.tls, face.face)
	if cName == 0 {
		return ""
	}
	return libc.GoString(cName)
}

/*
GetFSTypeFlags returns the FSType flags for a font.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_fstype_flags
*/
func (face Face) GetFSTypeFlags() FSType {
	return libfreetype.XFT_Get_FSType_Flags(face.tls, face.face)
}

/*
FSType is a list of bit flags used in the fsType field of the OS/2 table in a TrueType or OpenType font
and the FSType entry in a PostScript font.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_fstype_xxx
*/
type FSType = UShort

const (
	FSTYPE_INSTALLABLE_EMBEDDING        = FSType(0x0000)
	FSTYPE_RESTRICTED_LICENSE_EMBEDDING = FSType(0x0002)
	FSTYPE_PREVIEW_AND_PRINT_EMBEDDING  = FSType(0x0004)
	FSTYPE_EDITABLE_EMBEDDING           = FSType(0x0008)
	FSTYPE_NO_SUBSETTING                = FSType(0x0100)
	FSTYPE_BITMAP_EMBEDDING_ONLY        = FSType(0x0200)
)

/*
GetSubGlyphInfo retrieves a description of a given subglyph. Only use it if glyph->format is GLYPH_FORMAT_COMPOSITE; an error is returned otherwise.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_subglyph_info
*/
func (face Face) GetSubGlyphInfo(glyph *GlyphSlotRec, subIndex UInt) (Int, SubglyphFlag, Int, Int, Matrix, error) {
	var index Int
	var flags UInt
	var arg1 Int
	var arg2 Int
	var transform Matrix
	err := libfreetype.XFT_Get_SubGlyph_Info(
		face.tls,
		libfreetype.TFT_GlyphSlot(toUintptr(glyph)),
		subIndex, toUintptr(&index), toUintptr(&flags), toUintptr(&arg1), toUintptr(&arg2), toUintptr(&transform))
	return Int(index), SubglyphFlag(flags), Int(arg1), Int(arg2), transform,
		newError(err, "failed to get sub glyph info")
}

/*
SubglyphFlag is a list of constants describing subglyphs.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_subglyph_flag_xxx
*/
type SubglyphFlag = UInt

const (
	SUBGLYPH_FLAG_ARGS_ARE_WORDS     = SubglyphFlag(1)
	SUBGLYPH_FLAG_ARGS_ARE_XY_VALUES = SubglyphFlag(2)
	SUBGLYPH_FLAG_ROUND_XY_TO_GRID   = SubglyphFlag(4)
	SUBGLYPH_FLAG_SCALE              = SubglyphFlag(8)
	SUBGLYPH_FLAG_XY_SCALE           = SubglyphFlag(0x40)
	SUBGLYPH_FLAG_2X2                = SubglyphFlag(0x80)
	SUBGLYPH_FLAG_USE_MY_METRICS     = SubglyphFlag(0x200)
)
