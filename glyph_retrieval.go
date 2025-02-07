package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"unsafe"
)

// Functions to manage glyphs.

/*
GlyphSlot is a handle to a given ‘glyph slot’. A slot is a container that can hold any of the glyphs contained in its parent face.

In other words, each time you call Load_Glyph or Load_Char, the slot's content is erased by the new glyph data, i.e., the glyph's metrics, its image (bitmap or outline), and other control information.

See GlyphSlotRec for the publicly accessible glyph fields.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_glyphslot
*/
type GlyphSlot *GlyphSlotRec

func init() {
	assertSameSize(GlyphSlotRec{}, C.FT_GlyphSlotRec{})
}

/*
GlyphSlotRec is the FreeType root glyph slot class structure. A glyph slot is a container where individual glyphs can be loaded, be they in outline or bitmap format.

If LoadGlyph is called with default flags (see LOAD_DEFAULT) the glyph image is loaded in the glyph slot in its native format (e.g., an outline glyph for TrueType and Type 1 formats). [Since 2.9] The prospective bitmap metrics are calculated according to LOAD_TARGET_XXX and other flags even for the outline glyph, even if LOAD_RENDER is not set.

This image can later be converted into a bitmap by calling RenderGlyph. This function searches the current renderer for the native image's format, then invokes it.

The renderer is in charge of transforming the native image through the slot's face transformation fields, then converting it into a bitmap that is returned in slot->bitmap.

Note that slot->bitmap_left and slot->bitmap_top are also used to specify the position of the bitmap relative to the current pen position (e.g., coordinates (0,0) on the baseline). Of course, slot->format is also changed to GLYPH_FORMAT_BITMAP.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_glyphslotrec
*/
type GlyphSlotRec struct {
	// A handle to the FreeType library instance this slot belongs to.
	Library Library
	// A handle to the parent face object.
	Face Face
	// In some cases (like some font tools), several glyph slots per face object can be a good thing.
	// As this is rare, the glyph slots are listed through a direct, single-linked list using its next field.
	Next GlyphSlot
	// [Since 2.10] The glyph index passed as an argument to LoadGlyph while initializing the glyph slot.
	GlyphIndex UInt
	// A typeless pointer unused by the FreeType library or any of its drivers. It can be used by client
	// applications to link their own data to each glyph slot object.
	Generic Generic

	// The metrics of the last loaded glyph in the slot. The returned values depend on the last load flags
	// (see the LoadGlyph API function) and can be expressed either in 26.6 fractional pixels or font units.
	//
	// Note that even when the glyph image is transformed, the metrics are not.
	Metrics GlyphMetrics
	// The advance width of the unhinted glyph. Its value is expressed in 16.16 fractional pixels,unless
	// LOAD_LINEAR_DESIGN is set when loading the glyph. This field can be important to perform correct
	// WYSIWYG layout. Only relevant for scalable glyphs.
	LinearHoriAdvance Fixed
	// The advance height of the unhinted glyph. Its value is expressed in 16.16 fractional pixels, unless
	// LOAD_LINEAR_DESIGN is set when loading the glyph. This field can be important to perform correct
	// WYSIWYG layout. Only relevant for scalable glyphs.
	LinearVertAdvance Fixed
	// This shorthand is, depending on LOAD_IGNORE_TRANSFORM, the transformed (hinted) advance width for the glyph,
	// in 26.6 fractional pixel format. As specified with LOAD_VERTICAL_LAYOUT, it uses either the horiAdvance
	// or the vertAdvance value of metrics field.
	Advance Vector

	// This field indicates the format of the image contained in the glyph slot. Typically GLYPH_FORMAT_BITMAP,
	// GLYPH_FORMAT_OUTLINE, or GLYPH_FORMAT_COMPOSITE, but other values are possible.
	Format GlyphFormat

	// This field is used as a bitmap descriptor. Note that the address and content of the bitmap buffer
	// can change between calls of LoadGlyph and a few other functions.
	Bitmap Bitmap
	// The bitmap's left bearing expressed in integer pixels.
	BitmapLeft Int
	// The bitmap's top bearing expressed in integer pixels. This is the distance from the baseline to the
	// top-most glyph scanline, upwards y coordinates being positive.
	BitmapTop Int

	// The outline descriptor for the current glyph image if its format is GLYPH_FORMAT_OUTLINE. Once a
	// glyph is loaded, outline can be transformed, distorted, emboldened, etc. However, it must not be freed.
	//
	// [Since 2.10.1] If LOAD_NO_SCALE is set, outline coordinates of OpenType variation fonts for a
	// selected instance are internally handled as 26.6 fractional font units but returned as (rounded)
	// integers, as expected. To get unrounded font units, don't use LOAD_NO_SCALE but load the glyph with
	// LOAD_NO_HINTING and scale it, using the font's units_per_EM value as the ppem.
	Outline Outline

	// The number of subglyphs in a composite glyph. This field is only valid for the composite glyph
	// format that should normally only be loaded with the LOAD_NO_RECURSE flag.
	NumSubglyphs UInt
	// An array of subglyph descriptors for composite glyphs. There are num_subglyphs elements in there.
	// Currently internal to FreeType.
	subglyphs unsafe.Pointer

	control_data unsafe.Pointer
	control_len  C.long

	// The difference between hinted and unhinted left side bearing while auto-hinting is active.
	// Zero otherwise.
	LsbDelta Pos
	// The difference between hinted and unhinted right side bearing while auto-hinting is active.
	// Zero otherwise.
	RsbDelta Pos

	other    unsafe.Pointer
	internal unsafe.Pointer
}

