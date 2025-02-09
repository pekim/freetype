package freetype

import (
	"fmt"

	"modernc.org/libfreetype"
)

/*
FTError is the FreeType error code type. A value of 0 (Err_Ok) is always interpreted as a successful operation.

https://freetype.org/freetype2/docs/reference/ft2-basic_types.html#ft_error
*/
type FTError = libfreetype.TFT_Error

const (
	Err_Ok                    = FTError(0x00)
	Err_Cannot_Open_Resource  = FTError(0x01)
	Err_Unknown_File_Format   = FTError(0x02)
	Err_Invalid_File_Format   = FTError(0x03)
	Err_Invalid_Version       = FTError(0x04)
	Err_Lower_Module_Version  = FTError(0x05)
	Err_Invalid_Argument      = FTError(0x06)
	Err_Unimplemented_Feature = FTError(0x07)
	Err_Invalid_Table         = FTError(0x08)
	Err_Invalid_Offset        = FTError(0x09)
	Err_Array_Too_Large       = FTError(0x0A)
	Err_Missing_Module        = FTError(0x0B)
	Err_Missing_Property      = FTError(0x0C)

	/* glyph/character errors */

	Err_Invalid_Glyph_Index    = FTError(0x10)
	Err_Invalid_Character_Code = FTError(0x11)
	Err_Invalid_Glyph_Format   = FTError(0x12)
	Err_Cannot_Render_Glyph    = FTError(0x13)
	Err_Invalid_Outline        = FTError(0x14)
	Err_Invalid_Composite      = FTError(0x15)
	Err_Too_Many_Hints         = FTError(0x16)
	Err_Invalid_Pixel_Size     = FTError(0x17)
	Err_Invalid_SVG_Document   = FTError(0x18)

	/* handle errors */

	Err_Invalid_Handle         = FTError(0x20)
	Err_Invalid_Library_Handle = FTError(0x21)
	Err_Invalid_Driver_Handle  = FTError(0x22)
	Err_Invalid_Face_Handle    = FTError(0x23)
	Err_Invalid_Size_Handle    = FTError(0x24)
	Err_Invalid_Slot_Handle    = FTError(0x25)
	Err_Invalid_CharMap_Handle = FTError(0x26)
	Err_Invalid_Cache_Handle   = FTError(0x27)
	Err_Invalid_Stream_Handle  = FTError(0x28)

	/* driver errors */

	Err_Too_Many_Drivers    = FTError(0x30)
	Err_Too_Many_Extensions = FTError(0x31)

	/* memory errors */

	Err_Out_Of_Memory   = FTError(0x40)
	Err_Unlisted_Object = FTError(0x41)

	/* stream errors */

	Err_Cannot_Open_Stream       = FTError(0x51)
	Err_Invalid_Stream_Seek      = FTError(0x52)
	Err_Invalid_Stream_Skip      = FTError(0x53)
	Err_Invalid_Stream_Read      = FTError(0x54)
	Err_Invalid_Stream_Operation = FTError(0x55)
	Err_Invalid_Frame_Operation  = FTError(0x56)
	Err_Nested_Frame_Access      = FTError(0x57)
	Err_Invalid_Frame_Read       = FTError(0x58)

	/* raster errors */

	Err_Raster_Uninitialized   = FTError(0x60)
	Err_Raster_Corrupted       = FTError(0x61)
	Err_Raster_Overflow        = FTError(0x62)
	Err_Raster_Negative_Height = FTError(0x63)

	/* cache errors */

	Err_Too_Many_Caches = FTError(0x70)

	/* TrueType and SFNT errors */

	Err_Invalid_Opcode            = FTError(0x80)
	Err_Too_Few_Arguments         = FTError(0x81)
	Err_Stack_Overflow            = FTError(0x82)
	Err_Code_Overflow             = FTError(0x83)
	Err_Bad_Argument              = FTError(0x84)
	Err_Divide_By_Zero            = FTError(0x85)
	Err_Invalid_Reference         = FTError(0x86)
	Err_Debug_OpCode              = FTError(0x87)
	Err_ENDF_In_Exec_Stream       = FTError(0x88)
	Err_Nested_DEFS               = FTError(0x89)
	Err_Invalid_CodeRange         = FTError(0x8A)
	Err_Execution_Too_Long        = FTError(0x8B)
	Err_Too_Many_Function_Defs    = FTError(0x8C)
	Err_Too_Many_Instruction_Defs = FTError(0x8D)
	Err_Table_Missing             = FTError(0x8E)
	Err_Horiz_Header_Missing      = FTError(0x8F)
	Err_Locations_Missing         = FTError(0x90)
	Err_Name_Table_Missing        = FTError(0x91)
	Err_CMap_Table_Missing        = FTError(0x92)
	Err_Hmtx_Table_Missing        = FTError(0x93)
	Err_Post_Table_Missing        = FTError(0x94)
	Err_Invalid_Horiz_Metrics     = FTError(0x95)
	Err_Invalid_CharMap_Format    = FTError(0x96)
	Err_Invalid_PPem              = FTError(0x97)
	Err_Invalid_Vert_Metrics      = FTError(0x98)
	Err_Could_Not_Find_Context    = FTError(0x99)
	Err_Invalid_Post_Table_Format = FTError(0x9A)
	Err_Invalid_Post_Table        = FTError(0x9B)
	Err_DEF_In_Glyf_Bytecode      = FTError(0x9C)
	Err_Missing_Bitmap            = FTError(0x9D)
	Err_Missing_SVG_Hooks         = FTError(0x9E)

	/* CFF, CID, and Type 1 errors */

	Err_Syntax_Error          = FTError(0xA0)
	Err_Stack_Underflow       = FTError(0xA1)
	Err_Ignore                = FTError(0xA2)
	Err_No_Unicode_Glyph_Name = FTError(0xA3)
	Err_Glyph_Too_Big         = FTError(0xA4)

	/* BDF errors */

	Err_Missing_Startfont_Field       = FTError(0xB0)
	Err_Missing_Font_Field            = FTError(0xB1)
	Err_Missing_Size_Field            = FTError(0xB2)
	Err_Missing_Fontboundingbox_Field = FTError(0xB3)
	Err_Missing_Chars_Field           = FTError(0xB4)
	Err_Missing_Startchar_Field       = FTError(0xB5)
	Err_Missing_Encoding_Field        = FTError(0xB6)
	Err_Missing_Bbx_Field             = FTError(0xB7)
	Err_Bbx_Too_Big                   = FTError(0xB8)
	Err_Corrupted_Font_Header         = FTError(0xB9)
	Err_Corrupted_Font_Glyphs         = FTError(0xBA)
)

