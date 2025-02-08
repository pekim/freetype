package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
//
// #include <freetype/ttnameid.h>
import "C"

// TTName is the possible values of the ‘name’ identifier field in the name records of an SFNT ‘name’ table.
// These values are platform independent.
//
// https://freetype.org/freetype2/docs/reference/ft2-truetype_tables.html#tt_name_id_xxx
type TTName = UInt

const (
	NAME_ID_COPYRIGHT      = TTName(C.TT_NAME_ID_COPYRIGHT)
	NAME_ID_FONT_FAMILY    = TTName(C.TT_NAME_ID_FONT_FAMILY)
	NAME_ID_FONT_SUBFAMILY = TTName(C.TT_NAME_ID_FONT_SUBFAMILY)
	NAME_ID_UNIQUE_ID      = TTName(C.TT_NAME_ID_UNIQUE_ID)
	NAME_ID_FULL_NAME      = TTName(C.TT_NAME_ID_FULL_NAME)
	NAME_ID_VERSION_STRING = TTName(C.TT_NAME_ID_VERSION_STRING)
	NAME_ID_PS_NAME        = TTName(C.TT_NAME_ID_PS_NAME)
	NAME_ID_TRADEMARK      = TTName(C.TT_NAME_ID_TRADEMARK)

	/* the following values are from the OpenType spec */
	NAME_ID_MANUFACTURER = TTName(C.TT_NAME_ID_MANUFACTURER)
	NAME_ID_DESIGNER     = TTName(C.TT_NAME_ID_DESIGNER)
	NAME_ID_DESCRIPTION  = TTName(C.TT_NAME_ID_DESCRIPTION)
	NAME_ID_VENDOR_URL   = TTName(C.TT_NAME_ID_VENDOR_URL)
	NAME_ID_DESIGNER_URL = TTName(C.TT_NAME_ID_DESIGNER_URL)
	NAME_ID_LICENSE      = TTName(C.TT_NAME_ID_LICENSE)
	NAME_ID_LICENSE_URL  = TTName(C.TT_NAME_ID_LICENSE_URL)
	/* number 15 is reserved */
	NAME_ID_TYPOGRAPHIC_FAMILY    = TTName(C.TT_NAME_ID_TYPOGRAPHIC_FAMILY)
	NAME_ID_TYPOGRAPHIC_SUBFAMILY = TTName(C.TT_NAME_ID_TYPOGRAPHIC_SUBFAMILY)
	NAME_ID_MAC_FULL_NAME         = TTName(C.TT_NAME_ID_MAC_FULL_NAME)

	/* The following code is new as of 2000-01-21 */
	NAME_ID_SAMPLE_TEXT = TTName(C.TT_NAME_ID_SAMPLE_TEXT)

	/* This is new in OpenType 1.3 */
	NAME_ID_CID_FINDFONT_NAME = TTName(C.TT_NAME_ID_CID_FINDFONT_NAME)

	/* This is new in OpenType 1.5 */
	NAME_ID_WWS_FAMILY    = TTName(C.TT_NAME_ID_WWS_FAMILY)
	NAME_ID_WWS_SUBFAMILY = TTName(C.TT_NAME_ID_WWS_SUBFAMILY)

	/* This is new in OpenType 1.7 */
	NAME_ID_LIGHT_BACKGROUND = TTName(C.TT_NAME_ID_LIGHT_BACKGROUND)
	NAME_ID_DARK_BACKGROUND  = TTName(C.TT_NAME_ID_DARK_BACKGROUND)

	/* This is new in OpenType 1.8 */
	NAME_ID_VARIATIONS_PREFIX = TTName(C.TT_NAME_ID_VARIATIONS_PREFIX)
)
