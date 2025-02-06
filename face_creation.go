package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
//
// #include <freetype/ftparams.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

// Functions to manage fonts.

/*
Face is a handle to a typographic face object.
A face object models a given typeface, in a given style.

A face object also owns a single GlyphSlot object, as well as one or more Size objects.

https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_face
*/
type Face struct {
	face C.FT_Face
}

// Rec returns a pointer to the FaceRec that is referenced by the Face.
func (face Face) Rec() *FaceRec {
	return (*FaceRec)(unsafe.Pointer((face.face)))
}

func init() {
	assertSameSize(FaceRec{}, C.FT_FaceRec{})
}

// FaceRec is a FreeType root face class structure.
// A face object models a typeface in a font file.
//
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_facerec
type FaceRec struct {
	// The number of faces in the font file. Some font formats can have multiple faces in a single font file.
	NumFaces Long
	// This field holds two different values.
	// Bits 0-15 are the index of the face in the font file (starting with value 0).
	// They are set to 0 if there is only one face in the font file.
	//
	// [Since 2.6.1] Bits 16-30 are relevant to GX and OpenType variation fonts only,
	// holding the named instance index for the current face index (starting with value 1;
	// value 0 indicates font access without a named instance). For non-variation fonts,
	// bits 16-30 are ignored. If we have the third named instance of face 4, say,
	// face_index is set to 0x00030004.
	//
	// Bit 31 is always zero (that is, face_index is always a positive value).
	//
	// [Since 2.9] Changing the design coordinates with FT_Set_Var_Design_Coordinates or
	// FT_Set_Var_Blend_Coordinates does not influence the named instance index value
	// (only FT_Set_Named_Instance does that).
	FaceIndex Long

	// A set of bit flags that give important information about the face; see FACE_FLAG_XXX
	// for the details.
	FaceFlags FACE_FLAG
	// The lower 16 bits contain a set of bit flags indicating the style of the face;
	// see FT_STYLE_FLAG_XXX for the details.
	//
	// [Since 2.6.1] Bits 16-30 hold the number of named instances available for the current face
	// if we have a GX or OpenType variation (sub)font.
	// Bit 31 is always zero (that is, style_flags is always a positive value).
	// Note that a variation font has always at least one named instance, namely the default instance.
	StyleFlags STYLE_FLAG

	// The number of glyphs in the face. If the face is scalable and has sbits (see num_fixed_sizes),
	// it is set to the number of outline glyphs.
	//
	// For CID-keyed fonts (not in an SFNT wrapper) this value gives the highest CID used in the font.
	NumGlyphs Long

	family_name *String
	style_name  *String

	num_fixed_sizes Int
	available_sizes *BitmapSize

	num_charmaps Int
	charmaps     **CharMapRec

	generic Generic

	/* The following member variables (down to `underline_thickness`) */
	/* outlines are only relevant to scalable  cf. @FT_Bitmap_Size    */
	/* for bitmap fonts.                                              */

	// The font bounding box. Coordinates are expressed in font units (see units_per_EM).
	// The box is large enough to contain any glyph from the font.
	// Thus, bbox.yMax can be seen as the ‘maximum ascender’, and bbox.yMin as the ‘minimum descender’.
	// Only relevant for scalable formats.
	//
	// Note that the bounding box might be off by (at least) one pixel for hinted fonts.
	// See FT_Size_Metrics for further discussion.
	//
	// Note that the bounding box does not vary in OpenType variation fonts and should only be used in
	// relation to the default instance.
	Bbox BBox

	// The number of font units per EM square for this face.
	// This is typically 2048 for TrueType fonts, and 1000 for Type 1 fonts.
	// Only relevant for scalable formats.
	UnitsPerEM UShort
	// The typographic ascender of the face, expressed in font units.
	// For font formats not having this information, it is set to bbox.yMax.
	// Only relevant for scalable formats.
	Ascender Short
	// The typographic descender of the face, expressed in font units.
	// For font formats not having this information, it is set to bbox.yMin.
	// Note that this field is negative for values below the baseline.
	// Only relevant for scalable formats.
	Descender Short
	// This value is the vertical distance between two consecutive baselines, expressed in font units.
	// It is always positive. Only relevant for scalable formats.
	//
	// If you want the global glyph height, use ascender - descender.
	Height Short

	// The maximum advance width, in font units, for all glyphs in this face.
	// This can be used to make word wrapping computations faster.
	// Only relevant for scalable formats.
	MaxAdvanceWidth Short
	// The maximum advance height, in font units, for all glyphs in this face.
	// This is only relevant for vertical layouts, and is set to height for fonts that do not provide vertical metrics.
	// Only relevant for scalable formats.
	MaxAdvanceHeight Short

	// The position, in font units, of the underline line for this face.
	// It is the center of the underlining stem.
	// Only relevant for scalable formats.
	UnderlinePosition Short
	// The thickness, in font units, of the underline for this face.
	// Only relevant for scalable formats.
	UnderlineThickness Short

	// The face's associated glyph slot(s).
	Glyph *GlyphSlotRec
	// The current active size for this face.
	Size *SizeRec
	// The current active charmap for this face.
	Charmap *CharMapRec

	/* private fields, internal to FreeType */

	driver unsafe.Pointer
	memory unsafe.Pointer
	stream unsafe.Pointer

	sizes_list ListRec

	autohint   Generic        /* face-specific auto-hinter data */
	extensions unsafe.Pointer /* unused                         */

	internal unsafe.Pointer
}

