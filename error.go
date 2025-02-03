package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"fmt"
)

const Err_Ok = C.FT_Err_Ok
const Err_Cannot_Open_Resource = C.FT_Err_Cannot_Open_Resource
const Err_Unknown_File_Format = C.FT_Err_Unknown_File_Format
const Err_Invalid_File_Format = C.FT_Err_Invalid_File_Format
const Err_Invalid_Version = C.FT_Err_Invalid_Version
const Err_Lower_Module_Version = C.FT_Err_Lower_Module_Version
const Err_Invalid_Argument = C.FT_Err_Invalid_Argument
const Err_Unimplemented_Feature = C.FT_Err_Unimplemented_Feature
const Err_Invalid_Table = C.FT_Err_Invalid_Table
const Err_Invalid_Offset = C.FT_Err_Invalid_Offset
const Err_Array_Too_Large = C.FT_Err_Array_Too_Large
const Err_Missing_Module = C.FT_Err_Missing_Module
const Err_Missing_Property = C.FT_Err_Missing_Property

/* glyph/character errors */

const Err_Invalid_Glyph_Index = C.FT_Err_Invalid_Glyph_Index
const Err_Invalid_Character_Code = C.FT_Err_Invalid_Character_Code
const Err_Invalid_Glyph_Format = C.FT_Err_Invalid_Glyph_Format
const Err_Cannot_Render_Glyph = C.FT_Err_Cannot_Render_Glyph
const Err_Invalid_Outline = C.FT_Err_Invalid_Outline
const Err_Invalid_Composite = C.FT_Err_Invalid_Composite
const Err_Too_Many_Hints = C.FT_Err_Too_Many_Hints
const Err_Invalid_Pixel_Size = C.FT_Err_Invalid_Pixel_Size
const Err_Invalid_SVG_Document = C.FT_Err_Invalid_SVG_Document

/* handle errors */

const Err_Invalid_Handle = C.FT_Err_Invalid_Handle
const Err_Invalid_Library_Handle = C.FT_Err_Invalid_Library_Handle
const Err_Invalid_Driver_Handle = C.FT_Err_Invalid_Driver_Handle
const Err_Invalid_Face_Handle = C.FT_Err_Invalid_Face_Handle
const Err_Invalid_Size_Handle = C.FT_Err_Invalid_Size_Handle
const Err_Invalid_Slot_Handle = C.FT_Err_Invalid_Slot_Handle
const Err_Invalid_CharMap_Handle = C.FT_Err_Invalid_CharMap_Handle
const Err_Invalid_Cache_Handle = C.FT_Err_Invalid_Cache_Handle
const Err_Invalid_Stream_Handle = C.FT_Err_Invalid_Stream_Handle

/* driver errors */

const Err_Too_Many_Drivers = C.FT_Err_Too_Many_Drivers
const Err_Too_Many_Extensions = C.FT_Err_Too_Many_Extensions

/* memory errors */

const Err_Out_Of_Memory = C.FT_Err_Out_Of_Memory
const Err_Unlisted_Object = C.FT_Err_Unlisted_Object

/* stream errors */

const Err_Cannot_Open_Stream = C.FT_Err_Cannot_Open_Stream
const Err_Invalid_Stream_Seek = C.FT_Err_Invalid_Stream_Seek
const Err_Invalid_Stream_Skip = C.FT_Err_Invalid_Stream_Skip
const Err_Invalid_Stream_Read = C.FT_Err_Invalid_Stream_Read
const Err_Invalid_Stream_Operation = C.FT_Err_Invalid_Stream_Operation
const Err_Invalid_Frame_Operation = C.FT_Err_Invalid_Frame_Operation
const Err_Nested_Frame_Access = C.FT_Err_Nested_Frame_Access
const Err_Invalid_Frame_Read = C.FT_Err_Invalid_Frame_Read

/* raster errors */

const Err_Raster_Uninitialized = C.FT_Err_Raster_Uninitialized
const Err_Raster_Corrupted = C.FT_Err_Raster_Corrupted
const Err_Raster_Overflow = C.FT_Err_Raster_Overflow
const Err_Raster_Negative_Height = C.FT_Err_Raster_Negative_Height

