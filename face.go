package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
/*

FT_Bool c_FT_HAS_HORIZONTAL(FT_Face face) {
	return FT_HAS_HORIZONTAL(face);
}

FT_Bool c_FT_HAS_VERTICAL(FT_Face face) {
	return FT_HAS_VERTICAL(face);
}

FT_Bool c_FT_HAS_KERNING(FT_Face face) {
	return FT_HAS_KERNING(face);
}

FT_Bool c_FT_HAS_FIXED_SIZES(FT_Face face) {
	return FT_HAS_FIXED_SIZES(face);
}

FT_Bool c_FT_HAS_GLYPH_NAMES(FT_Face face) {
	return FT_HAS_GLYPH_NAMES(face);
}

FT_Bool c_FT_HAS_COLOR(FT_Face face) {
	return FT_HAS_COLOR(face);
}

FT_Bool c_FT_HAS_MULTIPLE_MASTERS(FT_Face face) {
	return FT_HAS_MULTIPLE_MASTERS(face);
}

FT_Bool c_FT_HAS_SVG(FT_Face face) {
	return FT_HAS_SVG(face);
}

FT_Bool c_FT_HAS_SBIX(FT_Face face) {
	return FT_HAS_SBIX(face);
}

FT_Bool c_FT_HAS_SBIX_OVERLAY(FT_Face face) {
	return FT_HAS_SBIX_OVERLAY(face);
}

FT_Bool c_FT_IS_SFNT(FT_Face face) {
	return FT_IS_SFNT(face);
}

FT_Bool c_FT_IS_SCALABLE(FT_Face face) {
	return FT_IS_SCALABLE(face);
}

FT_Bool c_FT_IS_FIXED_WIDTH(FT_Face face) {
	return FT_IS_FIXED_WIDTH(face);
}

FT_Bool c_FT_IS_CID_KEYED(FT_Face face) {
	return FT_IS_CID_KEYED(face);
}

FT_Bool c_FT_IS_TRICKY(FT_Face face) {
	return FT_IS_TRICKY(face);
}

FT_Bool c_FT_IS_NAMED_INSTANCE(FT_Face face) {
	return FT_IS_NAMED_INSTANCE(face);
}

FT_Bool c_FT_IS_VARIATION(FT_Face face) {
	return FT_IS_VARIATION(face);
}

*/
import "C"

type Face struct {
	face C.FT_Face
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

func (face Face) HasHorizontal() bool {
	return cBoolToGo(C.c_FT_HAS_HORIZONTAL(face.face))
}

func (face Face) HasVertical() bool {
	return cBoolToGo(C.c_FT_HAS_VERTICAL(face.face))
}

func (face Face) HasKerning() bool {
	return cBoolToGo(C.c_FT_HAS_KERNING(face.face))
}

func (face Face) HadFixedSizes() bool {
	return cBoolToGo(C.c_FT_HAS_FIXED_SIZES(face.face))
}

func (face Face) HasGlyphNames() bool {
	return cBoolToGo(C.c_FT_HAS_GLYPH_NAMES(face.face))
}

func (face Face) HasColor() bool {
	return cBoolToGo(C.c_FT_HAS_COLOR(face.face))
}

func (face Face) HasMultipleMasters() bool {
	return cBoolToGo(C.c_FT_HAS_MULTIPLE_MASTERS(face.face))
}

func (face Face) HaseSVG() bool {
	return cBoolToGo(C.c_FT_HAS_SVG(face.face))
}

func (face Face) HasSbix() bool {
	return cBoolToGo(C.c_FT_HAS_SBIX(face.face))
}

func (face Face) HasSbixOverlay() bool {
	return cBoolToGo(C.c_FT_HAS_SBIX_OVERLAY(face.face))
}

func (face Face) IsSFNT() bool {
	return cBoolToGo(C.c_FT_IS_SFNT(face.face))
}

func (face Face) IsScalable() bool {
	return cBoolToGo(C.c_FT_IS_SCALABLE(face.face))
}

func (face Face) IsFixedWidth() bool {
	return cBoolToGo(C.c_FT_IS_FIXED_WIDTH(face.face))
}

func (face Face) IsCIDKeyed() bool {
	return cBoolToGo(C.c_FT_IS_CID_KEYED(face.face))
}

func (face Face) IsTricky() bool {
	return cBoolToGo(C.c_FT_IS_TRICKY(face.face))
}

func (face Face) IsNamedInstance() bool {
	return cBoolToGo(C.c_FT_IS_NAMED_INSTANCE(face.face))
}
func (face Face) IsVariation() bool {
	return cBoolToGo(C.c_FT_IS_VARIATION(face.face))
}
