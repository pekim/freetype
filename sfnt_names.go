package freetype

import (
	"unsafe"

	"modernc.org/libfreetype"
)

// Access the names embedded in TrueType and OpenType files.

// SfntName is structure used to model an SFNT ‘name’ table entry.
//
// https://freetype.org/freetype2/docs/reference/ft2-sfnt_names.html#ft_sfntname
type SfntName struct {
	PlatformID UShort
	EncodingID UShort
	LanguageID UShort
	NameID     UShort

	string     *Byte /* this string is *not* null-terminated! */
	string_len UInt  /* in bytes                              */
}

func init() {
	assertSameSize(SfntName{}, libfreetype.TFT_SfntName{})
}

// String returns the table entry's string field.
func (sn SfntName) String() string {
	return unsafe.String(sn.string, sn.string_len)
}

// GetSfntNameCount retrieves the number of name strings in the SFNT ‘name’ table.
//
// https://freetype.org/freetype2/docs/reference/ft2-sfnt_names.html#ft_get_sfnt_name_count
func (face Face) GetSfntNameCount() int {
	return int(libfreetype.XFT_Get_Sfnt_Name_Count(face.tls, face.face))
}

func (face Face) GetSfntName(index UInt) (SfntName, error) {
	var sfntName SfntName
	err := libfreetype.XFT_Get_Sfnt_Name(face.tls, face.face, index, toUintptr(&sfntName))
	return sfntName, newError(err, "failed to get SFNT name table with index %d", index)
}

// SfntLangTag is a structure to model a language tag entry from an SFNT ‘name’ table.
//
// https://freetype.org/freetype2/docs/reference/ft2-sfnt_names.html#ft_sfntlangtag
type SfntLangTag struct {
	string     *Byte /* this string is *not* null-terminated! */
	string_len UInt  /* in bytes                              */
}

func init() {
	assertSameSize(SfntLangTag{}, libfreetype.TFT_SfntLangTag{})
}

// String returns the language tag as a string.
func (sn SfntLangTag) String() string {
	return unsafe.String(sn.string, sn.string_len)
}

// GetSfntLangTag retrieves the language tag associated with a language ID of an SFNT ‘name’ table entry.
//
// https://freetype.org/freetype2/docs/reference/ft2-sfnt_names.html#ft_get_sfnt_langtag
func (face Face) GetSfntLangTag(langID UInt) (SfntLangTag, error) {
	var langTag SfntLangTag
	err := libfreetype.XFT_Get_Sfnt_LangTag(face.tls, face.face, langID, toUintptr(&langTag))
	return langTag, newError(err, "failed to get SFNT language tage with langID %d", langID)
}
