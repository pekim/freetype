package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"unsafe"
)

// Functions to manage character-to-glyph maps.

/*
CharMap is a handle to a character map (usually abbreviated to ‘charmap’). A charmap is used to translate character codes in a given encoding into glyph indexes for its parent's face. Some font formats may provide several charmaps per font.

Each face object owns zero or more charmaps, but only one of them can be ‘active’, providing the data used by GetCharIndex or LoadChar.

The list of available charmaps in a face is available through Face.Charmaps.

The currently active charmap is available as Face.Charmap. You should call SetCharmap to change it.

When a new face is created (either through Library.NewFace or Library.OpenFace), the library looks for a Unicode charmap within the list and automatically activates it. If there is no Unicode charmap, FreeType doesn't set an ‘active’ charmap.

See CharMapRec for the publicly accessible fields of a given character map.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_charmap
*/
type CharMap *CharMapRec

func init() {
	assertSameSize(CharMapRec{}, C.FT_CharMapRec{})
}

/*
CharMapRec is the base charmap structure.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_charmaprec
*/
type CharMapRec struct {
	// A handle to the parent face object.
	Face Face
	// An Encoding tag identifying the charmap. Use this with SelectCharmap.
	Encoding Encoding
	// An ID number describing the platform for the following encoding ID.
	// This comes directly from the TrueType specification and gets emulated for other formats.
	PlatformID UShort
	// A platform-specific encoding number.
	// This also comes from the TrueType specification and gets emulated similarly.
	EncodingID UShort
}

/*
Encoding is an enumeration to specify character sets supported by charmaps. Used in the FT_Select_Charmap API function.

Despite the name, this enumeration lists specific character repertoires (i.e., charsets), and not text encoding methods (e.g., UTF-8, UTF-16, etc.).

Other encodings might be defined in the future.

When loading a font, FreeType makes a Unicode charmap active if possible (either if the font provides such a charmap, or if FreeType can synthesize one from PostScript glyph name dictionaries; in either case, the charmap is tagged with ENCODING_UNICODE). If such a charmap is synthesized, it is placed at the first position of the charmap array.

All other encodings are considered legacy and tagged only if explicitly defined in the font file. Otherwise, ENCODING_NONE is used.

ENCODING_NONE is set by the BDF and PCF drivers if the charmap is neither Unicode nor ISO-8859-1 (otherwise it is set to ENCODING_UNICODE). Use FT_Get_BDF_Charset_ID to find out which encoding is really present. If, for example, the cs_registry field is ‘KOI8’ and the cs_encoding field is ‘R’, the font is encoded in KOI8-R.

ENCODING_NONE is always set (with a single exception) by the winfonts driver. Use FT_Get_WinFNT_Header and examine the charset field of the FT_WinFNT_HeaderRec structure to find out which encoding is really present. For example, FT_WinFNT_ID_CP1251 (204) means Windows code page 1251 (for Russian).

ENCODING_NONE is set if platform_id is TT_PLATFORM_MACINTOSH and encoding_id is not TT_MAC_ID_ROMAN (otherwise it is set to ENCODING_APPLE_ROMAN).

If platform_id is TT_PLATFORM_MACINTOSH, use the function FT_Get_CMap_Language_ID to query the Mac language ID that may be needed to be able to distinguish Apple encoding variants. See https://www.unicode.org/Public/MAPPINGS/VENDORS/APPLE/Readme.txt
to get an idea how to do that. Basically, if the language ID is 0, don't use it, otherwise subtract 1 from the language ID. Then examine encoding_id. If, for example, encoding_id is TT_MAC_ID_ROMAN and the language ID (minus 1) is TT_MAC_LANGID_GREEK, it is the Greek encoding, not Roman. TT_MAC_ID_ARABIC with TT_MAC_LANGID_FARSI means the Farsi variant of the Arabic encoding.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_encoding
*/
type Encoding = C.FT_Encoding

