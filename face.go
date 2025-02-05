package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"unsafe"
)

type FACE_FLAG = Long

const (
	FACE_FLAG_SCALABLE         = FACE_FLAG(C.FT_FACE_FLAG_SCALABLE)
	FACE_FLAG_FIXED_SIZES      = FACE_FLAG(C.FT_FACE_FLAG_FIXED_SIZES)
	FACE_FLAG_FIXED_WIDTH      = FACE_FLAG(C.FT_FACE_FLAG_FIXED_WIDTH)
	FACE_FLAG_SFNT             = FACE_FLAG(C.FT_FACE_FLAG_SFNT)
	FACE_FLAG_HORIZONTAL       = FACE_FLAG(C.FT_FACE_FLAG_HORIZONTAL)
	FACE_FLAG_VERTICAL         = FACE_FLAG(C.FT_FACE_FLAG_VERTICAL)
	FACE_FLAG_KERNING          = FACE_FLAG(C.FT_FACE_FLAG_KERNING)
	FACE_FLAG_FAST_GLYPHS      = FACE_FLAG(C.FT_FACE_FLAG_FAST_GLYPHS)
	FACE_FLAG_MULTIPLE_MASTERS = FACE_FLAG(C.FT_FACE_FLAG_MULTIPLE_MASTERS)
	FACE_FLAG_GLYPH_NAMES      = FACE_FLAG(C.FT_FACE_FLAG_GLYPH_NAMES)
	FACE_FLAG_EXTERNAL_STREAM  = FACE_FLAG(C.FT_FACE_FLAG_EXTERNAL_STREAM)
	FACE_FLAG_HINTER           = FACE_FLAG(C.FT_FACE_FLAG_HINTER)
	FACE_FLAG_CID_KEYED        = FACE_FLAG(C.FT_FACE_FLAG_CID_KEYED)
	FACE_FLAG_TRICKY           = FACE_FLAG(C.FT_FACE_FLAG_TRICKY)
	FACE_FLAG_COLOR            = FACE_FLAG(C.FT_FACE_FLAG_COLOR)
	FACE_FLAG_VARIATION        = FACE_FLAG(C.FT_FACE_FLAG_VARIATION)
	FACE_FLAG_SVG              = FACE_FLAG(C.FT_FACE_FLAG_SVG)
	FACE_FLAG_SBIX             = FACE_FLAG(C.FT_FACE_FLAG_SBIX)
	FACE_FLAG_SBIX_OVERLAY     = FACE_FLAG(C.FT_FACE_FLAG_SBIX_OVERLAY)
)

type STYLE_FLAG = Long

const (
	STYLE_FLAG_ITALIC = STYLE_FLAG(C.FT_STYLE_FLAG_ITALIC)
	STYLE_FLAG_BOLD   = STYLE_FLAG(C.FT_STYLE_FLAG_BOLD)
)

type Face struct {
	face C.FT_Face
}

func (face Face) Rec() *FaceRec {
	return (*FaceRec)(unsafe.Pointer((face.face)))
}

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

func (fr *FaceRec) FamilyName() string {
	return C.GoString(fr.family_name)
}

func (fr *FaceRec) StyleName() string {
	return C.GoString(fr.style_name)
}

func (fr *FaceRec) AvailableSizes() []BitmapSize {
	return unsafe.Slice(fr.available_sizes, fr.num_fixed_sizes)
}

func (fr *FaceRec) Charmaps() []*CharMapRec {
	return unsafe.Slice(fr.charmaps, fr.num_charmaps)
}
