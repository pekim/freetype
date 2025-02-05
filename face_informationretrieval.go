package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"unsafe"
)

func (face Face) GetGlyphName(glyphIndex UInt) (string, error) {
	buffer := make([]C.char, 128)
	err := C.FT_Get_Glyph_Name(face.face, glyphIndex, C.FT_Pointer(unsafe.Pointer(&buffer[0])), UInt(len(buffer)))
	name := C.GoString(&buffer[0])
	return name, newError(err, "failed to get glyph name for glyph index %d", glyphIndex)

}
