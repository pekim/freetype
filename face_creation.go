package freetype

import (
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libfreetype"
)

// Functions to manage fonts.

/*
Face is a handle to a typographic face object.
A face object models a given typeface, in a given style.

https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_face
*/
type Face struct {
	face libfreetype.TFT_Face
	tls  *libc.TLS
}

// Rec returns a pointer to the FaceRec that is referenced by the Face.
func (face Face) Rec() *FaceRec {
	return fromUintptr[FaceRec](face.face)
}

func init() {
	assertSameSize(FaceRec{}, libfreetype.TFT_FaceRec{})
}

// FaceRec is a FreeType root face class structure.
// A face object models a typeface in a font file.
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_facerec
type FaceRec struct {
	NumFaces  Long
	FaceIndex Long

	FaceFlags  FACE_FLAG
	StyleFlags STYLE_FLAG

	NumGlyphs Long

	family_name uintptr
	style_name  uintptr

	num_fixed_sizes Int
	available_sizes *BitmapSize

	num_charmaps Int
	charmaps     *CharMap

	_ Generic // generic

	/* The following member variables (down to `underline_thickness`) */
	/* outlines are only relevant to scalable  cf. @FT_Bitmap_Size    */
	/* for bitmap fonts.                                              */

	Bbox BBox

	UnitsPerEM UShort
	Ascender   Short
	Descender  Short
	Height     Short

	MaxAdvanceWidth  Short
	MaxAdvanceHeight Short

	UnderlinePosition  Short
	UnderlineThickness Short

	Glyph   GlyphSlot
	Size    Size
	Charmap CharMap

	/* private fields, internal to FreeType */

	_ unsafe.Pointer // driver
	_ unsafe.Pointer // memory
	_ unsafe.Pointer // stream

	_ libfreetype.TFT_ListRec // sizes_list

	_ Generic        /* autohint - face-specific auto-hinter data */
	_ unsafe.Pointer /* extensions - unused                         */

	_ unsafe.Pointer // internal
}

/*
FamilyName returns the face's family name.

(This exposes the C string referenced by the unexported family_name field.)
*/
func (fr *FaceRec) FamilyName() string {
	return libc.GoString(fr.family_name)
}

/*
Stylename returns the face's style name.

(This exposes the C string referenced by the unexported style_name field.)
*/
func (fr *FaceRec) StyleName() string {
	return libc.GoString(fr.style_name)
}

/*
AvailableSizes returns a slice of Bitmap_Size for all bitmap strikes in the face.

(This exposes the data referenced by the unexported num_fixed_sizes and available_size fields.)
*/
func (fr *FaceRec) AvailableSizes() []BitmapSize {
	return unsafe.Slice(fr.available_sizes, fr.num_fixed_sizes)
}

/*
Charmaps returns the charmaps of the face.

(This exposes the data referenced by the unexported num_charmaps and charmap fields.)
*/
func (fr *FaceRec) Charmaps() []CharMap {
	return unsafe.Slice(fr.charmaps, fr.num_charmaps)
}

// A list of bit flags used in the face_flags field of the FaceRec structure.
// They inform client applications of properties of the corresponding face.
//
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_face_flag_xxx
type FACE_FLAG = Long

const (
	FACE_FLAG_SCALABLE         = FACE_FLAG(1 << 0)
	FACE_FLAG_FIXED_SIZES      = FACE_FLAG(1 << 1)
	FACE_FLAG_FIXED_WIDTH      = FACE_FLAG(1 << 2)
	FACE_FLAG_SFNT             = FACE_FLAG(1 << 3)
	FACE_FLAG_HORIZONTAL       = FACE_FLAG(1 << 4)
	FACE_FLAG_VERTICAL         = FACE_FLAG(1 << 5)
	FACE_FLAG_KERNING          = FACE_FLAG(1 << 6)
	FACE_FLAG_FAST_GLYPHS      = FACE_FLAG(1 << 7)
	FACE_FLAG_MULTIPLE_MASTERS = FACE_FLAG(1 << 8)
	FACE_FLAG_GLYPH_NAMES      = FACE_FLAG(1 << 9)
	FACE_FLAG_EXTERNAL_STREAM  = FACE_FLAG(1 << 10)
	FACE_FLAG_HINTER           = FACE_FLAG(1 << 11)
	FACE_FLAG_CID_KEYED        = FACE_FLAG(1 << 12)
	FACE_FLAG_TRICKY           = FACE_FLAG(1 << 13)
	FACE_FLAG_COLOR            = FACE_FLAG(1 << 14)
	FACE_FLAG_VARIATION        = FACE_FLAG(1 << 15)
	FACE_FLAG_SVG              = FACE_FLAG(1 << 16)
	FACE_FLAG_SBIX             = FACE_FLAG(1 << 17)
	FACE_FLAG_SBIX_OVERLAY     = FACE_FLAG(1 << 18)
)

