package freetype

// Macros to test various properties of fonts.

// HasHorizontal returns true whenever a face object contains horizontal metrics (this is true for all font formats though).
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_horizontal
func (face Face) HasHorizontal() bool {
	return face.Rec().FaceFlags&FACE_FLAG_HORIZONTAL != 0
}

// HasVertical returns true whenever a face object contains real vertical metrics (and not only synthesized ones).
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_vertical
func (face Face) HasVertical() bool {
	return face.Rec().FaceFlags&FACE_FLAG_VERTICAL != 0
}

// HasKerning returns true whenever a face object contains kerning data that can be accessed with Get_Kerning.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_kerning
func (face Face) HasKerning() bool {
	return face.Rec().FaceFlags&FACE_FLAG_KERNING != 0
}

// HasFixedSizes returns true whenever a face object contains some embedded bitmaps.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_fixed_sizes
func (face Face) HasFixedSizes() bool {
	return face.Rec().FaceFlags&FACE_FLAG_FIXED_SIZES != 0
}

// HasGlyphNames returns true whenever a face object contains some glyph names that can be accessed through Get_Glyph_Name.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_glyph_names
func (face Face) HasGlyphNames() bool {
	return face.Rec().FaceFlags&FACE_FLAG_GLYPH_NAMES != 0
}

// HasColor returns true whenever a face object contains tables for color glyphs.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_color
func (face Face) HasColor() bool {
	return face.Rec().FaceFlags&FACE_FLAG_COLOR != 0
}

// HasMultipleMasters returns true whenever a face object contains some multiple masters.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_multiple_masters
func (face Face) HasMultipleMasters() bool {
	return face.Rec().FaceFlags&FACE_FLAG_MULTIPLE_MASTERS != 0
}

// HasSVG returns true whenever a face object contains an ‘SVG ’ OpenType table.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_svg
func (face Face) HasSVG() bool {
	return face.Rec().FaceFlags&FACE_FLAG_SVG != 0
}

// HasSbix returns true whenever a face object contains an ‘sbix’ OpenType table and outline glyphs.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_sbix
func (face Face) HasSbix() bool {
	return face.Rec().FaceFlags&FACE_FLAG_SBIX != 0
}

// HasSbixOverlay returns true whenever a face object contains an ‘sbix’ OpenType table with bit 1
// in its flags field set, instructing the application to overlay the bitmap strike with the corresponding
// outline glyph.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_has_sbix_overlay
func (face Face) HasSbixOverlay() bool {
	return face.Rec().FaceFlags&FACE_FLAG_SBIX_OVERLAY != 0
}

// IsSfnt returns true whenever a face object contains a font whose format is based on the SFNT storage scheme.
// This usually means: TrueType fonts, OpenType fonts, as well as SFNT-based embedded bitmap fonts.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_sfnt
func (face Face) IsSfnt() bool {
	return face.Rec().FaceFlags&FACE_FLAG_SFNT != 0
}

// IsScalable returns true whenever a face object contains a scalable font face (true for TrueType, Type 1,
// Type 42, CID, OpenType/CFF, and PFR font formats).
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_scalable
func (face Face) IsScalable() bool {
	return face.Rec().FaceFlags&FACE_FLAG_SCALABLE != 0
}

// IsFixedWidth returns true whenever a face object contains a font face that contains fixed-width
// (or ‘monospace’, ‘fixed-pitch’, etc.) glyphs.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_fixed_width
func (face Face) IsFixedWidth() bool {
	return face.Rec().FaceFlags&FACE_FLAG_FIXED_WIDTH != 0
}

// IsCIDKeyed returns true whenever a face object contains a CID-keyed font.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_cid_keyed
func (face Face) IsCIDKeyed() bool {
	return face.Rec().FaceFlags&FACE_FLAG_CID_KEYED != 0
}

// IsTricky returns true whenever a face represents a ‘tricky’ font.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_tricky
func (face Face) IsTricky() bool {
	return face.Rec().FaceFlags&FACE_FLAG_TRICKY != 0
}

// IsNamedInstance returns true whenever a face object is a named instance of a GX or OpenType variation font.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_named_instance
func (face Face) IsNamedInstance() bool {
	return face.Rec().FaceIndex&0x7FFF0000 != 0
}

// IsVariation returns true whenever a face object has been altered by Set_MM_Design_Coordinates,
// Set_Var_Design_Coordinates, Set_Var_Blend_Coordinates, or Set_MM_WeightVector.
//
// https://freetype.org/freetype2/docs/reference/ft2-font_testing_macros.html#ft_is_variation
func (face Face) IsVariation() bool {
	return face.Rec().FaceFlags&FACE_FLAG_VARIATION != 0
}