const (
	ENCODING_NONE = Encoding(C.FT_ENCODING_NONE)

	ENCODING_MS_SYMBOL = Encoding(C.FT_ENCODING_MS_SYMBOL)
	ENCODING_UNICODE   = Encoding(C.FT_ENCODING_UNICODE)

	ENCODING_SJIS    = Encoding(C.FT_ENCODING_SJIS)
	ENCODING_PRC     = Encoding(C.FT_ENCODING_PRC)
	ENCODING_BIG5    = Encoding(C.FT_ENCODING_BIG5)
	ENCODING_WANSUNG = Encoding(C.FT_ENCODING_WANSUNG)
	ENCODING_JOHAB   = Encoding(C.FT_ENCODING_JOHAB)

	/* for backward compatibility */
	ENCODING_GB2312     = Encoding(C.FT_ENCODING_GB2312)
	ENCODING_MS_SJIS    = Encoding(C.FT_ENCODING_MS_SJIS)
	ENCODING_MS_GB2312  = Encoding(C.FT_ENCODING_MS_GB2312)
	ENCODING_MS_BIG5    = Encoding(C.FT_ENCODING_MS_BIG5)
	ENCODING_MS_WANSUNG = Encoding(C.FT_ENCODING_MS_WANSUNG)
	ENCODING_MS_JOHAB   = Encoding(C.FT_ENCODING_MS_JOHAB)

	ENCODING_ADOBE_STANDARD = Encoding(C.FT_ENCODING_ADOBE_STANDARD)
	ENCODING_ADOBE_EXPERT   = Encoding(C.FT_ENCODING_ADOBE_EXPERT)
	ENCODING_ADOBE_CUSTOM   = Encoding(C.FT_ENCODING_ADOBE_CUSTOM)
	ENCODING_ADOBE_LATIN_1  = Encoding(C.FT_ENCODING_ADOBE_LATIN_1)

	ENCODING_OLD_LATIN_2 = Encoding(C.FT_ENCODING_OLD_LATIN_2)

	ENCODING_APPLE_ROMAN = Encoding(C.FT_ENCODING_APPLE_ROMAN)
)

// String returns a formatted representation of the 4 bytes of the Encoding tag.
func (encoding Encoding) String() string {
	return formatTag(uint32(encoding))
}

// FT_ENC_TAG

/*
SelectCharmap select a given charmap by its encoding tag.

  - encoding - A handle to the selected encoding.

This function returns an error if no charmap in the face corresponds to the encoding queried here.

Because many fonts contain more than a single cmap for Unicode encoding, this function has some special code to select the one that covers Unicode best (‘best’ in the sense that a UCS-4 cmap is preferred to a UCS-2 cmap). It is thus preferable to SetCharmap in this case.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_select_charmap
*/
func (face Face) SelectCharmap(encoding Encoding) error {
	err := C.FT_Select_Charmap(face.face, encoding)
	return newError(err, "failed to select charmap for encoding %s (0x%04x)", encoding, int32(encoding))
}

/*
SetCharmap selects a given charmap for character code to glyph index mapping.

  - charmap - A handle to the selected charmap.

This function returns an error if the charmap is not part of the face (i.e., if it is not listed in the face->charmaps table).

It also fails if an OpenType type 14 charmap is selected (which doesn't map character codes to glyph indices at all).

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_set_charmap
*/
func (face Face) SetCharmap(charmap CharMap) error {
	ftCharmap := toPointer[C.FT_CharMapRec](charmap)
	ftCharmap.face = charmap.Face.face
	err := C.FT_Set_Charmap(face.face, ftCharmap)
	return newError(err, "failed to set charmap")
}

/*
GetCharmapIndex retrieves the index of a given charmap.

  - charmap - A handle to a charmap.

Returns the index into the array of character maps within the face to which charmap belongs. If an error occurs, -1 is returned.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_get_charmap_index
*/
func GetCharmapIndex(charmap CharMap) Int {
	return C.FT_Get_Charmap_Index((C.FT_CharMap)(unsafe.Pointer(charmap)))
}