func init() {
	assertSameSize(GlyphMetrics{}, C.FT_Glyph_Metrics{})
}

/*
GlyphMetrics is a structure to model the metrics of a single glyph. The values are expressed in 26.6 fractional pixel format; if the flag LOAD_NO_SCALE has been used while loading the glyph, values are expressed in font units instead.

If not disabled with LOAD_NO_HINTING, the values represent dimensions of the hinted glyph (in case hinting is applicable).

Stroking a glyph with an outside border does not increase horiAdvance or vertAdvance; you have to manually adjust these values to account for the added width and height.

FreeType doesn't use the ‘VORG’ table data for CFF fonts because it doesn't have an interface to quickly retrieve the glyph height. The y coordinate of the vertical origin can be simply computed as vertBearingY + height after loading a glyph.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_glyph_metrics
*/
type GlyphMetrics struct {
	// The glyph's width.
	Width Pos
	// The glyph's height.
	Height Pos

	// Left side bearing for horizontal layout.
	HoriBearingX Pos
	// Top side bearing for horizontal layout.
	HoriBearingY Pos
	// Advance width for horizontal layout.
	HoriAdvance Pos

	// Left side bearing for vertical layout.
	VertBearingX Pos
	// Top side bearing for vertical layout. Larger positive values mean further below the vertical glyph origin.
	VertBearingY Pos
	// Advance height for vertical layout. Positive values mean the glyph has a positive advance downward.
	VertAdvance Pos
}

/*
LoadGlyph loads a glyph into the glyph slot of a face object.

  - glyphIndex	- The index of the glyph in the font file. For CID-keyed fonts (either in PS or in CFF format) this argument specifies the CID value.
  - loadFlag - A flag indicating what to load for this glyph. The LOAD_XXX flags can be used to control the glyph loading process (e.g., whether the outline should be scaled, whether to load bitmaps or not, whether to hint the outline, etc).

For proper scaling and hinting, the active Size object owned by the face has to be meaningfully initialized by calling SetCharSize before this function, for example. The loaded glyph may be transformed. See SetTransform for the details.

For subsetted CID-keyed fonts, FT_Err_Invalid_Argument is returned for invalid CID values (that is, for CID values that don't have a corresponding glyph in the font). See the discussion of the FACE_FLAG_CID_KEYED flag for more details.

If you receive Err_Glyph_Too_Big, try getting the glyph outline at EM size, then scale it manually and fill it as a graphics operation.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_load_glyph
*/
func (face Face) LoadGlyph(glyphIndex UInt, loadFlags Int32) error {
	err := C.FT_Load_Glyph(face.face, glyphIndex, loadFlags)
	return newError(err, "failed to load glyph index %d with flags %04x", glyphIndex, loadFlags)
}

