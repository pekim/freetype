package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

type LoadFlag = Int32

const (
	LOAD_DEFAULT                     = LoadFlag(C.FT_LOAD_DEFAULT)
	LOAD_NO_SCALE                    = LoadFlag(C.FT_LOAD_NO_SCALE)
	LOAD_NO_HINTING                  = LoadFlag(C.FT_LOAD_NO_HINTING)
	LOAD_RENDER                      = LoadFlag(C.FT_LOAD_RENDER)
	LOAD_NO_BITMAP                   = LoadFlag(C.FT_LOAD_NO_BITMAP)
	LOAD_VERTICAL_LAYOUT             = LoadFlag(C.FT_LOAD_VERTICAL_LAYOUT)
	LOAD_FORCE_AUTOHINT              = LoadFlag(C.FT_LOAD_FORCE_AUTOHINT)
	LOAD_CROP_BITMAP                 = LoadFlag(C.FT_LOAD_CROP_BITMAP)
	LOAD_PEDANTIC                    = LoadFlag(C.FT_LOAD_PEDANTIC)
	LOAD_IGNORE_GLOBAL_ADVANCE_WIDTH = LoadFlag(C.FT_LOAD_IGNORE_GLOBAL_ADVANCE_WIDTH)
	LOAD_NO_RECURSE                  = LoadFlag(C.FT_LOAD_NO_RECURSE)
	LOAD_IGNORE_TRANSFORM            = LoadFlag(C.FT_LOAD_IGNORE_TRANSFORM)
	LOAD_MONOCHROME                  = LoadFlag(C.FT_LOAD_MONOCHROME)
	LOAD_LINEAR_DESIGN               = LoadFlag(C.FT_LOAD_LINEAR_DESIGN)
	LOAD_SBITS_ONLY                  = LoadFlag(C.FT_LOAD_SBITS_ONLY)
	LOAD_NO_AUTOHINT                 = LoadFlag(C.FT_LOAD_NO_AUTOHINT)

	/* Bits 16-19 are used by `FT_LOAD_TARGET_` */
	LOAD_COLOR               = LoadFlag(C.FT_LOAD_COLOR)
	LOAD_COMPUTE_METRICS     = LoadFlag(C.FT_LOAD_COMPUTE_METRICS)
	LOAD_BITMAP_METRICS_ONLY = LoadFlag(C.FT_LOAD_BITMAP_METRICS_ONLY)
	LOAD_NO_SVG              = LoadFlag(C.FT_LOAD_NO_SVG)
)

// func (face Face) LoadGlyph(glyphIndex UInt, loadFlags Int32) error {
// 	err := C.FT_Load_Glyph(face.face, glyphIndex, loadFlags)
// 	return newError(err, "failed to load glyph index %d with flags %04x", glyphIndex, loadFlags)
// }