// A list of bit flags to indicate the style of a given face.
// These are used in the style_flags field of FaceRec.
//
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_style_flag_xxx
type STYLE_FLAG = Long

const (
	STYLE_FLAG_ITALIC = STYLE_FLAG(1 << 0)
	STYLE_FLAG_BOLD   = STYLE_FLAG(1 << 1)
)

// NewFace opens a font by its pathname.
//
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_new_face
func (lib Library) NewFace(filepathname string, faceIndex int) (Face, error) {
	cFilepathname, err := libc.CString(filepathname)
	if err != nil {
		return Face{}, err
	}
	defer libc.Xfree(nil, cFilepathname)

	face, freeFace := alloc(lib.tls, Face{})
	face.tls = lib.tls

	err_ := libfreetype.XFT_New_Face(lib.tls, lib.library, cFilepathname, Long(faceIndex), toUintptr(&face.face))

	face_ := *face
	freeFace()
	return face_, newError(err_, "failed to create a face for file '%s'", filepathname)
}

// Done discards a given face object, as well as all of its child slots and sizes.
//
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_done_face
func (face Face) Done() error {
	err := libfreetype.XFT_Done_Face(face.tls, face.face)
	return newError(err, "failed to discard face")
}

// /*
// A counter gets initialized to 1 at the time a Face structure is created.
// This function increments the counter.

// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_reference_face
// */
// func (face Face) Reference() error {
// 	err := C.FT_Reference_Face(face.face)
// 	return newError(err, "failed to reference face")
// }

// NewMemoryFace opens a font that has been loaded into memory.
//
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_new_memory_face
func (lib Library) NewMemoryFace(data []byte, faceIndex int) (Face, error) {
	face, freeFace := alloc(lib.tls, Face{})
	face.tls = lib.tls

	err := libfreetype.XFT_New_Memory_Face(
		lib.tls, lib.library,
		toUintptr(&data[0]), libfreetype.TFT_Long(len(data)),
		libfreetype.TFT_Long(faceIndex), toUintptr(&face.face))

	face_ := *face
	freeFace()
	return face_, newError(err, "failed to create a new memory face")
}

// /*
// Properties sets or overrides certain (library or module-wide) properties on a face-by-face basis.

// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_face_properties
// */
// func (face Face) Properties(properties ...Parameter) error {
// 	err := C.FT_Face_Properties(face.face, C.FT_UInt(len(properties)), (*C.FT_Parameter)(&properties[0]))
// 	for _, param := range properties {
// 		param.freeData()
// 	}
// 	return newError(err, "failed to set face properties")
// }

// // FT_Open_Face
// // https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_open_face

// // FT_Open_Args
// // https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_open_args

// // FT_OPEN_XXX
// // https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_open_xxx

// /*
// Parameter is a simple structure to pass more or less generic parameters to Library.OpenFace and Face.Properties.

// Use one of the ParameterTagXXX functions to create a Parameter initialized to a specific parameter tag with data.
// Pass a nil argument to reset the property that the parameter represents.
// */
// type Parameter = C.FT_Parameter

// func (param Parameter) freeData() {
// 	C.free(unsafe.Pointer(param.data))
// }

// // ParameterTagIgnoreTypoGraphicFamily creates a Parameter for the FT_PARAM_TAG_IGNORE_TYPOGRAPHIC_FAMILY tag.
// //
// // https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_ignore_typographic_family
// func ParameterTagIgnoreTypoGraphicFamily(value *bool) Parameter {
// 	return booleanParamTag(C.FT_PARAM_TAG_IGNORE_TYPOGRAPHIC_FAMILY, value)
// }

