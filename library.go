package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
//
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

type Library struct {
	library C.FT_Library
}

func Init() (Library, error) {
	lib := Library{}
	err := C.FT_Init_FreeType(&lib.library)
	return lib, newError(err, "failed to init library")
}

func (lib Library) Done() error {
	err := C.FT_Done_FreeType(lib.library)
	return newError(err, "failed to destroy library")
}

func (lib Library) Version() (int, int, int) {
	var major, minor, patch C.FT_Int
	C.FT_Library_Version(lib.library, &major, &minor, &patch)
	return int(major), int(minor), int(patch)
}

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
