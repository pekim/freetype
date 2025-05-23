package freetype

import (
	"fmt"
	"unsafe"

	"modernc.org/libfreetype"
)

// TrueType-specific table types and functions.

func init() {
	assertSameSize(TT_Header{}, libfreetype.TTT_Header{})
}

// TT_Header is a structure to model a TrueType font header table.
// All fields follow the OpenType specification.
// The 64-bit timestamps are stored in two-element arrays Created and Modified, first the upper then the lower 32 bits.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_header
type TT_Header struct {
	TableVersion Fixed
	FontRevision Fixed

	CheckSumAdjust Long
	MagicNumber    Long

	Flags      UShort
	UnitsPerEM UShort

	Created   [2]ULong
	Modifield [2]ULong

	XMin Short
	YMin Short
	XMax Short
	YMax Short

	MacStyle      UShort
	LowestRecPPEM UShort

	FontDirection    Short
	IndexToLocFormat Short
	GlyphDataFormat  Short
}

func init() {
	assertSameSize(TT_HoriHeader{}, libfreetype.TTT_HoriHeader{})
}

// TT_HoriHeader is a structure to model a TrueType horizontal header, the ‘hhea’ table, as well as the
// corresponding horizontal metrics table, ‘hmtx’.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_horiheader
type TT_HoriHeader struct {
	Version   Fixed
	Ascender  Short
	Descender Short
	Line_Gap  Short

	AdvanceWidthMax UShort /* advance width maximum */

	MinLeftSideBearing  Short /* minimum left-sb       */
	MinRightSideBearing Short /* minimum right-sb      */
	XMax_Extent         Short /* xmax extents          */
	CaretSlopeRise      Short
	CaretSlopeRun       Short
	CaretOffset         Short

	Reserved [4]Short

	MetricDataFormat Short
	NumberOfHMetrics UShort

	/* The following fields are not defined by the OpenType specification */
	/* but they are used to connect the metrics header to the relevant    */
	/* 'hmtx' table.                                                      */

	LongMetrics  unsafe.Pointer
	ShortMetrics unsafe.Pointer
}

func init() {
	assertSameSize(TT_VertHeader{}, libfreetype.TTT_VertHeader{})
}

// TT_VertHeader is a structure used to model a TrueType vertical header, the ‘vhea’ table, as well
// as the corresponding vertical metrics table, ‘vmtx’.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_vertheader
type TT_VertHeader struct {
	Version   Fixed
	Ascender  Short
	Descender Short
	Line_Gap  Short

	AdvanceHeightMax UShort /* advance height maximum */

	MinTopSideBearing    Short /* minimum top-sb          */
	MinBottomSideBearing Short /* minimum bottom-sb       */
	YMax_Extent          Short /* ymax extents            */
	CaretSlopeRise       Short
	CaretSlopeRun        Short
	CaretOffset          Short

	Reserved [4]Short

	MetricDataFormat Short
	NumberOfVMetrics UShort

	/* The following fields are not defined by the OpenType specification */
	/* but they are used to connect the metrics header to the relevant    */
	/* 'vmtx' table.                                                      */

	LongMetrics  unsafe.Pointer
	ShortMetrics unsafe.Pointer
}

func init() {
	assertSameSize(TT_OS2{}, libfreetype.TTT_OS2{})
}

// TT_OS2 is a structure to model a TrueType ‘OS/2’ table.
// All fields comply to the OpenType specification.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_os2
type TT_OS2 struct {
	Version             UShort /* 0x0001 - more or 0xFFFF */
	XAvgCharWidth       Short
	UsWeightClass       UShort
	UssWidthClass       UShort
	FsType              UShort
	YSubscriptXSize     Short
	YSubscriptYSize     Short
	YSubscriptXOffset   Short
	YSubscriptYOffset   Short
	YSuperscriptXSize   Short
	YSuperscriptYSize   Short
	YSuperscriptXOffset Short
	YSuperscriptYOffset Short
	YStrikeoutSize      Short
	YStrikeoutPosition  Short
	SFamilyClass        Short

	Panose [10]Byte

	UlUnicodeRange1 ULong /* Bits 0-31   */
	UlUnicodeRange2 ULong /* Bits 32-63  */
	UlUnicodeRange3 ULong /* Bits 64-95  */
	UlUnicodeRange4 ULong /* Bits 96-127 */

	achVendID [4]Char

	FsSelection      UShort
	UsFirstCharIndex UShort
	UsLastCharIndex  UShort
	STypoAscender    Short
	STypoDescender   Short
	STypoLineGap     Short
	UsWinAscent      UShort
	UsWinDescent     UShort

	/* only version 1 and higher: */

	UlCodePageRange1 ULong /* Bits 0-31   */
	UlCodePageRange2 ULong /* Bits 32-63  */

	/* only version 2 and higher: */

	SxHeight       Short
	SCapHeight     Short
	USsDefaultChar UShort
	USsBreakChar   UShort
	USsMaxContext  UShort

	/* only version 5 and higher: */

	UsLowerOpticalPointSize UShort /* in twips (1/20 points) */
	UsUpperOpticalPointSize UShort /* in twips (1/20 points) */
}

