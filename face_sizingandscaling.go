package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

func (face Face) SetCharSize(
	charWidth F26Dot6, charHeight F26Dot6,
	horzResolution UInt, vertResolution UInt,
) error {
	err := C.FT_Set_Char_Size(face.face, charWidth, charHeight, horzResolution, vertResolution)
	return newError(err, "failed to set face char size")
}