/*
GetCharIndex returns the glyph index of a given character code. This function uses the currently selected charmap to do the mapping.

  - charcode - The character code.

Returns the glyph index. 0 means ‘undefined character code’.

If you use FreeType to manipulate the contents of font files directly, be aware that the glyph index returned by this function doesn't always correspond to the internal indices used within the file. This is done to ensure that value 0 always corresponds to the ‘missing glyph’. If the first glyph is not named ‘.notdef’, then for Type 1 and Type 42 fonts, ‘.notdef’ will be moved into the glyph ID 0 position, and whatever was there will be moved to the position ‘.notdef’ had. For Type 1 fonts, if there is no ‘.notdef’ glyph at all, then one will be created at index 0 and whatever was there will be moved to the last index – Type 42 fonts are considered invalid under this condition.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_get_char_index
*/
func (face Face) GetCharIndex(charcode rune) UInt {
	return C.FT_Get_Char_Index(face.face, ULong(charcode))
}

/*
GetFirstChar returns the first character code in the current charmap of a given face, together with its corresponding glyph index.

Returns the charmap's first character code, and the glyph index of first character code (0 if charmap is empty).

You should use this function together with GetNextChar to parse all character codes available in a given charmap. The code should look like this:

	FT_ULong  charcode;
	FT_UInt   gindex;


	charcode = FT_Get_First_Char( face, &gindex );
	while ( gindex != 0 )
	{
	  ... do something with (charcode,gindex) pair ...

	  charcode = FT_Get_Next_Char( face, charcode, &gindex );
	}

Be aware that character codes can have values up to 0xFFFFFFFF; this might happen for non-Unicode or malformed cmaps. However, even with regular Unicode encoding, so-called ‘last resort fonts’ (using SFNT cmap format 13, see function FT_Get_CMap_Format) normally have entries for all Unicode characters up to 0x1FFFFF, which can cause a lot of iterations.

Note that *agindex is set to 0 if the charmap is empty. The result itself can be 0 in two cases: if the charmap is empty or if the value 0 is the first valid character code.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_get_first_char
*/
func (face Face) GetFirstChar() (ULong, UInt) {
	var gindex UInt
	charCode := C.FT_Get_First_Char(face.face, &gindex)
	return charCode, gindex
}

/*
GetNextChar returns the next character code in the current charmap of a given face following the value char_code, as well as the corresponding glyph index.

  - char_code - The starting character code.

Returns the charmap's next character code, and the glyph index of next character code (0 if charmap is empty).

You should use this function with GetFirstChar to walk over all character codes available in a given charmap. See the note for that function for a simple code example.

Note that the character code returned is set to 0 when there are no more codes in the charmap.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_get_next_char
*/
func (face Face) GetNextChar(charCode ULong) (ULong, UInt) {
	var gindex UInt
	nextCharCode := C.FT_Get_Next_Char(face.face, charCode, &gindex)
	return nextCharCode, gindex
}

/*
LoadChar loads a glyph into the glyph slot of a face object, accessed by its character code.

  - char_code - The glyph's character code, according to the current charmap used in the face.
  - load_flags - A flag indicating what to load for this glyph. The LOAD_XXX constants can be used to control the glyph loading process (e.g., whether the outline should be scaled, whether to load bitmaps or not, whether to hint the outline, etc).

This function simply calls GetCharIndex and LoadGlyph.

Many fonts contain glyphs that can't be loaded by this function since its glyph indices are not listed in any of the font's charmaps.

If no active cmap is set up (i.e., face->charmap is zero), the call to GetCharIndex is omitted, and the function behaves identically to LoadGlyph.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_load_char
*/
func (face Face) LoadChar(charCode rune, loadFlags LoadFlag) error {
	err := C.FT_Load_Char(face.face, ULong(charCode), loadFlags)
	return newError(err, "failed to load char '%s' (0x%04x) with flags 0x%04x", string(charCode), charCode, loadFlags)
}