func init() {
	assertSameSize(TT_Postscript{}, libfreetype.TTT_Postscript{})
}

// TT_Postscript is a structure to model a TrueType ‘post’ table.
// All fields comply to the OpenType specification.
// This structure does not reference a font's PostScript glyph names; use GetGlyphName to retrieve them.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_postscript
type TT_Postscript struct {
	FormatType         Fixed
	ItalicAngle        Fixed
	UnderlinePosition  Short
	UnderlineThickness Short
	IsFixedPitch       ULong
	MinMemType42       ULong
	MaxMemType42       ULong
	MinMemType1        ULong
	MaxMemType1        ULong

	/* Glyph names follow in the 'post' table, but we don't */
	/* load them by default.                                */
}

func init() {
	assertSameSize(TT_PCLT{}, libfreetype.TTT_PCLT{})
}

// TT_PCLT is a structure to model a TrueType ‘PCLT’ table.
// All fields comply to the OpenType specification.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_pclt

type TT_PCLT struct {
	Version             Fixed
	FontNumber          ULong
	Pitch               UShort
	XHeight             UShort
	Style               UShort
	TypeFamily          UShort
	CapHeight           UShort
	SymbolSet           UShort
	TypeFace            [16]Char
	CharacterComplement [8]Char
	FileName            [6]Char
	StrokeWeight        Char
	WidthType           Char
	SerifStyle          Byte
	Reserved            Byte
}

func init() {
	assertSameSize(TT_MaxProfile{}, libfreetype.TTT_MaxProfile{})
}

// TT_MaxProfile is the maximum profile (‘maxp’) table contains many max values,
// which can be used to pre-allocate arrays for speeding up glyph loading and hinting.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_maxprofile
type TT_MaxProfile struct {
	Version               Fixed
	NumGlyphs             UShort
	MaxPoints             UShort
	MaxContours           UShort
	MaxCompositePoints    UShort
	MaxCompositeContours  UShort
	MaxZones              UShort
	MaxTwilightPoints     UShort
	MaxStorage            UShort
	MaxFunctionDefs       UShort
	MaxInstructionDefs    UShort
	MaxStackElements      UShort
	MaxSizeOfInstructions UShort
	MaxComponentElements  UShort
	MaxComponentDepth     UShort
}

// FT_Sfnt_Tag is an enumeration to specify indices of SFNT tables loaded and parsed by FreeType during
// initialization of an SFNT font. Used in the FT_Get_Sfnt_Table API function.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#ft_sfnt_tag
type SfntTag = libfreetype.TFT_Sfnt_Tag

const (
	SFNT_HEAD = SfntTag(0)
	SFNT_MAXP = SfntTag(1)
	SFNT_OS2  = SfntTag(2)
	SFNT_HHEA = SfntTag(3)
	SFNT_VHEA = SfntTag(4)
	SFNT_POST = SfntTag(5)
	SFNT_PCLT = SfntTag(6)
	SFNT_MAX  = SfntTag(7)
)

// GetSfntTable returns a pointer to a given SFNT table stored within a face.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#ft_get_sfnt_table
func (face Face) GetSfntTable(tag SfntTag) (unsafe.Pointer, error) {
	table := libfreetype.XFT_Get_Sfnt_Table(face.tls, face.face, tag)
	if table == 0 {
		return nil, fmt.Errorf("failed to get SFNT table with tag %d", tag)
	}
	return *(*unsafe.Pointer)(unsafe.Pointer(&table)), nil
}

// LoadSfntTable loads any SFNT font table into client memory.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#ft_load_sfnt_table
func (face Face) LoadSfntTable(tag uint32, offset Long, buffer []byte, length *ULong) error {
	var buffer_ uintptr
	if buffer != nil {
		buffer_ = toUintptr(&buffer[0])
	}
	err := libfreetype.XFT_Load_Sfnt_Table(face.tls, face.face, ULong(tag), Long(offset), buffer_, toUintptr(length))
	return newError(err, "failed to load sfnt table %s, with offset %d", formatTag(tag), offset)
}

// SfntTableInfo returns information on an SFNT table.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#ft_sfnt_table_info
func (face Face) SfntTableInfo(tableIndex UInt, tag *ULong) (ULong, error) {
	var length ULong
	err := libfreetype.XFT_Sfnt_Table_Info(face.tls, face.face, UInt(tableIndex), toUintptr(tag), toUintptr(&length))
	if tag == nil {
		return length, newError(err, "failed to get sfnt table info count")
	}
	return length, newError(err, "failed to get sfnt table info for index %d, tag %s", tableIndex, formatTag(uint32(*tag)))
}

