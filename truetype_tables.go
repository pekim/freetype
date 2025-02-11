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
type Sfnt_Tag = libfreetype.TFT_Sfnt_Tag

const (
	SFNT_HEAD = Sfnt_Tag(0)
	SFNT_MAXP = Sfnt_Tag(1)
	SFNT_OS2  = Sfnt_Tag(2)
	SFNT_HHEA = Sfnt_Tag(3)
	SFNT_VHEA = Sfnt_Tag(4)
	SFNT_POST = Sfnt_Tag(5)
	SFNT_PCLT = Sfnt_Tag(6)
	SFNT_MAX  = Sfnt_Tag(7)
)

// GetSFNTTable returns a pointer to a given SFNT table stored within a face.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#ft_get_sfnt_table
func (face Face) GetSFNTTable(tag Sfnt_Tag) (unsafe.Pointer, error) {
	table := libfreetype.XFT_Get_Sfnt_Table(face.tls, face.face, tag)
	if table == 0 {
		return nil, fmt.Errorf("failed to get SFNT table with tag %d", tag)
	}
	return *(*unsafe.Pointer)(unsafe.Pointer(&table)), nil
}

// FT_Load_Sfnt_Table
//
//

// FT_Sfnt_Table_Info
//
//

// FT_Get_CMap_Language_ID
//
//

// FT_Get_CMap_Format
//
//

// FT_PARAM_TAG_UNPATENTED_HINTING
//
//

// TT_PLATFORM_XXX
//
//

// TT_APPLE_ID_XXX
//
//

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

// TTName is the possible values of the ‘name’ identifier field in the name records of an SFNT ‘name’ table.
// These values are platform independent.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_name_id_xxx
type TTName = UInt

const (
	NAME_ID_COPYRIGHT      = TTName(0)
	NAME_ID_FONT_FAMILY    = TTName(1)
	NAME_ID_FONT_SUBFAMILY = TTName(2)
	NAME_ID_UNIQUE_ID      = TTName(3)
	NAME_ID_FULL_NAME      = TTName(4)
	NAME_ID_VERSION_STRING = TTName(5)
	NAME_ID_PS_NAME        = TTName(6)
	NAME_ID_TRADEMARK      = TTName(7)

	/* the following values are from the OpenType spec */
	NAME_ID_MANUFACTURER = TTName(8)
	NAME_ID_DESIGNER     = TTName(9)
	NAME_ID_DESCRIPTION  = TTName(10)
	NAME_ID_VENDOR_URL   = TTName(11)
	NAME_ID_DESIGNER_URL = TTName(12)
	NAME_ID_LICENSE      = TTName(13)
	NAME_ID_LICENSE_URL  = TTName(14)
	/* number 15 is reserved */
	NAME_ID_TYPOGRAPHIC_FAMILY    = TTName(16)
	NAME_ID_TYPOGRAPHIC_SUBFAMILY = TTName(17)
	NAME_ID_MAC_FULL_NAME         = TTName(18)

	/* The following code is new as of 2000-01-21 */
	NAME_ID_SAMPLE_TEXT = TTName(19)

	/* This is new in OpenType 1.3 */
	NAME_ID_CID_FINDFONT_NAME = TTName(20)

	/* This is new in OpenType 1.5 */
	NAME_ID_WWS_FAMILY    = TTName(21)
	NAME_ID_WWS_SUBFAMILY = TTName(22)

	/* This is new in OpenType 1.7 */
	NAME_ID_LIGHT_BACKGROUND = TTName(23)
	NAME_ID_DARK_BACKGROUND  = TTName(24)

	/* This is new in OpenType 1.8 */
	NAME_ID_VARIATIONS_PREFIX = TTName(25)
)

// TT_UCR_XXX
//
//
