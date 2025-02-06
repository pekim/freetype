package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
//
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

func (lib Library) NewFace(filepathname string, index int) (Face, error) {
	cFilepathname := C.CString(filepathname)
	defer C.free(unsafe.Pointer(cFilepathname))

	face := Face{}
	err := C.FT_New_Face(lib.library, cFilepathname, C.FT_Long(index), &face.face)
	return face, newError(err, "failed to create a face for file '%s'", filepathname)
}

func (lib Library) NewMemoryFace(data []byte, index int) (Face, error) {
	face := Face{}
	err := C.FT_New_Memory_Face(lib.library, (*C.FT_Byte)(unsafe.Pointer(&data[0])), C.FT_Long(len(data)), C.FT_Long(index), &face.face)
	return face, newError(err, "failed to create a new memory face")
}
