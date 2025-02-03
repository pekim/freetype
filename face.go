package freetype

// #include <ft2build.h>
// #include FT_FREETYPE_H
import "C"

type Face struct {
	face C.FT_Face
}