/*
FamilyName returns the face's family name.
This is an ASCII string, usually in English, that describes the typeface's family
(like ‘Times New Roman’, ‘Bodoni’, ‘Garamond’, etc).
This is a least common denominator used to list fonts.
Some formats (TrueType & OpenType) provide localized and Unicode versions of this string.
Applications should use the format-specific interface to access them.
Can be NULL (e.g., in fonts embedded in a PDF file).

In case the font doesn't provide a specific family name entry, FreeType tries to synthesize one,
deriving it from other name entries.

(This exposes the C string referenced by the unexported family_name field.)
*/
func (fr *FaceRec) FamilyName() string {
	return C.GoString(fr.family_name)
}

/*
Stylename returns the face's style name.
This is an ASCII string, usually in English, that describes the typeface's style
(like ‘Italic’, ‘Bold’, ‘Condensed’, etc).
Not all font formats provide a style name, so this field is optional, and can be set to NULL.
As for family_name, some formats provide localized and Unicode versions of this string.
Applications should use the format-specific interface to access them.

(This exposes the C string referenced by the unexported style_name field.)
*/
func (fr *FaceRec) StyleName() string {
	return C.GoString(fr.style_name)
}

/*
AvailableSizes returns a slice of Bitmap_Size for all bitmap strikes in the face.
It is set to NULL if there is no bitmap strike.

Note that FreeType tries to sanitize the strike data since they are sometimes sloppy or incorrect, but this can easily fail.

(This exposes the data referenced by the unexported num_fixed_sizes and available_size fields.)
*/
func (fr *FaceRec) AvailableSizes() []BitmapSize {
	return unsafe.Slice(fr.available_sizes, fr.num_fixed_sizes)
}

/*
Charmaps returns the charmaps of the face.

(This exposes the data referenced by the unexported num_charmaps and charmap fields.)
*/
func (fr *FaceRec) Charmaps() []*CharMapRec {
	return unsafe.Slice(fr.charmaps, fr.num_charmaps)
}

// A list of bit flags used in the face_flags field of the FaceRec structure.
// They inform client applications of properties of the corresponding face.
//
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_face_flag_xxx
type FACE_FLAG = Long

