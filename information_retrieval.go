package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
//
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

// Functions to retrieve font and glyph information.

/*
GetNameIndex returns the glyph index of a given glyph name.
This only works for those faces where HasGlyphNames returns true.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_name_index
*/
func (face Face) GetNameIndex(glyphName string) UInt {
	cName := C.CString(glyphName)
	defer C.free(unsafe.Pointer(cName))

	return C.FT_Get_Name_Index(face.face, cName)
}

/*
GetGlyphName retrieves the ASCII name of a given glyph in a face.
This only works for those faces where HasGlyphNames returns true.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_glyph_name
*/
func (face Face) GetGlyphName(glyphIndex UInt) (string, error) {
	buffer := make([]C.char, 128)
	err := C.FT_Get_Glyph_Name(face.face, glyphIndex, C.FT_Pointer(unsafe.Pointer(&buffer[0])), UInt(len(buffer)))
	name := C.GoString(&buffer[0])
	return name, newError(err, "failed to get glyph name for glyph index %d", glyphIndex)
}

/*
GetPostscriptName retrieves the ASCII PostScript name of a given face, if available.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_postscript_name
*/
func (face Face) GetPostscriptName() string {
	cName := C.FT_Get_Postscript_Name(face.face)
	if cName == nil {
		return ""
	}
	return C.GoString(cName)
}

/*
GetFSTypeFlags returns the FSType flags for a font.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_fstype_flags
*/
func (face Face) GetFSTypeFlags() FSType {
	return C.FT_Get_FSType_Flags(face.face)
}

/*
FSType is a list of bit flags used in the fsType field of the OS/2 table in a TrueType or OpenType font
and the FSType entry in a PostScript font.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_fstype_xxx
*/
type FSType = UShort

const (
	FSTYPE_INSTALLABLE_EMBEDDING        = FSType(C.FT_FSTYPE_INSTALLABLE_EMBEDDING)
	FSTYPE_RESTRICTED_LICENSE_EMBEDDING = FSType(C.FT_FSTYPE_RESTRICTED_LICENSE_EMBEDDING)
	FSTYPE_PREVIEW_AND_PRINT_EMBEDDING  = FSType(C.FT_FSTYPE_PREVIEW_AND_PRINT_EMBEDDING)
	FSTYPE_EDITABLE_EMBEDDING           = FSType(C.FT_FSTYPE_EDITABLE_EMBEDDING)
	FSTYPE_NO_SUBSETTING                = FSType(C.FT_FSTYPE_NO_SUBSETTING)
	FSTYPE_BITMAP_EMBEDDING_ONLY        = FSType(C.FT_FSTYPE_BITMAP_EMBEDDING_ONLY)
)

/*
GetSubGlyphInfo retrieves a description of a given subglyph. Only use it if glyph->format is GLYPH_FORMAT_COMPOSITE; an error is returned otherwise.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_subglyph_info
*/
func (face Face) GetSubGlyphInfo(glyph *GlyphSlotRec, subIndex UInt) (Int, SubglyphFlag, Int, Int, Matrix, error) {
	var index C.FT_Int
	var flags C.FT_UInt
	var arg1 C.FT_Int
	var arg2 C.FT_Int
	var transform C.FT_Matrix
	err := C.FT_Get_SubGlyph_Info(
		(*C.FT_GlyphSlotRec)(unsafe.Pointer(glyph)),
		subIndex, &index, &flags, &arg1, &arg2, &transform)
	return Int(index), SubglyphFlag(flags), Int(arg1), Int(arg2), to[Matrix](transform),
		newError(err, "failed to get sub glyph info")
}

/*
SubglyphFlag is a list of constants describing subglyphs.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_subglyph_flag_xxx
*/
type SubglyphFlag = UInt

const (
	SUBGLYPH_FLAG_ARGS_ARE_WORDS     = SubglyphFlag(C.FT_SUBGLYPH_FLAG_ARGS_ARE_WORDS)
	SUBGLYPH_FLAG_ARGS_ARE_XY_VALUES = SubglyphFlag(C.FT_SUBGLYPH_FLAG_ARGS_ARE_XY_VALUES)
	SUBGLYPH_FLAG_ROUND_XY_TO_GRID   = SubglyphFlag(C.FT_SUBGLYPH_FLAG_ROUND_XY_TO_GRID)
	SUBGLYPH_FLAG_SCALE              = SubglyphFlag(C.FT_SUBGLYPH_FLAG_SCALE)
	SUBGLYPH_FLAG_XY_SCALE           = SubglyphFlag(C.FT_SUBGLYPH_FLAG_XY_SCALE)
	SUBGLYPH_FLAG_2X2                = SubglyphFlag(C.FT_SUBGLYPH_FLAG_2X2)
	SUBGLYPH_FLAG_USE_MY_METRICS     = SubglyphFlag(C.FT_SUBGLYPH_FLAG_USE_MY_METRICS)
)
