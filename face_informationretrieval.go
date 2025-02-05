package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

type FSType = UShort

const (
	FSTYPE_INSTALLABLE_EMBEDDING        = FSType(C.FT_FSTYPE_INSTALLABLE_EMBEDDING)
	FSTYPE_RESTRICTED_LICENSE_EMBEDDING = FSType(C.FT_FSTYPE_RESTRICTED_LICENSE_EMBEDDING)
	FSTYPE_PREVIEW_AND_PRINT_EMBEDDING  = FSType(C.FT_FSTYPE_PREVIEW_AND_PRINT_EMBEDDING)
	FSTYPE_EDITABLE_EMBEDDING           = FSType(C.FT_FSTYPE_EDITABLE_EMBEDDING)
	FSTYPE_NO_SUBSETTING                = FSType(C.FT_FSTYPE_NO_SUBSETTING)
	FSTYPE_BITMAP_EMBEDDING_ONLY        = FSType(C.FT_FSTYPE_BITMAP_EMBEDDING_ONLY)
)

func (face Face) GetNameIndex(glyphName string) UInt {
	cName := C.CString(glyphName)
	defer C.free(unsafe.Pointer(cName))

	return C.FT_Get_Name_Index(face.face, cName)
}

func (face Face) GetGlyphName(glyphIndex UInt) (string, error) {
	buffer := make([]C.char, 128)
	err := C.FT_Get_Glyph_Name(face.face, glyphIndex, C.FT_Pointer(unsafe.Pointer(&buffer[0])), UInt(len(buffer)))
	name := C.GoString(&buffer[0])
	return name, newError(err, "failed to get glyph name for glyph index %d", glyphIndex)
}

func (face Face) GetPostscriptName() string {
	cName := C.FT_Get_Postscript_Name(face.face)
	if cName == nil {
		return ""
	}
	return C.GoString(cName)
}

func (face Face) GetFSTypeFlags() FSType {
	return C.FT_Get_FSType_Flags(face.face)
}
