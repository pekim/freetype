package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"unsafe"
)

// Functions to manage glyphs.

/*
GlyphSlot is a handle to a given ‘glyph slot’.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_glyphslot
*/
type GlyphSlot *GlyphSlotRec

func init() {
	assertSameSize(GlyphSlotRec{}, C.FT_GlyphSlotRec{})
}

/*
GlyphSlotRec is the FreeType root glyph slot class structure.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_glyphslotrec
*/
type GlyphSlotRec struct {
	Library    Library
	Face       Face
	Next       GlyphSlot
	GlyphIndex UInt
	Generic    Generic

	Metrics           GlyphMetrics
	LinearHoriAdvance Fixed
	LinearVertAdvance Fixed
	Advance           Vector

	Format GlyphFormat

	Bitmap     Bitmap
	BitmapLeft Int
	BitmapTop  Int

	Outline C.FT_Outline

	NumSubglyphs UInt
	subglyphs    unsafe.Pointer

	control_data unsafe.Pointer
	control_len  C.long

	LsbDelta Pos
	RsbDelta Pos

	other    unsafe.Pointer
	internal unsafe.Pointer
}

func init() {
	assertSameSize(GlyphMetrics{}, C.FT_Glyph_Metrics{})
}

/*
GlyphMetrics is a structure to model the metrics of a single glyph.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_glyph_metrics
*/
type GlyphMetrics struct {
	Width  Pos
	Height Pos

	HoriBearingX Pos
	HoriBearingY Pos
	HoriAdvance  Pos

	VertBearingX Pos
	VertBearingY Pos
	VertAdvance  Pos
}

/*
LoadGlyph loads a glyph into the glyph slot of a face object.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_load_glyph
*/
func (face Face) LoadGlyph(glyphIndex UInt, loadFlags Int32) error {
	err := C.FT_Load_Glyph(face.face, glyphIndex, loadFlags)
	return newError(err, "failed to load glyph index %d with flags %04x", glyphIndex, loadFlags)
}

/*
LoadFlag is a list of bit field constants for LoadGlyph to indicate what kind of operations to perform during glyph loading.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_load_xxx
*/
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

/*
LoadTarget is a list of values to select a specific hinting algorithm for the hinter.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_load_target_xxx
*/
type LoadTarget = Int32

const (
	LOAD_TARGET_NORMAL = LoadTarget(C.FT_LOAD_TARGET_NORMAL)
	LOAD_TARGET_LIGHT  = LoadTarget(C.FT_LOAD_TARGET_LIGHT)
	LOAD_TARGET_MONO   = LoadTarget(C.FT_LOAD_TARGET_MONO)
	LOAD_TARGET_LCD    = LoadTarget(C.FT_LOAD_TARGET_LCD)
	LOAD_TARGET_LCD_V  = LoadTarget(C.FT_LOAD_TARGET_LCD_V)
)

/*
RenderGlyph convert a given glyph image to a bitmap.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_render_glyph
*/
func (face Face) RenderGlyph(renderMode RenderMode) error {
	glyph := face.Rec().Glyph
	err := C.FT_Render_Glyph((*C.FT_GlyphSlotRec)(unsafe.Pointer(glyph)), renderMode)
	return newError(err, "failed to render glyph with index %d for render mode %04x", glyph.GlyphIndex, renderMode)
}

/*
RenderMode is a list of modes supported by FreeType 2.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_render_mode
*/
type RenderMode = C.FT_Render_Mode

const (
	RENDER_MODE_NORMAL = RenderMode(C.FT_RENDER_MODE_NORMAL)
	RENDER_MODE_LIGHT  = RenderMode(C.FT_RENDER_MODE_LIGHT)
	RENDER_MODE_MONO   = RenderMode(C.FT_RENDER_MODE_MONO)
	RENDER_MODE_LCD    = RenderMode(C.FT_RENDER_MODE_LCD)
	RENDER_MODE_LCD_V  = RenderMode(C.FT_RENDER_MODE_LCD_V)
	RENDER_MODE_SDF    = RenderMode(C.FT_RENDER_MODE_SDF)
)

/*
GetKerning returns the kerning vector between two glyphs of the same face.

Use HasKerning to find out whether a font has data that can be extracted with GetKerning.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_get_kerning
*/
func (face Face) GetKerning(leftGlyph UInt, rightGlyph UInt, kernMode KerningMode) (Vector, error) {
	var kerning Vector
	err := C.FT_Get_Kerning(face.face, leftGlyph, rightGlyph, kernMode, (*C.FT_Vector)(unsafe.Pointer(&kerning)))
	return kerning, newError(err, "failed to get kerning for %d and %d with kern mode %d", leftGlyph, rightGlyph, kernMode)
}

/*
KerningMode is an enumeration to specify the format of kerning values returned by GetKerning.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_kerning_mode
*/
type KerningMode = UInt

const (
	KERNING_DEFAULT  = KerningMode(C.FT_KERNING_DEFAULT)
	KERNING_UNFITTED = KerningMode(C.FT_KERNING_UNFITTED)
	KERNING_UNSCALED = KerningMode(C.FT_KERNING_UNSCALED)
)

/*
GetTrackKerning returns the track kerning for a given face object at a given size.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_get_track_kerning
*/
func (face Face) GetTrackKerning(pointSize Fixed, degree Int) (Fixed, error) {
	var kerning Fixed
	err := C.FT_Get_Track_Kerning(face.face, pointSize, degree, (*C.FT_Fixed)(unsafe.Pointer(&kerning)))
	return kerning, newError(err, "failed to get track kerning")
}
