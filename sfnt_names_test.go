package freetype

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pekim/freetype/internal/font"
)

func TestFaceGetSfntNameCount(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	assert.Equal(t, 22, face.GetSfntNameCount())
}

func TestFaceGetSFNTName(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	tableName, err := face.GetSfntName(NAME_ID_FONT_SUBFAMILY)
	assert.Nil(t, err)
	assert.Equal(t, "Book", tableName.String())
}
