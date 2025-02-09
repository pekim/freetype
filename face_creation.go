package freetype

import (
	"modernc.org/libc"
	"modernc.org/libfreetype"
)

// Functions to manage fonts.

/*
Face is a handle to a typographic face object.
A face object models a given typeface, in a given style.

https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_face
*/
type Face struct {
	face libfreetype.TFT_Face
	tls  *libc.TLS
}

// NewMemoryFace opens a font that has been loaded into memory.
//
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_new_memory_face
func (lib Library) NewMemoryFace(data []byte, faceIndex int) (Face, error) {
	face, freeFace := alloc(lib.tls, Face{})
	face.tls = lib.tls

	err := libfreetype.XFT_New_Memory_Face(
		lib.tls, lib.library,
		toUintptr(&data[0]), libfreetype.TFT_Long(len(data)),
		libfreetype.TFT_Long(faceIndex), toUintptr(&face.face))

	face_ := *face
	freeFace()
	return face_, newError(err, "failed to create a new memory face")
}
