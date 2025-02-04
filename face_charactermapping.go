package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"fmt"
)

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

func (encoding Encoding) String() string {
	return fmt.Sprintf("'%s', '%s', '%s', '%s'",
		string(rune(encoding>>24&0x000000ff)),
		string(rune(encoding>>16&0x000000ff)),
		string(rune(encoding>>8&0x000000ff)),
		string(rune(encoding>>0&0x000000ff)),
	)
}

func (face Face) SelectCharmap(encoding Encoding) error {
	err := C.FT_Select_Charmap(face.face, encoding)
	return newError(err, "failed to select charmap for encoding %s (0x%04x)", encoding, int32(encoding))
}

func (face Face) SetCharmap(charmap CharMapRec) error {
	ftCharmap := toPointer[C.FT_CharMapRec](charmap)
	ftCharmap.face = charmap.Face.face
	err := C.FT_Set_Charmap(face.face, ftCharmap)
	return newError(err, "failed to set charmap")
}

func GetCharmapIndex(charmap CharMapRec) Int {
	ftCharmap := toPointer[C.FT_CharMapRec](charmap)
	ftCharmap.face = charmap.Face.face
	return C.FT_Get_Charmap_Index(ftCharmap)
}

func (face Face) GetCharIndex(charcode rune) UInt {
	return C.FT_Get_Char_Index(face.face, ULong(charcode))
}

func (face Face) GetFirstChar() (ULong, UInt) {
	var gindex UInt
	charCode := C.FT_Get_First_Char(face.face, &gindex)
	return charCode, gindex
}

func (face Face) GetNextChar(charCode ULong) (ULong, UInt) {
	var gindex UInt
	nextCharCode := C.FT_Get_Next_Char(face.face, charCode, &gindex)
	return nextCharCode, gindex
}
