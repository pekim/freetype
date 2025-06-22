package freetype

import (
	"modernc.org/libfreetype"
)

// Functions to manage glyphs.

/*
GlyphSlot is a handle to a given ‘glyph slot’.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_glyphslot
*/
type GlyphSlot uintptr

// Rec returns a pointer to the FaceRec that is referenced by the Face.
func (glyphSlot GlyphSlot) Rec() *GlyphSlotRec {
	return fromUintptr[GlyphSlotRec](uintptr(glyphSlot))
}

func init() {
	assertSameSize(GlyphSlotRec{}, libfreetype.TFT_GlyphSlotRec{})
}

/*
GlyphSlotRec is the FreeType root glyph slot class structure.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_glyphslotrec
*/
type GlyphSlotRec struct {
	Library    libfreetype.TFT_Library
	Face       libfreetype.TFT_Face
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

	Outline libfreetype.TFT_Outline

	NumSubglyphs UInt
	_            uintptr // subglyphs

	_ uintptr // control_data
	_ int64   // control_len

	LsbDelta Pos
	RsbDelta Pos

	Other uintptr
	_     uintptr // internal
}

func init() {
	assertSameSize(GlyphMetrics{}, libfreetype.TFT_Glyph_Metrics{})
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
	err := libfreetype.XFT_Load_Glyph(face.tls, face.face, glyphIndex, loadFlags)
	return newError(err, "failed to load glyph index %d with flags %04x", glyphIndex, loadFlags)
}

/*
LoadFlag is a list of bit field constants for LoadGlyph to indicate what kind of operations to perform during glyph loading.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_load_xxx
*/
type LoadFlag = Int32

const (
	LOAD_DEFAULT                     = LoadFlag(0x0)
	LOAD_NO_SCALE                    = LoadFlag(1 << 1)
	LOAD_NO_HINTING                  = LoadFlag(1 << 2)
	LOAD_RENDER                      = LoadFlag(1 << 3)
	LOAD_NO_BITMAP                   = LoadFlag(1 << 4)
	LOAD_VERTICAL_LAYOUT             = LoadFlag(1 << 5)
	LOAD_FORCE_AUTOHINT              = LoadFlag(1 << 6)
	LOAD_CROP_BITMAP                 = LoadFlag(1 << 7)
	LOAD_PEDANTIC                    = LoadFlag(1 << 8)
	LOAD_IGNORE_GLOBAL_ADVANCE_WIDTH = LoadFlag(1 << 9)
	LOAD_NO_RECURSE                  = LoadFlag(1 << 10)
	LOAD_IGNORE_TRANSFORM            = LoadFlag(1 << 11)
	LOAD_MONOCHROME                  = LoadFlag(1 << 12)
	LOAD_LINEAR_DESIGN               = LoadFlag(1 << 13)
	LOAD_SBITS_ONLY                  = LoadFlag(1 << 14)
	LOAD_NO_AUTOHINT                 = LoadFlag(1 << 15)

	/* Bits 16-19 are used by `FT_LOAD_TARGET_` */
	LOAD_COLOR               = LoadFlag(1 << 20)
	LOAD_COMPUTE_METRICS     = LoadFlag(1 << 21)
	LOAD_BITMAP_METRICS_ONLY = LoadFlag(1 << 22)
	LOAD_NO_SVG              = LoadFlag(1 << 24)
)

/*
LoadTarget is a list of values to select a specific hinting algorithm for the hinter.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_load_target_xxx
*/
type LoadTarget = Int32

const (
	LOAD_TARGET_NORMAL = LoadTarget(RENDER_MODE_NORMAL)
	LOAD_TARGET_LIGHT  = LoadTarget(RENDER_MODE_LIGHT)
	LOAD_TARGET_MONO   = LoadTarget(RENDER_MODE_MONO)
	LOAD_TARGET_LCD    = LoadTarget(RENDER_MODE_LCD)
	LOAD_TARGET_LCD_V  = LoadTarget(RENDER_MODE_LCD_V)
)

/*
RenderGlyph convert a given glyph image to a bitmap.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_render_glyph
*/
func (face Face) RenderGlyph(renderMode RenderMode) error {
	glyph := face.Rec().Glyph
	err := libfreetype.XFT_Render_Glyph(face.tls, libfreetype.TFT_GlyphSlot(glyph), renderMode)
	return newError(err, "failed to render glyph with index %d for render mode %04x", glyph.Rec().GlyphIndex, renderMode)
}

/*
RenderMode is a list of modes supported by FreeType 2.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_render_mode
*/
type RenderMode = libfreetype.TFT_Render_Mode

const (
	RENDER_MODE_NORMAL = RenderMode(0)
	RENDER_MODE_LIGHT  = RenderMode(1)
	RENDER_MODE_MONO   = RenderMode(2)
	RENDER_MODE_LCD    = RenderMode(3)
	RENDER_MODE_LCD_V  = RenderMode(4)
	RENDER_MODE_SDF    = RenderMode(5)
)

/*
GetKerning returns the kerning vector between two glyphs of the same face.

Use HasKerning to find out whether a font has data that can be extracted with GetKerning.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_get_kerning
*/
func (face Face) GetKerning(leftGlyph UInt, rightGlyph UInt, kernMode KerningMode) (Vector, error) {
	var kerning Vector
	err := libfreetype.XFT_Get_Kerning(face.tls, face.face, leftGlyph, rightGlyph, kernMode, toUintptr(&kerning))
	return kerning, newError(err, "failed to get kerning for %d and %d with kern mode %d", leftGlyph, rightGlyph, kernMode)
}

/*
KerningMode is an enumeration to specify the format of kerning values returned by GetKerning.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_kerning_mode
*/
type KerningMode = UInt

const (
	KERNING_DEFAULT  = KerningMode(0)
	KERNING_UNFITTED = KerningMode(1)
	KERNING_UNSCALED = KerningMode(2)
)

/*
GetTrackKerning returns the track kerning for a given face object at a given size.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_get_track_kerning
*/
func (face Face) GetTrackKerning(pointSize Fixed, degree Int) (Fixed, error) {
	var kerning Fixed
	err := libfreetype.XFT_Get_Track_Kerning(face.tls, face.face, pointSize, degree, toUintptr(&kerning))
	return kerning, newError(err, "failed to get track kerning")
}