const (
	// The face contains outline glyphs.
	// Note that a face can contain bitmap strikes also, i.e., a face can have both this flag
	// and FACE_FLAG_FIXED_SIZES set.
	FACE_FLAG_SCALABLE = FACE_FLAG(C.FT_FACE_FLAG_SCALABLE)
	// The face contains bitmap strikes.
	// See also the num_fixed_sizes and available_sizes fields of FT_FaceRec.
	FACE_FLAG_FIXED_SIZES = FACE_FLAG(C.FT_FACE_FLAG_FIXED_SIZES)
	// The face contains fixed-width characters (like Courier, Lucida, MonoType, etc.).
	FACE_FLAG_FIXED_WIDTH = FACE_FLAG(C.FT_FACE_FLAG_FIXED_WIDTH)
	// The face uses the SFNT storage scheme. For now, this means TrueType and OpenType.
	FACE_FLAG_SFNT = FACE_FLAG(C.FT_FACE_FLAG_SFNT)
	// The face contains horizontal glyph metrics. This should be set for all common formats.
	FACE_FLAG_HORIZONTAL = FACE_FLAG(C.FT_FACE_FLAG_HORIZONTAL)
	// The face contains vertical glyph metrics. This is only available in some formats, not all of them.
	FACE_FLAG_VERTICAL = FACE_FLAG(C.FT_FACE_FLAG_VERTICAL)
	// The face contains kerning information.
	// If set, the kerning distance can be retrieved using the function FT_Get_Kerning.
	// Otherwise the function always returns the vector (0,0).
	//
	// Note that for TrueType fonts only, FreeType supports both the ‘kern’ table and the basic,
	// pair-wise kerning feature from the ‘GPOS’ table (with TT_CONFIG_OPTION_GPOS_KERNING enabled),
	// though FreeType does not support the more advanced GPOS layout features; use a library like
	// HarfBuzz for those instead.
	FACE_FLAG_KERNING = FACE_FLAG(C.FT_FACE_FLAG_KERNING)
	// THIS FLAG IS DEPRECATED. DO NOT USE OR TEST IT.
	FACE_FLAG_FAST_GLYPHS = FACE_FLAG(C.FT_FACE_FLAG_FAST_GLYPHS)
	// The face contains multiple masters and is capable of interpolating between them.
	// Supported formats are Adobe MM, TrueType GX, and OpenType variation fonts.
	//
	// See section ‘Multiple Masters’ for API details.
	FACE_FLAG_MULTIPLE_MASTERS = FACE_FLAG(C.FT_FACE_FLAG_MULTIPLE_MASTERS)
	// The face contains glyph names, which can be retrieved using FT_Get_Glyph_Name.
	// Note that some TrueType fonts contain broken glyph name tables.
	// Use the function Has_PS_Glyph_Names when needed.
	FACE_FLAG_GLYPH_NAMES = FACE_FLAG(C.FT_FACE_FLAG_GLYPH_NAMES)
	// Used internally by FreeType to indicate that a face's stream was provided by the
	// client application and should not be destroyed when FT_Done_Face is called.
	// Don't read or test this flag.
	FACE_FLAG_EXTERNAL_STREAM = FACE_FLAG(C.FT_FACE_FLAG_EXTERNAL_STREAM)
	// The font driver has a hinting machine of its own.
	// For example, with TrueType fonts, it makes sense to use data from the SFNT ‘gasp’ table only
	// if the native TrueType hinting engine (with the bytecode interpreter) is available and active.
	FACE_FLAG_HINTER = FACE_FLAG(C.FT_FACE_FLAG_HINTER)
	// The face is CID-keyed.
	// In that case, the face is not accessed by glyph indices but by CID values.
	// For subsetted CID-keyed fonts this has the consequence that not all index values are a valid
	// argument to FT_Load_Glyph.
	// Only the CID values for which corresponding glyphs in the subsetted font exist make
	// Load_Glyph return successfully; in all other cases you get an FT_Err_Invalid_Argument error.
	//
	// Note that CID-keyed fonts that are in an SFNT wrapper (that is, all OpenType/CFF fonts) don't
	//  have this flag set since the glyphs are accessed in the normal way (using contiguous indices);
	// the ‘CID-ness’ isn't visible to the application.
	FACE_FLAG_CID_KEYED = FACE_FLAG(C.FT_FACE_FLAG_CID_KEYED)
	// The face is ‘tricky’, that is, it always needs the font format's native hinting engine to get
	// a reasonable result. A typical example is the old Chinese font mingli.ttf (but not mingliu.ttc)
	// that uses TrueType bytecode instructions to move and scale all of its subglyphs.
	//
	// It is not possible to auto-hint such fonts using FT_LOAD_FORCE_AUTOHINT; it will also ignore
	// LOAD_NO_HINTING. You have to set both LOAD_NO_HINTING and LOAD_NO_AUTOHINT to really disable hinting;
	// however, you probably never want this except for demonstration purposes.
	//
	// Currently, there are about a dozen TrueType fonts in the list of tricky fonts;
	// they are hard-coded in file ttobjs.c.
	FACE_FLAG_TRICKY = FACE_FLAG(C.FT_FACE_FLAG_TRICKY)
	// [Since 2.5.1] The face has color glyph tables.
	// See LOAD_COLOR for more information.
	FACE_FLAG_COLOR = FACE_FLAG(C.FT_FACE_FLAG_COLOR)
	// [Since 2.9] Set if the current face (or named instance) has been altered with Set_MM_Design_Coordinates,
	// Set_Var_Design_Coordinates, Set_Var_Blend_Coordinates, or Set_MM_WeightVector to select a non-default instance.
	FACE_FLAG_VARIATION = FACE_FLAG(C.FT_FACE_FLAG_VARIATION)
	// [Since 2.12] The face has an ‘SVG ’ OpenType table.
	FACE_FLAG_SVG = FACE_FLAG(C.FT_FACE_FLAG_SVG)
	// [Since 2.12] The face has an ‘sbix’ OpenType table and outlines. For such fonts,
	// FACE_FLAG_SCALABLE is not set by default to retain backward compatibility.
	FACE_FLAG_SBIX = FACE_FLAG(C.FT_FACE_FLAG_SBIX)
	// [Since 2.12] The face has an ‘sbix’ OpenType table where outlines should be drawn on top of bitmap strikes.
	FACE_FLAG_SBIX_OVERLAY = FACE_FLAG(C.FT_FACE_FLAG_SBIX_OVERLAY)
)