/*
LoadFlag is a list of bit field constants for LoadGlyph to indicate what kind of operations to perform during glyph loading.

By default, hinting is enabled and the font's native hinter (see FACE_FLAG_HINTER) is preferred over the auto-hinter. You can disable hinting by setting LOAD_NO_HINTING or change the precedence by setting LOAD_FORCE_AUTOHINT. You can also set LOAD_NO_AUTOHINT in case you don't want the auto-hinter to be used at all.

See the description of FACE_FLAG_TRICKY for a special exception (affecting only a handful of Asian fonts).

Besides deciding which hinter to use, you can also decide which hinting algorithm to use. See LOAD_TARGET_XXX for details.

Note that the auto-hinter needs a valid Unicode cmap (either a native one or synthesized by FreeType) for producing correct results. If a font provides an incorrect mapping (for example, assigning the character code U+005A, LATIN CAPITAL LETTER Z, to a glyph depicting a mathematical integral sign), the auto-hinter might produce useless results.

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
LoadTarget is a list of values to select a specific hinting algorithm for the hinter. You should OR one of these values to your load_flags when calling LoadGlyph.

Note that a font's native hinters may ignore the hinting algorithm you have specified (e.g., the TrueType bytecode interpreter). You can set LOAD_FORCE_AUTOHINT to ensure that the auto-hinter is used.

You should use only one of the LOAD_TARGET_XXX values in your load_flags. They can't be ORed.

If LOAD_RENDER is also set, the glyph is rendered in the corresponding mode (i.e., the mode that matches the used algorithm best). An exception is LOAD_TARGET_MONO since it implies LOAD_MONOCHROME.

You can use a hinting algorithm that doesn't correspond to the same rendering mode. As an example, it is possible to use the ‘light’ hinting algorithm and have the results rendered in horizontal LCD pixel mode, with code like

	FT_Load_Glyph( face, glyph_index,
	               load_flags | FT_LOAD_TARGET_LIGHT );

	FT_Render_Glyph( face->glyph, FT_RENDER_MODE_LCD );

In general, you should stick with one rendering mode. For example, switching between LOAD_TARGET_NORMAL and LOAD_TARGET_MONO enforces a lot of recomputation for TrueType fonts, which is slow. Another reason is caching: Selecting a different mode usually causes changes in both the outlines and the rasterized bitmaps; it is thus necessary to empty the cache after a mode switch to avoid false hits.

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
RenderGlyph convert a given glyph image to a bitmap. It does so by inspecting the glyph image format, finding the relevant renderer, and invoking it.

  - render_mode	- The render mode used to render the glyph image into a bitmap. See FT_Render_Mode for a list of possible values.
    If FT_RENDER_MODE_NORMAL is used, a previous call of FT_Load_Glyph with flag FT_LOAD_COLOR makes FT_Render_Glyph provide a default blending of colored glyph layers associated with the current glyph slot (provided the font contains such layers) instead of rendering the glyph slot's outline. This is an experimental feature; see FT_LOAD_COLOR for more information.

When FreeType outputs a bitmap of a glyph, it really outputs an alpha coverage map. If a pixel is completely covered by a filled-in outline, the bitmap contains 0xFF at that pixel, meaning that 0xFF/0xFF fraction of that pixel is covered, meaning the pixel is 100% black (or 0% bright). If a pixel is only 50% covered (value 0x80), the pixel is made 50% black (50% bright or a middle shade of grey). 0% covered means 0% black (100% bright or white).

On high-DPI screens like on smartphones and tablets, the pixels are so small that their chance of being completely covered and therefore completely black are fairly good. On the low-DPI screens, however, the situation is different. The pixels are too large for most of the details of a glyph and shades of gray are the norm rather than the exception.

This is relevant because all our screens have a second problem: they are not linear. 1 + 1 is not 2. Twice the value does not result in twice the brightness. When a pixel is only 50% covered, the coverage map says 50% black, and this translates to a pixel value of 128 when you use 8 bits per channel (0-255). However, this does not translate to 50% brightness for that pixel on our sRGB and gamma 2.2 screens. Due to their non-linearity, they dwell longer in the darks and only a pixel value of about 186 results in 50% brightness – 128 ends up too dark on both bright and dark backgrounds. The net result is that dark text looks burnt-out, pixely and blotchy on bright background, bright text too frail on dark backgrounds, and colored text on colored background (for example, red on green) seems to have dark halos or ‘dirt’ around it. The situation is especially ugly for diagonal stems like in ‘w’ glyph shapes where the quality of FreeType's anti-aliasing depends on the correct display of grays. On high-DPI screens where smaller, fully black pixels reign supreme, this doesn't matter, but on our low-DPI screens with all the gray shades, it does. 0% and 100% brightness are the same things in linear and non-linear space, just all the shades in-between aren't.

The blending function for placing text over a background is

	dst = alpha * src + (1 - alpha) * dst    ,

which is known as the OVER operator.

To correctly composite an anti-aliased pixel of a glyph onto a surface,

take the foreground and background colors (e.g., in sRGB space) and apply gamma to get them in a linear space,

use OVER to blend the two linear colors using the glyph pixel as the alpha value (remember, the glyph bitmap is an alpha coverage bitmap), and

apply inverse gamma to the blended pixel and write it back to the image.

Internal testing at Adobe found that a target inverse gamma of 1.8 for step 3 gives good results across a wide range of displays with an sRGB gamma curve or a similar one.

This process can cost performance. There is an approximation that does not need to know about the background color; see https://bel.fi/alankila/lcd/ and https://bel.fi/alankila/lcd/alpcor.html for details.

ATTENTION: Linear blending is even more important when dealing with subpixel-rendered glyphs to prevent color-fringing! A subpixel-rendered glyph must first be filtered with a filter that gives equal weight to the three color primaries and does not exceed a sum of 0x100, see section ‘Subpixel Rendering’. Then the only difference to gray linear blending is that subpixel-rendered linear blending is done 3 times per pixel: red foreground subpixel to red background subpixel and so on for green and blue.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_render_glyph
*/
func (face Face) RenderGlyph(renderMode RenderMode) error {
	glyph := face.Rec().Glyph
	err := C.FT_Render_Glyph((*C.FT_GlyphSlotRec)(unsafe.Pointer(glyph)), renderMode)
	return newError(err, "failed to render glyph with index %d for render mode %04x", glyph.GlyphIndex, renderMode)
}

