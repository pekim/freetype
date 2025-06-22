//go:build linux

package freetype

import "modernc.org/libfreetype"

/*
SetTransform sets the transformation that is applied to glyph images when they are loaded into a
glyph slot through FT_Load_Glyph.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_set_transform
*/
func (face Face) SetTransform(matrix *Matrix, delta *Vector) {
	libfreetype.XFT_Set_Transform(face.tls, face.face, toUintptr(matrix), toUintptr(delta))
}

/*
GetTransform returns the transformation that is applied to glyph images when they are loaded into
a glyph slot through Load_Glyph. See SetTransform for more details.

https://freetype.org/freetype2/docs/reference/ft2-sizing_and_scaling.html#ft_get_transform
*/
func (face Face) GetTransform() (Matrix, Vector) {
	var matrix Matrix
	var vector Vector
	libfreetype.XFT_Get_Transform(face.tls, face.face, toUintptr(&matrix), toUintptr(&vector))
	return matrix, vector
}