// // ParameterTagIgnoreTypoGraphicSubfamily creates a Parameter for the FT_PARAM_TAG_IGNORE_TYPOGRAPHIC_SUBFAMILY tag.
// //
// // https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_ignore_typographic_subfamily
// func ParameterTagIgnoreTypoGraphicSubfamily(value *bool) Parameter {
// 	return booleanParamTag(C.FT_PARAM_TAG_IGNORE_TYPOGRAPHIC_SUBFAMILY, value)
// }

// // ParameterTagIncremental creates a Parameter for the FT_PARAM_TAG_INCREMENTAL tag.
// //
// // https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_incremental
// func ParameterTagIncremental(value *bool) Parameter {
// 	return booleanParamTag(C.FT_PARAM_TAG_INCREMENTAL, value)
// }

// // ParameterTagIgnoreSbix creates a Parameter for the FT_PARAM_TAG_IGNORE_SBIX tag.
// //
// // https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_ignore_sbix
// func ParameterTagIgnoreSbix(value *bool) Parameter {
// 	return booleanParamTag(C.FT_PARAM_TAG_IGNORE_SBIX, value)
// }

// const LCDFilterWeightsLen = 5

// // ParameterTagLCDFilterWeights creates a Parameter for the FT_PARAM_TAG_LCD_FILTER_WEIGHTS tag.
// //
// // https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_lcd_filter_weights
// func ParameterTagLCDFilterWeights(weights *[LCDFilterWeightsLen]byte) Parameter {
// 	if weights == nil {
// 		return C.FT_Parameter{
// 			tag:  C.FT_PARAM_TAG_LCD_FILTER_WEIGHTS,
// 			data: nil,
// 		}
// 	}

// 	cWeights := (*C.uchar)(C.malloc(LCDFilterWeightsLen))
// 	C.memcpy(unsafe.Pointer(cWeights), unsafe.Pointer(&(*weights)[0]), LCDFilterWeightsLen)

// 	return C.FT_Parameter{
// 		tag:  C.FT_PARAM_TAG_LCD_FILTER_WEIGHTS,
// 		data: C.FT_Pointer(cWeights),
// 	}
// }

// // ParameterTagRandomSeed creates a Parameter for the FT_PARAM_TAG_RANDOM_SEED tag.
// //
// // https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_random_seed
// func ParameterTagRandomSeed(value *int) Parameter {
// 	return integerParamTag(C.FT_PARAM_TAG_RANDOM_SEED, value)
// }

// // ParameterTagStemDarkening creates a Parameter for the FT_PARAM_TAG_STEM_DARKENING tag.
// //
// // https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_stem_darkening
// func ParameterTagStemDarkening(value *bool) Parameter {
// 	return booleanParamTag(C.FT_PARAM_TAG_STEM_DARKENING, value)
// }

// // ParameterTagUnpatentedHinting creates a Parameter for the FT_PARAM_TAG_UNPATENTED_HINTING tag.
// //
// // https://freetype.org/freetype2/docs/reference/ft2-parameter_tags.html#ft_param_tag_unpatented_hinting
// func ParameterTagUnpatentedHinting(value *bool) Parameter {
// 	return booleanParamTag(C.FT_PARAM_TAG_UNPATENTED_HINTING, value)
// }

// func booleanParamTag(tag C.FT_ULong, value *bool) Parameter {
// 	if value == nil {
// 		return C.FT_Parameter{
// 			tag:  tag,
// 			data: nil,
// 		}
// 	}

// 	var cBool C.FT_Bool
// 	cValue := (*C.FT_Bool)(C.malloc(C.size_t(unsafe.Sizeof(cBool))))
// 	if *value {
// 		*cValue = 1
// 	} else {
// 		*cValue = 0
// 	}

// 	return C.FT_Parameter{
// 		tag:  tag,
// 		data: C.FT_Pointer(cValue),
// 	}
// }

// func integerParamTag(tag C.FT_ULong, value *int) Parameter {
// 	if value == nil {
// 		return C.FT_Parameter{
// 			tag:  tag,
// 			data: nil,
// 		}
// 	}

// 	var cInt C.FT_Int32
// 	cValue := (*C.FT_Int32)(C.malloc(C.size_t(unsafe.Sizeof(cInt))))
// 	*cValue = C.FT_Int32(*value)

// 	return C.FT_Parameter{
// 		tag:  tag,
// 		data: C.FT_Pointer(cValue),
// 	}
// }

// FT_Attach_File
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_attach_file

// FT_Attach_Stream
// https://freetype.org/freetype2/docs/reference/ft2-face_creation.html#ft_attach_stream
