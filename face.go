package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

type Face struct {
	face C.FT_Face
}

func (face Face) Done() error {
	err := C.FT_Done_Face(face.face)
	return newError(err, "failed to discard a face")
}