/*
RenderMode is a list of modes supported by FreeType 2. Each mode corresponds to a specific type of scanline conversion performed on the outline.

For bitmap fonts and embedded bitmaps the bitmap->pixel_mode field in the GlyphSlotRec structure gives the format of the returned bitmap.

All modes except RENDER_MODE_MONO use 256 levels of opacity, indicating pixel coverage. Use linear alpha blending and gamma correction to correctly render non-monochrome glyph bitmaps onto a surface; see Render_Glyph.

The RENDER_MODE_SDF is a special render mode that uses up to 256 distance values, indicating the signed distance from the grid position to the nearest outline.

The selected render mode only affects scalable vector glyphs of a font. Embedded bitmaps often have a different pixel mode like PIXEL_MODE_MONO. You can use FT_Bitmap_Convert to transform them into 8-bit pixmaps.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_render_mode
*/
type RenderMode = C.FT_Render_Mode

const (
	// Default render mode; it corresponds to 8-bit anti-aliased bitmaps.
	RENDER_MODE_NORMAL = RenderMode(C.FT_RENDER_MODE_NORMAL)
	// This is equivalent to FT_RENDER_MODE_NORMAL. It is only defined as a separate value because
	// render modes are also used indirectly to define hinting algorithm selectors.
	// See LOAD_TARGET_XXX for details.
	RENDER_MODE_LIGHT = RenderMode(C.FT_RENDER_MODE_LIGHT)
	// This mode corresponds to 1-bit bitmaps (with 2 levels of opacity).
	RENDER_MODE_MONO = RenderMode(C.FT_RENDER_MODE_MONO)
	// This mode corresponds to horizontal RGB and BGR subpixel displays like LCD screens.
	// It produces 8-bit bitmaps that are 3 times the width of the original glyph outline in pixels,
	// and which use the PIXEL_MODE_LCD mode.
	RENDER_MODE_LCD = RenderMode(C.FT_RENDER_MODE_LCD)
	// This mode corresponds to vertical RGB and BGR subpixel displays (like PDA screens, rotated LCD
	// displays, etc.). It produces 8-bit bitmaps that are 3 times the height of the original glyph
	// outline in pixels and use the PIXEL_MODE_LCD_V mode.
	RENDER_MODE_LCD_V = RenderMode(C.FT_RENDER_MODE_LCD_V)
	// The positive (unsigned) 8-bit bitmap values can be converted to the single-channel signed
	// distance field (SDF) by subtracting 128, with the positive and negative results corresponding
	// to the inside and the outside of a glyph contour, respectively. The distance units are
	// arbitrarily determined by an adjustable spread property.
	RENDER_MODE_SDF = RenderMode(C.FT_RENDER_MODE_SDF)
)

