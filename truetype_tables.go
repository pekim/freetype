package freetype

import "modernc.org/libfreetype"

// TrueType-specific table types and functions.

// TTHeader is a structure to model a TrueType font header table.
// All fields follow the OpenType specification.
// The 64-bit timestamps are stored in two-element arrays Created and Modified, first the upper then the lower 32 bits.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_header
type TTHeader struct {
	Table_Version Fixed
	Font_Revision Fixed

	CheckSum_Adjust Long
	Magic_Number    Long

	Flags        UShort
	Units_Per_EM UShort

	_ [2]ULong // created
	_ [2]ULong // modified

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

// func (header TTHeader)Created()[2]ULong{

// }

func init() {
	assertSameSize(TTHeader{}, libfreetype.TTT_Header{})
}

// TT_HoriHeader
//
//

// TT_VertHeader
//
//

// TT_OS2
//
//

// TT_Postscript
//
//

// TT_PCLT
//
//

// TT_MaxProfile
//
//

// FT_Sfnt_Tag
//
//

// FT_Get_Sfnt_Table
//
//

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
