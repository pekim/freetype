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

// Macros to test various properties of fonts.

// HasHorizontal returns true whenever a face object contains horizontal metrics (this is true for all font formats though).
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_horizontal
func (face Face) HasHorizontal() bool {
	return cBoolToGo(C.c_FT_HAS_HORIZONTAL(face.face))
}

// HasVertical returns true whenever a face object contains real vertical metrics (and not only synthesized ones).
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_vertical
func (face Face) HasVertical() bool {
	return cBoolToGo(C.c_FT_HAS_VERTICAL(face.face))
}

// HasKerning returns true whenever a face object contains kerning data that can be accessed with Get_Kerning.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_kerning
func (face Face) HasKerning() bool {
	return cBoolToGo(C.c_FT_HAS_KERNING(face.face))
}

// HasFixedSizes returns true whenever a face object contains some embedded bitmaps.
// See the available_sizes field of the FaceRec structure.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_fixed_sizes
func (face Face) HasFixedSizes() bool {
	return cBoolToGo(C.c_FT_HAS_FIXED_SIZES(face.face))
}

// HasGlyphNames returns true whenever a face object contains some glyph names that can be accessed through Get_Glyph_Name.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_glyph_names
func (face Face) HasGlyphNames() bool {
	return cBoolToGo(C.c_FT_HAS_GLYPH_NAMES(face.face))
}

// HasColor returns true whenever a face object contains tables for color glyphs.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_color
func (face Face) HasColor() bool {
	return cBoolToGo(C.c_FT_HAS_COLOR(face.face))
}

// HasMultipleMasters returns true whenever a face object contains some multiple masters.
// The functions provided by FT_MULTIPLE_MASTERS_H are then available to choose the exact design you want.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_multiple_masters
func (face Face) HasMultipleMasters() bool {
	return cBoolToGo(C.c_FT_HAS_MULTIPLE_MASTERS(face.face))
}

// HasSVG returns true whenever a face object contains an ‘SVG ’ OpenType table.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_svg
func (face Face) HasSVG() bool {
	return cBoolToGo(C.c_FT_HAS_SVG(face.face))
}

// HasSbix returns true whenever a face object contains an ‘sbix’ OpenType table and outline glyphs.
//
// Currently, FreeType only supports bitmap glyphs in PNG format for this table
// (i.e., JPEG and TIFF formats are unsupported,
// as are Apple-specific formats not part of the OpenType specification).
//
// For backward compatibility, a font with an ‘sbix’ table is treated as a bitmap-only face.
// Using FT_Open_Face with FT_PARAM_TAG_IGNORE_SBIX, an application can switch off ‘sbix’ handling
// so that the face is treated as an ordinary outline font with scalable outlines.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_sbix
func (face Face) HasSbix() bool {
	return cBoolToGo(C.c_FT_HAS_SBIX(face.face))
}

// HasSbixOverlay returns true whenever a face object contains an ‘sbix’ OpenType table with bit 1
// in its flags field set, instructing the application to overlay the bitmap strike with the corresponding
// outline glyph. See HasSbix for pseudo code how to use it.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_sbix_overlay
func (face Face) HasSbixOverlay() bool {
	return cBoolToGo(C.c_FT_HAS_SBIX_OVERLAY(face.face))
}

// IsSFNT returns true whenever a face object contains a font whose format is based on the SFNT storage scheme.
// This usually means: TrueType fonts, OpenType fonts, as well as SFNT-based embedded bitmap fonts.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_sfnt
func (face Face) IsSFNT() bool {
	return cBoolToGo(C.c_FT_IS_SFNT(face.face))
}

// IsScalable returns true whenever a face object contains a scalable font face (true for TrueType, Type 1,
// Type 42, CID, OpenType/CFF, and PFR font formats).
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_scalable
func (face Face) IsScalable() bool {
	return cBoolToGo(C.c_FT_IS_SCALABLE(face.face))
}

// IsFixedWidth returns true whenever a face object contains a font face that contains fixed-width
// (or ‘monospace’, ‘fixed-pitch’, etc.) glyphs.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_fixed_width
func (face Face) IsFixedWidth() bool {
	return cBoolToGo(C.c_FT_IS_FIXED_WIDTH(face.face))
}

// IsCIDKeyed returns true whenever a face object contains a CID-keyed font.
// See the discussion of FACE_FLAG_CID_KEYED for more details.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_cid_keyed
func (face Face) IsCIDKeyed() bool {
	return cBoolToGo(C.c_FT_IS_CID_KEYED(face.face))
}

// IsTricky returns true whenever a face represents a ‘tricky’ font.
// See the discussion of FT_FACE_FLAG_TRICKY for more details.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_tricky
func (face Face) IsTricky() bool {
	return cBoolToGo(C.c_FT_IS_TRICKY(face.face))
}

// IsNamedInstance returns true whenever a face object is a named instance of a GX or OpenType variation font.
//
// [Since 2.9] Changing the design coordinates with FT_Set_Var_Design_Coordinates or
// FT_Set_Var_Blend_Coordinates does not influence the return value of this macro
// (only Set_Named_Instance does that).
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_named_instance
func (face Face) IsNamedInstance() bool {
	return cBoolToGo(C.c_FT_IS_NAMED_INSTANCE(face.face))
}

// IsVariation returns true whenever a face object has been altered by Set_MM_Design_Coordinates,
// Set_Var_Design_Coordinates, Set_Var_Blend_Coordinates, or Set_MM_WeightVector.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_variation
func (face Face) IsVariation() bool {
	return cBoolToGo(C.c_FT_IS_VARIATION(face.face))
}
