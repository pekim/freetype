package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"unsafe"
)

func init() {
	// assertSameSize(Bitmap{}, C.FT_Bitmap{})
	assertSameSize(BitmapSize{}, C.FT_Bitmap_Size{})
	assertSameSize(CharMapRec{}, C.FT_CharMapRec{})
	// assertSameSize(FaceRec{}, C.FT_FaceRec{})
	// assertSameSize(Generic{}, C.FT_Generic{})
	assertSameSize(GlyphMetrics{}, C.FT_Glyph_Metrics{})
	assertSameSize(GlyphSlotRec{}, C.FT_GlyphSlotRec{})
	assertSameSize(ListNodeRec{}, C.FT_ListNodeRec{})
	assertSameSize(ListRec{}, C.FT_ListRec{})
	// assertSameSize(Matrix{}, C.FT_Matrix{})
	assertSameSize(Outline{}, C.FT_Outline{})
	assertSameSize(SizeMetrics{}, C.FT_Size_Metrics{})
	assertSameSize(SizeRec{}, C.FT_SizeRec{})
	assertSameSize(SizeRequestRec{}, C.FT_Size_RequestRec{})
	// assertSameSize(UnitVector{}, C.FT_UnitVector{})
	// assertSameSize(Vector{}, C.FT_Vector{})
}

type BitmapSize struct {
	Height Short
	Width  Short

	Size Pos

	Xppem Pos
	YPpem Pos
}

type CharMapRec struct {
	Face       Face
	Encoding   Encoding
	PlatformID UShort
	EncodingID UShort
}

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

type GlyphSlotRec struct {
	Library Library
	Face    Face
	Next    *GlyphSlotRec
	GlyphIndex/* new in 2.10; was reserved previously */ UInt
	generic Generic

	Metrics           GlyphMetrics
	LinearHoriAdvance Fixed
	LinearVertAdvance Fixed
	Advance           Vector

	Format GlyphFormat

	Bitmap     Bitmap
	BitmapLeft Int
	BitmapTop  Int

	Outline Outline

	NumSubglyphs UInt
	subglyphs    unsafe.Pointer

	control_data unsafe.Pointer
	control_len  C.long

	LsbDelta Pos
	RsbDelta Pos

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
	Xppem UShort /* horizontal pixels per EM               */
	Yppem UShort /* vertical pixels per EM                 */

	XScale Fixed /* scaling values used to convert font    */
	YScale Fixed /* units to 26.6 fractional pixels        */

	Ascender   Pos /* ascender in 26.6 frac. pixels          */
	Descender  Pos /* descender in 26.6 frac. pixels         */
	Height     Pos /* text height in 26.6 frac. pixels       */
	MaxAdvance Pos /* max horizontal advance, in 26.6 pixels */
}

type SizeRec struct {
	Face     Face        /* parent face object              */
	generic  Generic     /* generic pointer for client uses */
	Metrics  SizeMetrics /* size metrics                    */
	internal unsafe.Pointer
}

type SizeRequestRec struct {
	Type           SizeRequestType
	Width          Long
	Height         Long
	HoriResolution UInt
	VertResolution UInt
}
