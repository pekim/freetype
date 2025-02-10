package freetype

import (
	"modernc.org/libfreetype"
)

// Access the names embedded in TrueType and OpenType files.

// SFNTName is structure used to model an SFNT ‘name’ table entry.
//
// https://freetype.org/freetype2/docs/reference/ft2-sfnt_names.html#ft_sfntname
type SFNTName struct {
	PlatformID UShort
	EncodingID UShort
	LanguageID UShort
	NameID     UShort

	string     *Byte /* this string is *not* null-terminated! */
	string_len UInt  /* in bytes                              */
}

func init() {
	assertSameSize(SFNTName{}, libfreetype.TFT_SfntName{})
}

// String returns the table entry's string field.
func (sn SFNTName) String() string {
	return goStringForNotNullTerminatedCString(sn.string, sn.string_len)
}

// GetSFNTNameCount retrieves the number of name strings in the SFNT ‘name’ table.
//
// https://freetype.org/freetype2/docs/reference/ft2-sfnt_names.html#ft_get_sfnt_name_count
func (face Face) GetSFNTNameCount() int {
	return int(libfreetype.XFT_Get_Sfnt_Name_Count(face.tls, face.face))
}

func (face Face) GetSFNTName(index UInt) (SFNTName, error) {
	var sfntName SFNTName
	err := libfreetype.XFT_Get_Sfnt_Name(face.tls, face.face, index, toUintptr(&sfntName))
	return sfntName, newError(err, "failed to get SFNT name table with index %d", index)
}

// SFNTLangTag is a structure to model a language tag entry from an SFNT ‘name’ table.
//
// https://freetype.org/freetype2/docs/reference/ft2-sfnt_names.html#ft_sfntlangtag
type SFNTLangTag struct {
	string     *Byte /* this string is *not* null-terminated! */
	string_len UInt  /* in bytes                              */
}

func init() {
	assertSameSize(SFNTLangTag{}, libfreetype.TFT_SfntLangTag{})
}

// String returns the language tag as a string.
func (sn SFNTLangTag) String() string {
	return goStringForNotNullTerminatedCString(sn.string, sn.string_len)
}

// GetSFNTLangTag retrieves the language tag associated with a language ID of an SFNT ‘name’ table entry.
//
// https://freetype.org/freetype2/docs/reference/ft2-sfnt_names.html#ft_get_sfnt_langtag
func (face Face) GetSFNTLangTag(langID UInt) (SFNTLangTag, error) {
	var langTag SFNTLangTag
	err := libfreetype.XFT_Get_Sfnt_LangTag(face.tls, face.face, langID, toUintptr(&langTag))
	return langTag, newError(err, "failed to get SFNT language tage with langID %d", langID)
}
