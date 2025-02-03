package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

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
