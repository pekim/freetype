package freetype

import (
	"testing"

	"github.com/pekim/freetype-go/internal/font"
	"github.com/stretchr/testify/assert"
)

func TestFaceSelectCharmap(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	err := face.SelectCharmap(ENCODING_UNICODE)
	assert.Nil(t, err)

	err = face.SelectCharmap(ENCODING_BIG5)
	assert.Error(t, err)
}

func TestFaceSetCharmap(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	err := face.SetCharmap(CharMapRec{
		Face:       face.face,
		Encoding:   ENCODING_UNICODE,
		PlatformID: 0,
		EncodingID: 1,
	})
	assert.Error(t, err)
}

func TestFaceGetCharmapIndex(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	charmaps := face.Rec().Charmaps()
	index := face.GetCharmapIndex(charmaps[2])
	assert.Equal(t, Int(2), index)
}

func TestFaceGetCharIndex(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	assert.Equal(t, UInt(0x3), face.GetCharIndex(' '))
	assert.Equal(t, UInt(0x44), face.GetCharIndex('a'))
}

// func TestFaceGetFirstCharGetNextChar(t *testing.T) {
// 	lib, _ := Init()
// 	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

// 	// first char
// 	charCode, glyphIndex := face.GetFirstChar()
// 	assert.Equal(t, ULong(' '), charCode)
// 	assert.Equal(t, UInt(3), glyphIndex)

// 	// second char
// 	charCode, glyphIndex = face.GetNextChar(charCode)
// 	assert.Equal(t, ULong('!'), charCode)
// 	assert.Equal(t, UInt(4), glyphIndex)

// 	// remaining chars
// 	c := 2
// 	for glyphIndex != 0 {
// 		charCode, glyphIndex = face.GetNextChar(charCode)
// 		c++
// 	}
// 	c-- // discount the last one, as it wasn't a char
// 	assert.Equal(t, 3322, c)
// }

// func TestFaceLoadChar(t *testing.T) {
// 	lib, _ := Init()
// 	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

// 	// face.SetPixelSizes(0, 20)
// 	err := face.LoadChar('A', LOAD_DEFAULT)
// 	assert.Nil(t, err)
// 	err = face.RenderGlyph(RENDER_MODE_NORMAL)
// 	assert.Nil(t, err)
// 	assertGlyphRecFieldsForUppercaseA(t, face.Rec().Glyph)
// }
