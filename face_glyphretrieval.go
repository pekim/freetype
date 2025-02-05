package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"
import (
	"unsafe"
)

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

type RenderMode = C.FT_Render_Mode

const (
	RENDER_MODE_NORMAL = RenderMode(C.FT_RENDER_MODE_NORMAL)
	RENDER_MODE_LIGHT  = RenderMode(C.FT_RENDER_MODE_LIGHT)
	RENDER_MODE_MONO   = RenderMode(C.FT_RENDER_MODE_MONO)
	RENDER_MODE_LCD    = RenderMode(C.FT_RENDER_MODE_LCD)
	RENDER_MODE_LCD_V  = RenderMode(C.FT_RENDER_MODE_LCD_V)
	RENDER_MODE_SDF    = RenderMode(C.FT_RENDER_MODE_SDF)
)

func (face Face) LoadGlyph(glyphIndex UInt, loadFlags Int32) error {
	err := C.FT_Load_Glyph(face.face, glyphIndex, loadFlags)
	return newError(err, "failed to load glyph index %d with flags %04x", glyphIndex, loadFlags)
}

func (face Face) RenderGlyph(renderMode RenderMode) error {
	glyph := face.Rec().Glyph
	err := C.FT_Render_Glyph((*C.FT_GlyphSlotRec)(unsafe.Pointer(glyph)), renderMode)
	return newError(err, "failed to render glyph with index %d for render mode %04x", glyph.GlyphIndex, renderMode)
}