/* cache errors */

const Err_Too_Many_Caches = C.FT_Err_Too_Many_Caches

/* TrueType and SFNT errors */

const Err_Invalid_Opcode = C.FT_Err_Invalid_Opcode
const Err_Too_Few_Arguments = C.FT_Err_Too_Few_Arguments
const Err_Stack_Overflow = C.FT_Err_Stack_Overflow
const Err_Code_Overflow = C.FT_Err_Code_Overflow
const Err_Bad_Argument = C.FT_Err_Bad_Argument
const Err_Divide_By_Zero = C.FT_Err_Divide_By_Zero
const Err_Invalid_Reference = C.FT_Err_Invalid_Reference
const Err_Debug_OpCode = C.FT_Err_Debug_OpCode
const Err_ENDF_In_Exec_Stream = C.FT_Err_ENDF_In_Exec_Stream
const Err_Nested_DEFS = C.FT_Err_Nested_DEFS
const Err_Invalid_CodeRange = C.FT_Err_Invalid_CodeRange
const Err_Execution_Too_Long = C.FT_Err_Execution_Too_Long
const Err_Too_Many_Function_Defs = C.FT_Err_Too_Many_Function_Defs
const Err_Too_Many_Instruction_Defs = C.FT_Err_Too_Many_Instruction_Defs
const Err_Table_Missing = C.FT_Err_Table_Missing
const Err_Horiz_Header_Missing = C.FT_Err_Horiz_Header_Missing
const Err_Locations_Missing = C.FT_Err_Locations_Missing
const Err_Name_Table_Missing = C.FT_Err_Name_Table_Missing
const Err_CMap_Table_Missing = C.FT_Err_CMap_Table_Missing
const Err_Hmtx_Table_Missing = C.FT_Err_Hmtx_Table_Missing
const Err_Post_Table_Missing = C.FT_Err_Post_Table_Missing
const Err_Invalid_Horiz_Metrics = C.FT_Err_Invalid_Horiz_Metrics
const Err_Invalid_CharMap_Format = C.FT_Err_Invalid_CharMap_Format
const Err_Invalid_PPem = C.FT_Err_Invalid_PPem
const Err_Invalid_Vert_Metrics = C.FT_Err_Invalid_Vert_Metrics
const Err_Could_Not_Find_Context = C.FT_Err_Could_Not_Find_Context
const Err_Invalid_Post_Table_Format = C.FT_Err_Invalid_Post_Table_Format
const Err_Invalid_Post_Table = C.FT_Err_Invalid_Post_Table
const Err_DEF_In_Glyf_Bytecode = C.FT_Err_DEF_In_Glyf_Bytecode
const Err_Missing_Bitmap = C.FT_Err_Missing_Bitmap
const Err_Missing_SVG_Hooks = C.FT_Err_Missing_SVG_Hooks

/* CFF, CID, and Type 1 errors */

const Err_Syntax_Error = C.FT_Err_Syntax_Error
const Err_Stack_Underflow = C.FT_Err_Stack_Underflow
const Err_Ignore = C.FT_Err_Ignore
const Err_No_Unicode_Glyph_Name = C.FT_Err_No_Unicode_Glyph_Name
const Err_Glyph_Too_Big = C.FT_Err_Glyph_Too_Big

/* BDF errors */

const Err_Missing_Startfont_Field = C.FT_Err_Missing_Startfont_Field
const Err_Missing_Font_Field = C.FT_Err_Missing_Font_Field
const Err_Missing_Size_Field = C.FT_Err_Missing_Size_Field
const Err_Missing_Fontboundingbox_Field = C.FT_Err_Missing_Fontboundingbox_Field
const Err_Missing_Chars_Field = C.FT_Err_Missing_Chars_Field
const Err_Missing_Startchar_Field = C.FT_Err_Missing_Startchar_Field
const Err_Missing_Encoding_Field = C.FT_Err_Missing_Encoding_Field
const Err_Missing_Bbx_Field = C.FT_Err_Missing_Bbx_Field
const Err_Bbx_Too_Big = C.FT_Err_Bbx_Too_Big
const Err_Corrupted_Font_Header = C.FT_Err_Corrupted_Font_Header
const Err_Corrupted_Font_Glyphs = C.FT_Err_Corrupted_Font_Glyphs