// A list of bit flags to indicate the style of a given face.
// These are used in the style_flags field of FaceRec.
//
// The style information as provided by FreeType is very basic.
// More details are beyond the scope and should be done on a higher level (for example,
// by analyzing various fields of the ‘OS/2’ table in SFNT based fonts).
//
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_style_flag_xxx
type STYLE_FLAG = Long

const (
	// The face style is italic or oblique.
	STYLE_FLAG_ITALIC = STYLE_FLAG(C.FT_STYLE_FLAG_ITALIC)
	// The face is bold.
	STYLE_FLAG_BOLD = STYLE_FLAG(C.FT_STYLE_FLAG_BOLD)
)

// NewFace opens a font by its pathname.
//
//   - filepathname - A path to the font file.
//   - faceIndex - See OpenFace for a detailed description of this parameter.
//
// Use Done method to destroy the created Face object (along with its slot and sizes).
//
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_new_face
func (lib Library) NewFace(filepathname string, faceIndex int) (Face, error) {
	cFilepathname := C.CString(filepathname)
	defer C.free(unsafe.Pointer(cFilepathname))

	face := Face{}
	err := C.FT_New_Face(lib.library, cFilepathname, C.FT_Long(faceIndex), &face.face)
	return face, newError(err, "failed to create a face for file '%s'", filepathname)
}

