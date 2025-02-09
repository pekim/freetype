package freetype

import "modernc.org/libfreetype"

// Functions to manage character-to-glyph maps.

/*
GetCharIndex returns the glyph index of a given character code.
This function uses the currently selected charmap to do the mapping.

https://freetype.org/freetype2/docs/reference/ft2-character_mapping.html#ft_get_char_index
*/
func (face Face) GetCharIndex(charcode rune) uint {
	return uint(libfreetype.XFT_Get_Char_Index(face.tls, face.face, libfreetype.TFT_ULong(charcode)))
}
