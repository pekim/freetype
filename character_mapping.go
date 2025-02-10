package freetype

import (
	"modernc.org/libfreetype"
)

// Functions to manage character-to-glyph maps.

/*
CharMap is a handle to a character map (usually abbreviated to ‘charmap’). A charmap is used to translate character codes in a given encoding into glyph indexes for its parent's face. Some font formats may provide several charmaps per font.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_charmap
*/
type CharMap uintptr

// Rec returns a pointer to the FaceRec that is referenced by the Face.
func (charmap CharMap) Rec() *CharMapRec {
	return fromUintptr[CharMapRec](uintptr(charmap))
}

func init() {
	assertSameSize(CharMapRec{}, libfreetype.TFT_CharMapRec{})
}

/*
CharMapRec is the base charmap structure.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_charmaprec
*/
type CharMapRec struct {
	Face       libfreetype.TFT_Face
	Encoding   Encoding
	PlatformID UShort
	EncodingID UShort
}

/*
Encoding is an enumeration to specify character sets supported by charmaps. Used in the FT_Select_Charmap API function.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_encoding
*/
type Encoding libfreetype.TFT_Encoding

var (
	ENCODING_NONE = Encoding(imageTag(0, 0, 0, 0))

	ENCODING_MS_SYMBOL = Encoding(imageTag('s', 'y', 'm', 'b'))
	ENCODING_UNICODE   = Encoding(imageTag('u', 'n', 'i', 'c'))

	ENCODING_SJIS    = Encoding(imageTag('s', 'j', 'i', 's'))
	ENCODING_PRC     = Encoding(imageTag('g', 'b', ' ', ' '))
	ENCODING_BIG5    = Encoding(imageTag('b', 'i', 'g', '5'))
	ENCODING_WANSUNG = Encoding(imageTag('w', 'a', 'n', 's'))
	ENCODING_JOHAB   = Encoding(imageTag('j', 'o', 'h', 'a'))

	/* for backward compatibility */
	ENCODING_GB2312     = ENCODING_PRC
	ENCODING_MS_SJIS    = ENCODING_SJIS
	ENCODING_MS_GB2312  = ENCODING_PRC
	ENCODING_MS_BIG5    = ENCODING_BIG5
	ENCODING_MS_WANSUNG = ENCODING_WANSUNG
	ENCODING_MS_JOHAB   = ENCODING_JOHAB

	ENCODING_ADOBE_STANDARD = Encoding(imageTag('A', 'D', 'O', 'B'))
	ENCODING_ADOBE_EXPERT   = Encoding(imageTag('A', 'D', 'B', 'E'))
	ENCODING_ADOBE_CUSTOM   = Encoding(imageTag('A', 'D', 'B', 'C'))
	ENCODING_ADOBE_LATIN_1  = Encoding(imageTag('l', 'a', 't', '1'))

	ENCODING_OLD_LATIN_2 = Encoding(imageTag('l', 'a', 't', '2'))

	ENCODING_APPLE_ROMAN = Encoding(imageTag('a', 'r', 'm', 'n'))
)

// String returns a formatted representation of the 4 bytes of the Encoding tag.
func (encoding Encoding) String() string {
	return formatTag(uint32(encoding))
}

// // FT_ENC_TAG

/*
SelectCharmap selects a given charmap by its encoding tag.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_select_charmap
*/
func (face Face) SelectCharmap(encoding Encoding) error {
	err := libfreetype.XFT_Select_Charmap(face.tls, face.face, libfreetype.TFT_Encoding(encoding))
	return newError(err, "failed to select charmap for encoding %s (0x%04x)", encoding, int32(encoding))
}

/*
SetCharmap selects a given charmap for character code to glyph index mapping.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_set_charmap
*/
func (face Face) SetCharmap(rec CharMapRec) error {
	err := libfreetype.XFT_Set_Charmap(face.tls, face.face, toUintptr(&rec))
	return newError(err, "failed to set charmap")
}

/*
GetCharmapIndex retrieves the index of a given charmap.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_get_charmap_index
*/
func (face Face) GetCharmapIndex(charmap CharMap) Int {
	return libfreetype.XFT_Get_Charmap_Index(face.tls, libfreetype.TFT_CharMap(charmap))
}

/*
GetCharIndex returns the glyph index of a given character code.
This function uses the currently selected charmap to do the mapping.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_get_char_index
*/
func (face Face) GetCharIndex(charcode rune) UInt {
	return libfreetype.XFT_Get_Char_Index(face.tls, face.face, libfreetype.TFT_ULong(charcode))
}

// /*
// GetFirstChar returns the first character code in the current charmap of a given face, together with its corresponding glyph index.

// https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_get_first_char
// */
// func (face Face) GetFirstChar() (ULong, UInt) {
// 	var gindex UInt
// 	charCode := C.FT_Get_First_Char(face.face, &gindex)
// 	return charCode, gindex
// }

// /*
// GetNextChar returns the next character code in the current charmap of a given face following the value char_code, as well as the corresponding glyph index.

// https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_get_next_char
// */
// func (face Face) GetNextChar(charCode ULong) (ULong, UInt) {
// 	var gindex UInt
// 	nextCharCode := C.FT_Get_Next_Char(face.face, charCode, &gindex)
// 	return nextCharCode, gindex
// }

// /*
// LoadChar loads a glyph into the glyph slot of a face object, accessed by its character code.

// https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_load_char
// */
// func (face Face) LoadChar(charCode rune, loadFlags LoadFlag) error {
// 	err := C.FT_Load_Char(face.face, ULong(charCode), loadFlags)
// 	return newError(err, "failed to load char '%s' (0x%04x) with flags 0x%04x", string(charCode), charCode, loadFlags)
// }