// GetCMapLanguageID returns cmap language ID as specified in the OpenType standard.
// Definitions of language ID values are in file FT_TRUETYPE_IDS_H.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#ft_get_cmap_language_id
func (face Face) GetCMapLanguageID(charmap CharMap) ULong {
	return libfreetype.XFT_Get_CMap_Language_ID(face.tls, libfreetype.TFT_CharMap(charmap))
}

// GetCMapFormat returns the format of an SFNT ‘cmap’ table.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#ft_get_cmap_format
func (face Face) GetCMapFormat(charmap CharMap) Long {
	return libfreetype.XFT_Get_CMap_Format(face.tls, libfreetype.TFT_CharMap(charmap))
}

// FT_PARAM_TAG_UNPATENTED_HINTING
//
//

// TT_PLatform is list of valid values for the platform_id identifier code in FT_CharMapRec and FT_SfntName structures.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_platform_xxx
type TT_Platform UShort

const (
	PLATFORM_APPLE_UNICODE = TT_Platform(0)
	PLATFORM_MACINTOSH     = TT_Platform(1)
	PLATFORM_ISO           = TT_Platform(2) /* deprecated */
	PLATFORM_MICROSOFT     = TT_Platform(3)
	PLATFORM_CUSTOM        = TT_Platform(4)
	PLATFORM_ADOBE         = TT_Platform(7) /* artificial */
)

// TT_AppleID is a list of valid values for the encoding_id for TT_PLATFORM_APPLE_UNICODE charmaps and name entries.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_apple_id_xxx
type TT_AppleID UShort

const (
	APPLE_ID_DEFAULT          = TT_AppleID(0) /* Unicode 1.0                   */
	APPLE_ID_UNICODE_1_1      = TT_AppleID(1) /* specify Hangul at U+34xx      */
	APPLE_ID_ISO_10646        = TT_AppleID(2) /* deprecated                    */
	APPLE_ID_UNICODE_2_0      = TT_AppleID(3) /* or later                      */
	APPLE_ID_UNICODE_32       = TT_AppleID(4) /* 2.0 or later, full repertoire */
	APPLE_ID_VARIANT_SELECTOR = TT_AppleID(5) /* variation selector data       */
	APPLE_ID_FULL_UNICODE     = TT_AppleID(6) /* used with type 13 cmaps       */
)

// TT_MAC_ID_XXX
//
//

// TT_ISO_ID_XXX
//
//

// TT_MS_ID_XXX
//
//

// TT_ADOBE_ID_XXX
//
//

// TT_MAC_LANGID_XXX
//
//

// TT_MS_LANGID_XXX
//
//

// TT_NameID is the possible values of the ‘name’ identifier field in the name records of an SFNT ‘name’ table.
// These values are platform independent.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_name_id_xxx
type TT_NameID = UInt

const (
	NAME_ID_COPYRIGHT      = TT_NameID(0)
	NAME_ID_FONT_FAMILY    = TT_NameID(1)
	NAME_ID_FONT_SUBFAMILY = TT_NameID(2)
	NAME_ID_UNIQUE_ID      = TT_NameID(3)
	NAME_ID_FULL_NAME      = TT_NameID(4)
	NAME_ID_VERSION_STRING = TT_NameID(5)
	NAME_ID_PS_NAME        = TT_NameID(6)
	NAME_ID_TRADEMARK      = TT_NameID(7)

	/* the following values are from the OpenType spec */
	NAME_ID_MANUFACTURER = TT_NameID(8)
	NAME_ID_DESIGNER     = TT_NameID(9)
	NAME_ID_DESCRIPTION  = TT_NameID(10)
	NAME_ID_VENDOR_URL   = TT_NameID(11)
	NAME_ID_DESIGNER_URL = TT_NameID(12)
	NAME_ID_LICENSE      = TT_NameID(13)
	NAME_ID_LICENSE_URL  = TT_NameID(14)
	/* number 15 is reserved */
	NAME_ID_TYPOGRAPHIC_FAMILY    = TT_NameID(16)
	NAME_ID_TYPOGRAPHIC_SUBFAMILY = TT_NameID(17)
	NAME_ID_MAC_FULL_NAME         = TT_NameID(18)

	/* The following code is new as of 2000-01-21 */
	NAME_ID_SAMPLE_TEXT = TT_NameID(19)

	/* This is new in OpenType 1.3 */
	NAME_ID_CID_FINDFONT_NAME = TT_NameID(20)

	/* This is new in OpenType 1.5 */
	NAME_ID_WWS_FAMILY    = TT_NameID(21)
	NAME_ID_WWS_SUBFAMILY = TT_NameID(22)

	/* This is new in OpenType 1.7 */
	NAME_ID_LIGHT_BACKGROUND = TT_NameID(23)
	NAME_ID_DARK_BACKGROUND  = TT_NameID(24)

	/* This is new in OpenType 1.8 */
	NAME_ID_VARIATIONS_PREFIX = TT_NameID(25)
)

// TT_UCR_XXX
//
//
