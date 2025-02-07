package freetype

import (
	"testing"

	"github.com/pekim/freetype-go/internal/font"
	"github.com/stretchr/testify/assert"
)

func TestFaceLoadGlyph(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	err := face.LoadGlyph(face.GetCharIndex('A'), LOAD_DEFAULT)
	assert.Nil(t, err)
	err = face.RenderGlyph(RENDER_MODE_NORMAL)
	assert.Nil(t, err)
	assertGlyphRecFieldsForUppercaseA(t, face.Rec().Glyph)
}

func TestFaceGetKerning(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSans, 0)

	leftGlyph := face.GetCharIndex('V')
	rightGlyph := face.GetCharIndex('A')
	kerning, err := face.GetKerning(leftGlyph, rightGlyph, KERNING_UNSCALED)
	assert.Nil(t, err)
	assert.Equal(t, Pos(-131), kerning.X)
	assert.Equal(t, Pos(0), kerning.Y)
}
