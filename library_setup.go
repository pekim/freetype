package freetype

import (
	"modernc.org/libc"
	"modernc.org/libfreetype"
)

// Functions to start and end the usage of the FreeType library.

/*
Library is a handle to a FreeType library instance.

https://freetype.org/freetype2/docs/reference/ft2-library_setup.html#ft_library
*/
type Library struct {
	library libfreetype.TFT_Library
	tls     *libc.TLS
}

/*
Init initializes a new FreeType library object.

https://freetype.org/freetype2/docs/reference/ft2-library_setup.html#ft_init_freetype
*/
func Init() (Library, error) {
	tls := libc.NewTLS()
	lib, freeLib := alloc(tls, Library{})
	defer freeLib()
	lib.tls = tls

	err := libfreetype.XFT_Init_FreeType(tls, toUintptr(&lib.library))
	if err != Err_Ok {
		return Library{}, newError(err, "failed to init library")
	}
	return *lib, nil
}

// Done destroys the FreeType library object represented by Library,
// and all of its children, including resources, drivers, faces, sizes, etc.
//
// https://freetype.org/freetype2/docs/reference/ft2-library_setup.html#ft_done_freetype
func (lib Library) Done() error {
	err := libfreetype.XFT_Done_FreeType(lib.tls, lib.library)
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
	var major, minor, patch libfreetype.TFT_Int
	libfreetype.XFT_Library_Version(lib.tls, lib.library,
		toUintptr(&major), toUintptr(&minor), toUintptr(&patch))
	return int(major), int(minor), int(patch)
}

// FREETYPE_XXX
// https://freetype.org/freetype2/docs/reference/ft2-library_setup.html#freetype_xxx
