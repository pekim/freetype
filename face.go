package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

func (face Face) Done() error {
	err := C.FT_Done_Face(face.face)
	return newError(err, "failed to discard face")
}

func (face Face) Reference() error {
	err := C.FT_Reference_Face(face.face)
	return newError(err, "failed to reference face")
}

func (face Face) Properties(properties ...Parameter) error {
	err := C.FT_Face_Properties(face.face, C.FT_UInt(len(properties)), (*C.FT_Parameter)(&properties[0]))
	for _, param := range properties {
		param.freeData()
	}
	return newError(err, "failed to set face properties")
}
