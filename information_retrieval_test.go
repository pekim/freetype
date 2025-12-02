package freetype

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pekim/freetype/internal/font"
)

func TestFaceGetNameIndex(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSans, 0)

	assert.True(t, face.HasGlyphNames())
	index, err := face.GetNameIndex("asciitilde")
	assert.Nil(t, err)
	assert.Equal(t, face.GetCharIndex('~'), index)
}

func TestFaceGetGlyphName(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSans, 0)

	assert.True(t, face.HasGlyphNames())
	name, err := face.GetGlyphName(face.GetCharIndex('~'))
	assert.Nil(t, err)
	assert.Equal(t, "asciitilde", name)
}

func TestFaceGetPostscriptName(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSans, 0)

	assert.Equal(t, "DejaVuSans", face.GetPostscriptName())
}

func TestFaceGetFSTypeFlags(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSans, 0)

	assert.Equal(t, FSType(0), face.GetFSTypeFlags())
}
