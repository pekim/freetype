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
GetNameIndex returns the glyph index of a given glyph name. This only works for those faces where HasGlyphNames returns true.

  - glyphName - The glyph name.

Returns the glyph index. 0 means ‘undefined character code’.

Acceptable glyph names might come from the Adobe Glyph List. See GetGlyphName for the inverse functionality.

This function has limited capabilities if the config macro FT_CONFIG_OPTION_POSTSCRIPT_NAMES is not defined in ftoption.h: It then works only for fonts that actually embed glyph names (which many recent OpenType fonts do not).

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_name_index
*/
func (face Face) GetNameIndex(glyphName string) UInt {
	cName := C.CString(glyphName)
	defer C.free(unsafe.Pointer(cName))

	return C.FT_Get_Name_Index(face.face, cName)
}

/*
GetGlyphName retrieves the ASCII name of a given glyph in a face. This only works for those faces where HasGlyphNames returns true.

  - glyphIndex - The glyph index.

An error is returned if the face doesn't provide glyph names or if the glyph index is invalid. In all cases of failure, the returned name will be an empty string.

The glyph name is truncated to fit within the buffer if it is too long.

Be aware that FreeType reorders glyph indices internally so that glyph index 0 always corresponds to the ‘missing glyph’ (called ‘.notdef’).

This function has limited capabilities if the config macro FT_CONFIG_OPTION_POSTSCRIPT_NAMES is not defined in ftoption.h: It then works only for fonts that actually embed glyph names (which many recent OpenType fonts do not).

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_glyph_name
*/
func (face Face) GetGlyphName(glyphIndex UInt) (string, error) {
	buffer := make([]C.char, 128)
	err := C.FT_Get_Glyph_Name(face.face, glyphIndex, C.FT_Pointer(unsafe.Pointer(&buffer[0])), UInt(len(buffer)))
	name := C.GoString(&buffer[0])
	return name, newError(err, "failed to get glyph name for glyph index %d", glyphIndex)
}

/*
GetPostscriptName retrieves the ASCII PostScript name of a given face, if available. This only works with PostScript, TrueType, and OpenType fonts.

Returns the face's PostScript name, or an empty string if unavailable.

For variation fonts, this string changes if you select a different instance, and you have to call GetPostScript_Name again to retrieve it. FreeType follows Adobe TechNote #5902, ‘Generating PostScript Names for Fonts Using OpenType Font Variations’.

https://download.macromedia.com/pub/developer/opentype/tech-notes/5902.AdobePSNameGeneration.html

[Since 2.9] Special PostScript names for named instances are only returned if the named instance is set with FT_Set_Named_Instance (and the font has corresponding entries in its ‘fvar’ table or is the default named instance). If IsVariation returns true, the algorithmically derived PostScript name is provided, not looking up special entries for named instances.

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

Use this function rather than directly reading the fs_type field in the PS_FontInfoRec structure, which is only guaranteed to return the correct results for Type 1 fonts.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_fstype_flags
*/
func (face Face) GetFSTypeFlags() FSType {
	return C.FT_Get_FSType_Flags(face.face)
}

/*
FSType is a list of bit flags used in the fsType field of the OS/2 table in a TrueType or OpenType font and the FSType entry in a PostScript font. These bit flags are returned by GetFSTypeFlags; they inform client applications of embedding and subsetting restrictions associated with a font.

See https://www.adobe.com/content/dam/Adobe/en/devnet/acrobat/pdfs/FontPolicies.pdf for more details.

The flags are ORed together, thus more than a single value can be returned.

While the FSType flags can indicate that a font may be embedded, a license with the font vendor may be separately required to use the font in this way.

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_fstype_xxx
*/
type FSType = UShort

