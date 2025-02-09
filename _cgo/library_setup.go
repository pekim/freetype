package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
//
// #include <stdlib.h>
import "C"

// Functions to start and end the usage of the FreeType library.

/*
Library is a handle to a FreeType library instance.

https://freetype.org/freetype2/docs/reference/ft2-library_setup.html#ft_library
*/
type Library struct {
	library C.FT_Library
}

/*
Init initializes a new FreeType library object.

https://freetype.org/freetype2/docs/reference/ft2-library_setup.html#ft_init_freetype
*/
func Init() (Library, error) {
	lib := Library{}
	err := C.FT_Init_FreeType(&lib.library)
	return lib, newError(err, "failed to init library")
}

// Done destroys the FreeType library object represented by Library,
// and all of its children, including resources, drivers, faces, sizes, etc.
//
// https://freetype.org/freetype2/docs/reference/ft2-library_setup.html#ft_done_freetype
func (lib Library) Done() error {
	err := C.FT_Done_FreeType(lib.library)
	return newError(err, "failed to destroy library")
}

// Version returns the version of the FreeType library being used.
//
// The 3 values returned are
//   - The major version number.
//   - The minor version number.
//   - The patch version number.
//
// https://freetype.org/freetype2/docs/reference/ft2-library_setup.html#ft_library_version
func (lib Library) Version() (int, int, int) {
	var major, minor, patch C.FT_Int
	C.FT_Library_Version(lib.library, &major, &minor, &patch)
	return int(major), int(minor), int(patch)
}

// FREETYPE_XXX
// https://freetype.org/freetype2/docs/reference/ft2-library_setup.html#freetype_xxx
