package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

func cBoolToGo(value C.FT_Bool) bool {
	return value != 0
}
