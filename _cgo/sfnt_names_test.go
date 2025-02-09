package freetype

import (
	"testing"

	"github.com/pekim/freetype-go/internal/font"
	"github.com/stretchr/testify/assert"
)

func TestFaceGetSfntNameCount(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	assert.Equal(t, 22, face.GetSFNTNameCount())
}

func TestFaceGetSFNTName(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	tableName, err := face.GetSFNTName(NAME_ID_FONT_SUBFAMILY)
	assert.Nil(t, err)
	assert.Equal(t, "Book", tableName.String())
}