/*
GetKerning returns the kerning vector between two glyphs of the same face.

  - leftGlyph - The index of the left glyph in the kern pair.
  - rightGlyph - The index of the right glyph in the kern pair.
  - kernMode - See Kerning_Mode for more information. Determines the scale and dimension of the returned kerning vector.

Returns the kerning vector. This is either in font units, fractional pixels (26.6 format), or pixels for scalable formats, and in pixels for fixed-sizes formats.

Only horizontal layouts (left-to-right & right-to-left) are supported by this method. Other layouts, or more sophisticated kernings, are out of the scope of this API function – they can be implemented through format-specific interfaces.

Note that, for TrueType fonts only, this can extract data from both the ‘kern’ table and the basic, pair-wise kerning feature from the GPOS table (with TT_CONFIG_OPTION_GPOS_KERNING enabled), though FreeType does not support the more advanced GPOS layout features; use a library like HarfBuzz for those instead. If a font has both a ‘kern’ table and kern features of a GPOS table, the ‘kern’ table will be used.

Also note for right-to-left scripts, the functionality may differ for fonts with GPOS tables vs. ‘kern’ tables. For GPOS, right-to-left fonts typically use both a placement offset and an advance for pair positioning, which this API does not support, so it would output kerning values of zero; though if the right-to-left font used only advances in GPOS pair positioning, then this API could output kerning values for it, but it would use left_glyph to mean the first glyph for that case. Whereas ‘kern’ tables are always advance-only and always store the left glyph first.

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

KERNING_DEFAULT returns full pixel values; it also makes FreeType heuristically scale down kerning distances at small ppem values so that they don't become too big.

Both KERNING_DEFAULT and KERNING_UNFITTED use the current horizontal scaling factor (as set e.g. with SetCharSize) to convert font units to pixels.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_kerning_mode
*/
type KerningMode = UInt

const (
	// Return grid-fitted kerning distances in 26.6 fractional pixels.
	KERNING_DEFAULT = KerningMode(C.FT_KERNING_DEFAULT)
	// Return un-grid-fitted kerning distances in 26.6 fractional pixels.
	KERNING_UNFITTED = KerningMode(C.FT_KERNING_UNFITTED)
	// Return the kerning vector in original font units.
	KERNING_UNSCALED = KerningMode(C.FT_KERNING_UNSCALED)
)

/*
GetTrackKerning returns the track kerning for a given face object at a given size.

  - point_size - The point size in 16.16 fractional points.
  - degree - The degree of tightness. Increasingly negative values represent tighter track kerning, while increasingly positive values represent looser track kerning. Value zero means no track kerning.

Returns the kerning in 16.16 fractional points, to be uniformly applied between all glyphs.

Currently, only the Type 1 font driver supports track kerning, using data from AFM files (if attached with AttachFile or AttachStream).

Only very few AFM files come with track kerning data; please refer to Adobe's AFM specification for more details.

https://freetype.org/freetype2/docs/reference/ft2-glyph_retrieval.html#ft_get_track_kerning
*/
func (face Face) GetTrackKerning(pointSize Fixed, degree Int) (Fixed, error) {
	var kerning Fixed
	err := C.FT_Get_Track_Kerning(face.face, pointSize, degree, (*C.FT_Fixed)(unsafe.Pointer(&kerning)))
	return kerning, newError(err, "failed to get track kerning")
}