const (
	// Fonts with no fsType bit set may be embedded and permanently installed on the remote
	// system by an application.
	FSTYPE_INSTALLABLE_EMBEDDING = FSType(C.FT_FSTYPE_INSTALLABLE_EMBEDDING)
	// Fonts that have only this bit set must not be modified, embedded or exchanged in any manner
	// without first obtaining permission of the font software copyright owner.
	FSTYPE_RESTRICTED_LICENSE_EMBEDDING = FSType(C.FT_FSTYPE_RESTRICTED_LICENSE_EMBEDDING)
	// The font may be embedded and temporarily loaded on the remote system. Documents containing
	// Preview & Print fonts must be opened ‘read-only’; no edits can be applied to the document.
	FSTYPE_PREVIEW_AND_PRINT_EMBEDDING = FSType(C.FT_FSTYPE_PREVIEW_AND_PRINT_EMBEDDING)
	// The font may be embedded but must only be installed temporarily on other systems. In
	// contrast to Preview & Print fonts, documents containing editable fonts may be opened for
	// reading, editing is permitted, and changes may be saved.
	FSTYPE_EDITABLE_EMBEDDING = FSType(C.FT_FSTYPE_EDITABLE_EMBEDDING)
	// The font may not be subsetted prior to embedding.
	FSTYPE_NO_SUBSETTING = FSType(C.FT_FSTYPE_NO_SUBSETTING)
	// Only bitmaps contained in the font may be embedded; no outline data may be embedded.
	// If there are no bitmaps available in the font, then the font is unembeddable.
	FSTYPE_BITMAP_EMBEDDING_ONLY = FSType(C.FT_FSTYPE_BITMAP_EMBEDDING_ONLY)
)

/*
GetSubGlyphInfo retrieves a description of a given subglyph. Only use it if glyph->format is GLYPH_FORMAT_COMPOSITE; an error is returned otherwise.

  - glyph - The source glyph slot.
  - subIndex - The index of the subglyph. Must be less than glyph->num_subglyphs.

returned
  - The glyph index of the subglyph.
  - The subglyph flags, see FT_SUBGLYPH_FLAG_XXX.
  - The subglyph's first argument (if any).
  - The subglyph's second argument (if any).
  - The subglyph transformation (if any).

The return values for first argument, second argument, and transformation must be interpreted depending on the flags returned. See the OpenType specification for details.

https://docs.microsoft.com/en-us/typography/opentype/spec/glyf#composite-glyph-description

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_get_subglyph_info
*/
func (face Face) GetSubGlyphInfo(glyph *GlyphSlotRec, subIndex UInt) (Int, SUBGLYPH_FLAG, Int, Int, Matrix, error) {
	var index C.FT_Int
	var flags C.FT_UInt
	var arg1 C.FT_Int
	var arg2 C.FT_Int
	var transform C.FT_Matrix
	err := C.FT_Get_SubGlyph_Info(
		(*C.FT_GlyphSlotRec)(unsafe.Pointer(glyph)),
		subIndex, &index, &flags, &arg1, &arg2, &transform)
	return Int(index), SUBGLYPH_FLAG(flags), Int(arg1), Int(arg2), to[Matrix](transform),
		newError(err, "failed to get sub glyph info")
}

/*
SUBGLYPH_FLAG is a list of constants describing subglyphs. Please refer to the ‘glyf’ table description in the OpenType specification for the meaning of the various flags (which get synthesized for non-OpenType subglyphs).

https://docs.microsoft.com/en-us/typography/opentype/spec/glyf#composite-glyph-description

https://freetype.org/freetype2/docs/reference/ft2-information_retrieval.html#ft_subglyph_flag_xxx
*/
type SUBGLYPH_FLAG = UInt

const (
	SUBGLYPH_FLAG_ARGS_ARE_WORDS     = SUBGLYPH_FLAG(C.FT_SUBGLYPH_FLAG_ARGS_ARE_WORDS)
	SUBGLYPH_FLAG_ARGS_ARE_XY_VALUES = SUBGLYPH_FLAG(C.FT_SUBGLYPH_FLAG_ARGS_ARE_XY_VALUES)
	SUBGLYPH_FLAG_ROUND_XY_TO_GRID   = SUBGLYPH_FLAG(C.FT_SUBGLYPH_FLAG_ROUND_XY_TO_GRID)
	SUBGLYPH_FLAG_SCALE              = SUBGLYPH_FLAG(C.FT_SUBGLYPH_FLAG_SCALE)
	SUBGLYPH_FLAG_XY_SCALE           = SUBGLYPH_FLAG(C.FT_SUBGLYPH_FLAG_XY_SCALE)
	SUBGLYPH_FLAG_2X2                = SUBGLYPH_FLAG(C.FT_SUBGLYPH_FLAG_2X2)
	SUBGLYPH_FLAG_USE_MY_METRICS     = SUBGLYPH_FLAG(C.FT_SUBGLYPH_FLAG_USE_MY_METRICS)
)
