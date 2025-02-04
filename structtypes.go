package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"fmt"
	"unsafe"
)

func init() {
	assertSameSize(BitmapSize{}, C.FT_Bitmap_Size{})
	assertSameSize(CharMapRec{}, C.FT_CharMapRec{})
	assertSameSize(FaceRec{}, C.FT_FaceRec{})
	assertSameSize(Generic{}, C.FT_Generic{})
	assertSameSize(GlyphMetrics{}, C.FT_Glyph_Metrics{})
	assertSameSize(GlyphSlotRec{}, C.FT_GlyphSlotRec{})
	assertSameSize(ListNodeRec{}, C.FT_ListNodeRec{})
	assertSameSize(ListRec{}, C.FT_ListRec{})
	assertSameSize(Matrix{}, C.FT_Matrix{})
	assertSameSize(Outline{}, C.FT_Outline{})
	assertSameSize(SizeMetrics{}, C.FT_Size_Metrics{})
	assertSameSize(SizeRec{}, C.FT_SizeRec{})
	assertSameSize(SizeRequestRec{}, C.FT_Size_RequestRec{})
	assertSameSize(Vector{}, C.FT_Vector{})
}

func assertSameSize[A any, B any](a A, b B) {
	if unsafe.Sizeof(a) != unsafe.Sizeof(b) {
		panic(fmt.Sprintf("size of %T (%d) != size of %T (%d)", a, unsafe.Sizeof(a), b, unsafe.Sizeof(b)))
	}
}

type BitmapSize struct {
	height Short
	width  Short

	size Pos

	x_ppem Pos
	y_ppem Pos
}

type CharMapRec struct {
	Face       Face
	Encoding   Encoding
	PlatformID UShort
	EncodingID UShort
}

type FaceRec struct {
	num_faces  Long
	face_index Long

	face_flags  Long
	style_flags Long

	num_glyphs Long

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
	bbox BBox

	units_per_EM UShort
	ascender     Short
	descender    Short
	height       Short

	max_advance_width  Short
	max_advance_height Short

	underline_position  Short
	underline_thickness Short

	glyph   *GlyphSlotRec
	size    *SizeRec
	charmap *CharMapRec

	/* private fields, internal to FreeType */

	driver unsafe.Pointer
	memory unsafe.Pointer
	stream unsafe.Pointer

	sizes_list ListRec

	autohint   Generic        /* face-specific auto-hinter data */
	extensions unsafe.Pointer /* unused                         */

	internal unsafe.Pointer
}

type Generic struct {
	data      unsafe.Pointer
	finalizer unsafe.Pointer
}

type GlyphMetrics struct {
	width  Pos
	height Pos

	horiBearingX Pos
	horiBearingY Pos
	horiAdvance  Pos

	vertBearingX Pos
	vertBearingY Pos
	vertAdvance  Pos
}

type GlyphSlotRec struct {
	library Library
	face    Face
	next    *GlyphSlotRec
	glyph_index/* new in 2.10; was reserved previously */ UInt
	generic Generic

	metrics           GlyphMetrics
	linearHoriAdvance Fixed
	linearVertAdvance Fixed
	advance           Vector

	format Glyph_Format

	bitmap      Bitmap
	bitmap_left Int
	bitmap_top  Int

	outline Outline

	num_subglyphs UInt
	subglyphs     unsafe.Pointer

	control_data unsafe.Pointer
	control_len  C.long

	lsb_delta Pos
	rsb_delta Pos

	other    unsafe.Pointer
	internal unsafe.Pointer
}

type ListNodeRec struct {
	prev *ListNodeRec
	next *ListNodeRec
	data unsafe.Pointer
}

type ListRec struct {
	head *ListNodeRec
	tail *ListNodeRec
}

type Matrix struct {
	xx, xy Fixed
	yx, yy Fixed
}

type Outline struct {
	n_contours C.ushort /* number of contours in glyph        */
	n_points   C.ushort /* number of points in the glyph      */

	points   *Vector  /* the outline's points               */
	tags     *C.uchar /* the points flags                   */
	contours C.ushort /* the contour end points             */

	flags C.int /* outline masks                      */
	_     [4]byte
}

type SizeMetrics struct {
	x_ppem UShort /* horizontal pixels per EM               */
	y_ppem UShort /* vertical pixels per EM                 */

	x_scale Fixed /* scaling values used to convert font    */
	y_scale Fixed /* units to 26.6 fractional pixels        */

	ascender    Pos /* ascender in 26.6 frac. pixels          */
	descender   Pos /* descender in 26.6 frac. pixels         */
	height      Pos /* text height in 26.6 frac. pixels       */
	max_advance Pos /* max horizontal advance, in 26.6 pixels */
}

type SizeRec struct {
	face     Face        /* parent face object              */
	generic  Generic     /* generic pointer for client uses */
	metrics  SizeMetrics /* size metrics                    */
	internal unsafe.Pointer
}

type SizeRequestRec struct {
	Type           SizeRequestType
	Width          Long
	Height         Long
	HoriResolution UInt
	VertResolution UInt
}

type Vector struct {
	x Pos
	y Pos
}
