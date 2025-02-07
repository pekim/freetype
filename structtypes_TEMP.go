package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

import (
	"unsafe"
)

func init() {
	assertSameSize(ListNodeRec{}, C.FT_ListNodeRec{})
	assertSameSize(ListRec{}, C.FT_ListRec{})
	assertSameSize(Outline{}, C.FT_Outline{})
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