// Done discards a given face object, as well as all of its child slots and sizes.
//
// See the discussion of reference counters in the description of ReferenceFace.
//
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_done_face
func (face Face) Done() error {
	err := C.FT_Done_Face(face.face)
	return newError(err, "failed to discard face")
}

/*
A counter gets initialized to 1 at the time a Face structure is created.
This function increments the counter.
FT_Done_Face then only destroys a face if the counter is 1, otherwise it simply decrements the counter.

This function helps in managing life-cycles of structures that reference Face objects.

https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_reference_face
*/
func (face Face) Reference() error {
	err := C.FT_Reference_Face(face.face)
	return newError(err, "failed to reference face")
}

// NewMemoryFace opens a font that has been loaded into memory.
//
//   - data - the font's data
//   - faceIndex	- See OpenFace for a detailed description of this parameter.
//
// You must not deallocate the memory before calling Face.Done.
//
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_new_memory_face
func (lib Library) NewMemoryFace(data []byte, faceIndex int) (Face, error) {
	face := Face{}
	err := C.FT_New_Memory_Face(lib.library, (*C.FT_Byte)(unsafe.Pointer(&data[0])), C.FT_Long(len(data)), C.FT_Long(faceIndex), &face.face)
	return face, newError(err, "failed to create a new memory face")
}

/*
Properties sets or overrides certain (library or module-wide) properties on a face-by-face basis.
Useful for finer-grained control and avoiding locks on shared structures (threads can modify their own faces as they see fit).

Contrary to PropertySet, this function uses Parameter so that you can pass multiple properties to the target face in one call.
Note that only a subset of the available properties can be controlled.

  - PARAM_TAG_STEM_DARKENING - (stem darkening, corresponding to the property no-stem-darkening provided by the ‘autofit’, ‘cff’, ‘type1’, and ‘t1cid’ modules; see no-stem-darkening).
  - PARAM_TAG_LCD_FILTER_WEIGHTS - (LCD filter weights, corresponding to function FT_Library_SetLcdFilterWeights).
  - PARAM_TAG_RANDOM_SEED - (seed value for the CFF, Type 1, and CID ‘random’ operator, corresponding to the random-seed property provided by the ‘cff’, ‘type1’, and ‘t1cid’ modules; see random-seed).

Pass NULL as data in Parameter for a given tag to reset the option and use the library or module default again.

https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_face_properties
*/
func (face Face) Properties(properties ...Parameter) error {
	err := C.FT_Face_Properties(face.face, C.FT_UInt(len(properties)), (*C.FT_Parameter)(&properties[0]))
	for _, param := range properties {
		param.freeData()
	}
	return newError(err, "failed to set face properties")
}

// FT_Open_Face
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_open_face

// FT_Open_Args
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_open_args

// FT_OPEN_XXX
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_open_xxx

/*
Parameter is a simple structure to pass more or less generic parameters to Library.OpenFace and Face.Properties.

Use one of the ParameterTagXXX functions to create a Parameter initialized to a specific parameter tag with data.
Pass a nil argument to reset the property that the parameter represents.
*/
type Parameter = C.FT_Parameter

func (param Parameter) freeData() {
	C.free(unsafe.Pointer(param.data))
}

// ParameterTagIgnoreTypoGraphicFamily creates a Parameter for the FT_PARAM_TAG_IGNORE_TYPOGRAPHIC_FAMILY tag.
//
// https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_ignore_typographic_family
func ParameterTagIgnoreTypoGraphicFamily(value *bool) Parameter {
	return booleanParamTag(C.FT_PARAM_TAG_IGNORE_TYPOGRAPHIC_FAMILY, value)
}