var errorsText = map[C.FT_Error]string{
	C.FT_Err_Ok:                            "no error",
	C.FT_Err_Cannot_Open_Resource:          "cannot open resource",
	C.FT_Err_Unknown_File_Format:           "unknown file format",
	C.FT_Err_Invalid_File_Format:           "broken file",
	C.FT_Err_Invalid_Version:               "invalid FreeType version",
	C.FT_Err_Lower_Module_Version:          "module version is too low",
	C.FT_Err_Invalid_Argument:              "invalid argument",
	C.FT_Err_Unimplemented_Feature:         "unimplemented feature",
	C.FT_Err_Invalid_Table:                 "broken table",
	C.FT_Err_Invalid_Offset:                "broken offset within table",
	C.FT_Err_Array_Too_Large:               "array allocation size too large",
	C.FT_Err_Missing_Module:                "missing module",
	C.FT_Err_Missing_Property:              "missing property",
	C.FT_Err_Invalid_Glyph_Index:           "invalid glyph index",
	C.FT_Err_Invalid_Character_Code:        "invalid character code",
	C.FT_Err_Invalid_Glyph_Format:          "unsupported glyph image format",
	C.FT_Err_Cannot_Render_Glyph:           "cannot render this glyph format",
	C.FT_Err_Invalid_Outline:               "invalid outline",
	C.FT_Err_Invalid_Composite:             "invalid composite glyph",
	C.FT_Err_Too_Many_Hints:                "too many hints",
	C.FT_Err_Invalid_Pixel_Size:            "invalid pixel size",
	C.FT_Err_Invalid_SVG_Document:          "invalid SVG document",
	C.FT_Err_Invalid_Handle:                "invalid object handle",
	C.FT_Err_Invalid_Library_Handle:        "invalid library handle",
	C.FT_Err_Invalid_Driver_Handle:         "invalid module handle",
	C.FT_Err_Invalid_Face_Handle:           "invalid face handle",
	C.FT_Err_Invalid_Size_Handle:           "invalid size handle",
	C.FT_Err_Invalid_Slot_Handle:           "invalid glyph slot handle",
	C.FT_Err_Invalid_CharMap_Handle:        "invalid charmap handle",
	C.FT_Err_Invalid_Cache_Handle:          "invalid cache manager handle",
	C.FT_Err_Invalid_Stream_Handle:         "invalid stream handle",
	C.FT_Err_Too_Many_Drivers:              "too many modules",
	C.FT_Err_Too_Many_Extensions:           "too many extensions",
	C.FT_Err_Out_Of_Memory:                 "out of memory",
	C.FT_Err_Unlisted_Object:               "unlisted object",
	C.FT_Err_Cannot_Open_Stream:            "cannot open stream",
	C.FT_Err_Invalid_Stream_Seek:           "invalid stream seek",
	C.FT_Err_Invalid_Stream_Skip:           "invalid stream skip",
	C.FT_Err_Invalid_Stream_Read:           "invalid stream read",
	C.FT_Err_Invalid_Stream_Operation:      "invalid stream operation",
	C.FT_Err_Invalid_Frame_Operation:       "invalid frame operation",
	C.FT_Err_Nested_Frame_Access:           "nested frame access",
	C.FT_Err_Invalid_Frame_Read:            "invalid frame read",
	C.FT_Err_Raster_Uninitialized:          "raster uninitialized",
	C.FT_Err_Raster_Corrupted:              "raster corrupted",
	C.FT_Err_Raster_Overflow:               "raster overflow",
	C.FT_Err_Raster_Negative_Height:        "negative height while rastering",
	C.FT_Err_Too_Many_Caches:               "too many registered caches",
	C.FT_Err_Invalid_Opcode:                "invalid opcode",
	C.FT_Err_Too_Few_Arguments:             "too few arguments",
	C.FT_Err_Stack_Overflow:                "stack overflow",
	C.FT_Err_Code_Overflow:                 "code overflow",
	C.FT_Err_Bad_Argument:                  "bad argument",
	C.FT_Err_Divide_By_Zero:                "division by zero",
	C.FT_Err_Invalid_Reference:             "invalid reference",
	C.FT_Err_Debug_OpCode:                  "found debug opcode",
	C.FT_Err_ENDF_In_Exec_Stream:           "found ENDF opcode in execution stream",
	C.FT_Err_Nested_DEFS:                   "nested DEFS",
	C.FT_Err_Invalid_CodeRange:             "invalid code range",
	C.FT_Err_Execution_Too_Long:            "execution context too long",
	C.FT_Err_Too_Many_Function_Defs:        "too many function definitions",
	C.FT_Err_Too_Many_Instruction_Defs:     "too many instruction definitions",
	C.FT_Err_Table_Missing:                 "SFNT font table missing",
	C.FT_Err_Horiz_Header_Missing:          "horizontal header (hhea) table missing",
	C.FT_Err_Locations_Missing:             "locations (loca) table missing",
	C.FT_Err_Name_Table_Missing:            "name table missing",
	C.FT_Err_CMap_Table_Missing:            "character map (cmap) table missing",
	C.FT_Err_Hmtx_Table_Missing:            "horizontal metrics (hmtx) table missing",
	C.FT_Err_Post_Table_Missing:            "PostScript (post) table missing",
	C.FT_Err_Invalid_Horiz_Metrics:         "invalid horizontal metrics",
	C.FT_Err_Invalid_CharMap_Format:        "invalid character map (cmap) format",
	C.FT_Err_Invalid_PPem:                  "invalid ppem value",
	C.FT_Err_Invalid_Vert_Metrics:          "invalid vertical metrics",
	C.FT_Err_Could_Not_Find_Context:        "could not find context",
	C.FT_Err_Invalid_Post_Table_Format:     "invalid PostScript (post) table format",
	C.FT_Err_Invalid_Post_Table:            "invalid PostScript (post) table",
	C.FT_Err_DEF_In_Glyf_Bytecode:          "found FDEF or IDEF opcode in glyf bytecode",
	C.FT_Err_Missing_Bitmap:                "missing bitmap in strike",
	C.FT_Err_Missing_SVG_Hooks:             "SVG hooks have not been set",
	C.FT_Err_Syntax_Error:                  "opcode syntax error",
	C.FT_Err_Stack_Underflow:               "argument stack underflow",
	C.FT_Err_Ignore:                        "ignore",
	C.FT_Err_No_Unicode_Glyph_Name:         "no Unicode glyph name found",
	C.FT_Err_Glyph_Too_Big:                 "glyph too big for hinting",
	C.FT_Err_Missing_Startfont_Field:       "`STARTFONT' field missing",
	C.FT_Err_Missing_Font_Field:            "`FONT' field missing",
	C.FT_Err_Missing_Size_Field:            "`SIZE' field missing",
	C.FT_Err_Missing_Fontboundingbox_Field: "`FONTBOUNDINGBOX' field missing",
	C.FT_Err_Missing_Chars_Field:           "`CHARS' field missing",
	C.FT_Err_Missing_Startchar_Field:       "`STARTCHAR' field missing",
	C.FT_Err_Missing_Encoding_Field:        "`ENCODING' field missing",
	C.FT_Err_Missing_Bbx_Field:             "`BBX' field missing",
	C.FT_Err_Bbx_Too_Big:                   "`BBX' too big",
	C.FT_Err_Corrupted_Font_Header:         "Font header corrupted or missing fields",
	C.FT_Err_Corrupted_Font_Glyphs:         "Font glyphs corrupted or missing fields",
}

type Error struct {
	ftError C.FT_Error
	action  string
}

func newError(err C.FT_Error, action string) error {
	if err == C.FT_Err_Ok {
		return nil
	}
	return Error{
		ftError: err,
		action:  action,
	}
}

func (err Error) Error() string {
	return fmt.Sprintf("%s : error %d : %s", err.action, err.ftError, errorsText[err.ftError])
}