var errorsText = map[FTError]string{
	Err_Ok:                            "no error",
	Err_Cannot_Open_Resource:          "cannot open resource",
	Err_Unknown_File_Format:           "unknown file format",
	Err_Invalid_File_Format:           "broken file",
	Err_Invalid_Version:               "invalid FreeType version",
	Err_Lower_Module_Version:          "module version is too low",
	Err_Invalid_Argument:              "invalid argument",
	Err_Unimplemented_Feature:         "unimplemented feature",
	Err_Invalid_Table:                 "broken table",
	Err_Invalid_Offset:                "broken offset within table",
	Err_Array_Too_Large:               "array allocation size too large",
	Err_Missing_Module:                "missing module",
	Err_Missing_Property:              "missing property",
	Err_Invalid_Glyph_Index:           "invalid glyph index",
	Err_Invalid_Character_Code:        "invalid character code",
	Err_Invalid_Glyph_Format:          "unsupported glyph image format",
	Err_Cannot_Render_Glyph:           "cannot render this glyph format",
	Err_Invalid_Outline:               "invalid outline",
	Err_Invalid_Composite:             "invalid composite glyph",
	Err_Too_Many_Hints:                "too many hints",
	Err_Invalid_Pixel_Size:            "invalid pixel size",
	Err_Invalid_SVG_Document:          "invalid SVG document",
	Err_Invalid_Handle:                "invalid object handle",
	Err_Invalid_Library_Handle:        "invalid library handle",
	Err_Invalid_Driver_Handle:         "invalid module handle",
	Err_Invalid_Face_Handle:           "invalid face handle",
	Err_Invalid_Size_Handle:           "invalid size handle",
	Err_Invalid_Slot_Handle:           "invalid glyph slot handle",
	Err_Invalid_CharMap_Handle:        "invalid charmap handle",
	Err_Invalid_Cache_Handle:          "invalid cache manager handle",
	Err_Invalid_Stream_Handle:         "invalid stream handle",
	Err_Too_Many_Drivers:              "too many modules",
	Err_Too_Many_Extensions:           "too many extensions",
	Err_Out_Of_Memory:                 "out of memory",
	Err_Unlisted_Object:               "unlisted object",
	Err_Cannot_Open_Stream:            "cannot open stream",
	Err_Invalid_Stream_Seek:           "invalid stream seek",
	Err_Invalid_Stream_Skip:           "invalid stream skip",
	Err_Invalid_Stream_Read:           "invalid stream read",
	Err_Invalid_Stream_Operation:      "invalid stream operation",
	Err_Invalid_Frame_Operation:       "invalid frame operation",
	Err_Nested_Frame_Access:           "nested frame access",
	Err_Invalid_Frame_Read:            "invalid frame read",
	Err_Raster_Uninitialized:          "raster uninitialized",
	Err_Raster_Corrupted:              "raster corrupted",
	Err_Raster_Overflow:               "raster overflow",
	Err_Raster_Negative_Height:        "negative height while rastering",
	Err_Too_Many_Caches:               "too many registered caches",
	Err_Invalid_Opcode:                "invalid opcode",
	Err_Too_Few_Arguments:             "too few arguments",
	Err_Stack_Overflow:                "stack overflow",
	Err_Code_Overflow:                 "code overflow",
	Err_Bad_Argument:                  "bad argument",
	Err_Divide_By_Zero:                "division by zero",
	Err_Invalid_Reference:             "invalid reference",
	Err_Debug_OpCode:                  "found debug opcode",
	Err_ENDF_In_Exec_Stream:           "found ENDF opcode in execution stream",
	Err_Nested_DEFS:                   "nested DEFS",
	Err_Invalid_CodeRange:             "invalid code range",
	Err_Execution_Too_Long:            "execution context too long",
	Err_Too_Many_Function_Defs:        "too many function definitions",
	Err_Too_Many_Instruction_Defs:     "too many instruction definitions",
	Err_Table_Missing:                 "SFNT font table missing",
	Err_Horiz_Header_Missing:          "horizontal header (hhea) table missing",
	Err_Locations_Missing:             "locations (loca) table missing",
	Err_Name_Table_Missing:            "name table missing",
	Err_CMap_Table_Missing:            "character map (cmap) table missing",
	Err_Hmtx_Table_Missing:            "horizontal metrics (hmtx) table missing",
	Err_Post_Table_Missing:            "PostScript (post) table missing",
	Err_Invalid_Horiz_Metrics:         "invalid horizontal metrics",
	Err_Invalid_CharMap_Format:        "invalid character map (cmap) format",
	Err_Invalid_PPem:                  "invalid ppem value",
	Err_Invalid_Vert_Metrics:          "invalid vertical metrics",
	Err_Could_Not_Find_Context:        "could not find context",
	Err_Invalid_Post_Table_Format:     "invalid PostScript (post) table format",
	Err_Invalid_Post_Table:            "invalid PostScript (post) table",
	Err_DEF_In_Glyf_Bytecode:          "found FDEF or IDEF opcode in glyf bytecode",
	Err_Missing_Bitmap:                "missing bitmap in strike",
	Err_Missing_SVG_Hooks:             "SVG hooks have not been set",
	Err_Syntax_Error:                  "opcode syntax error",
	Err_Stack_Underflow:               "argument stack underflow",
	Err_Ignore:                        "ignore",
	Err_No_Unicode_Glyph_Name:         "no Unicode glyph name found",
	Err_Glyph_Too_Big:                 "glyph too big for hinting",
	Err_Missing_Startfont_Field:       "`STARTFONT' field missing",
	Err_Missing_Font_Field:            "`FONT' field missing",
	Err_Missing_Size_Field:            "`SIZE' field missing",
	Err_Missing_Fontboundingbox_Field: "`FONTBOUNDINGBOX' field missing",
	Err_Missing_Chars_Field:           "`CHARS' field missing",
	Err_Missing_Startchar_Field:       "`STARTCHAR' field missing",
	Err_Missing_Encoding_Field:        "`ENCODING' field missing",
	Err_Missing_Bbx_Field:             "`BBX' field missing",
	Err_Bbx_Too_Big:                   "`BBX' too big",
	Err_Corrupted_Font_Header:         "Font header corrupted or missing fields",
	Err_Corrupted_Font_Glyphs:         "Font glyphs corrupted or missing fields",
}

/*
Error is used to represent errors returned from FreeType functions.

Its FTError method can be used to get the FTError value that was returned from the failing function.
*/
type Error struct {
	ftError FTError
	message string
}

func newError(err FTError, format string, args ...any) error {
	if err == Err_Ok {
		return nil
	}
	return Error{
		ftError: err,
		message: fmt.Sprintf(format, args...),
	}
}

func (err Error) Error() string {
	return fmt.Sprintf("%s : error %d : %s", err.message, err.ftError, errorsText[err.ftError])
}

// FTError returns the FTError value that was returned from the failing function.
func (err Error) FTError() FTError {
	return err.ftError
}