// ParameterTagIgnoreTypoGraphicSubfamily creates a Parameter for the FT_PARAM_TAG_IGNORE_TYPOGRAPHIC_SUBFAMILY tag.
//
// https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_ignore_typographic_subfamily
func ParameterTagIgnoreTypoGraphicSubfamily(value *bool) Parameter {
	return booleanParamTag(C.FT_PARAM_TAG_IGNORE_TYPOGRAPHIC_SUBFAMILY, value)
}

// ParameterTagIncremental creates a Parameter for the FT_PARAM_TAG_INCREMENTAL tag.
//
// https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_incremental
func ParameterTagIncremental(value *bool) Parameter {
	return booleanParamTag(C.FT_PARAM_TAG_INCREMENTAL, value)
}

// ParameterTagIgnoreSbix creates a Parameter for the FT_PARAM_TAG_IGNORE_SBIX tag.
//
// https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_ignore_sbix
func ParameterTagIgnoreSbix(value *bool) Parameter {
	return booleanParamTag(C.FT_PARAM_TAG_IGNORE_SBIX, value)
}

const LCDFilterWeightsLen = 5

// ParameterTagLCDFilterWeights creates a Parameter for the FT_PARAM_TAG_LCD_FILTER_WEIGHTS tag.
//
// https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_lcd_filter_weights
func ParameterTagLCDFilterWeights(weights *[LCDFilterWeightsLen]byte) Parameter {
	if weights == nil {
		return C.FT_Parameter{
			tag:  C.FT_PARAM_TAG_LCD_FILTER_WEIGHTS,
			data: nil,
		}
	}

	cWeights := (*C.uchar)(C.malloc(LCDFilterWeightsLen))
	C.memcpy(unsafe.Pointer(cWeights), unsafe.Pointer(&(*weights)[0]), LCDFilterWeightsLen)

	return C.FT_Parameter{
		tag:  C.FT_PARAM_TAG_LCD_FILTER_WEIGHTS,
		data: C.FT_Pointer(cWeights),
	}
}

// ParameterTagRandomSeed creates a Parameter for the FT_PARAM_TAG_RANDOM_SEED tag.
//
// https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_random_seed
func ParameterTagRandomSeed(value *int) Parameter {
	return integerParamTag(C.FT_PARAM_TAG_RANDOM_SEED, value)
}

// ParameterTagStemDarkening creates a Parameter for the FT_PARAM_TAG_STEM_DARKENING tag.
//
// https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_stem_darkening
func ParameterTagStemDarkening(value *bool) Parameter {
	return booleanParamTag(C.FT_PARAM_TAG_STEM_DARKENING, value)
}

// ParameterTagUnpatentedHinting creates a Parameter for the FT_PARAM_TAG_UNPATENTED_HINTING tag.
//
// https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_unpatented_hinting
func ParameterTagUnpatentedHinting(value *bool) Parameter {
	return booleanParamTag(C.FT_PARAM_TAG_UNPATENTED_HINTING, value)
}

func booleanParamTag(tag C.FT_ULong, value *bool) Parameter {
	if value == nil {
		return C.FT_Parameter{
			tag:  tag,
			data: nil,
		}
	}

	var cBool C.FT_Bool
	cValue := (*C.FT_Bool)(C.malloc(C.size_t(unsafe.Sizeof(cBool))))
	if *value {
		*cValue = 1
	} else {
		*cValue = 0
	}

	return C.FT_Parameter{
		tag:  tag,
		data: C.FT_Pointer(cValue),
	}
}

func integerParamTag(tag C.FT_ULong, value *int) Parameter {
	if value == nil {
		return C.FT_Parameter{
			tag:  tag,
			data: nil,
		}
	}

	var cInt C.FT_Int32
	cValue := (*C.FT_Int32)(C.malloc(C.size_t(unsafe.Sizeof(cInt))))
	*cValue = C.FT_Int32(*value)

	return C.FT_Parameter{
		tag:  tag,
		data: C.FT_Pointer(cValue),
	}
}

// FT_Attach_File
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_attach_file

// FT_Attach_Stream
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_attach_stream
