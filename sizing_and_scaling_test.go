package freetype

import (
	_ "embed"
	"testing"

	"github.com/pekim/freetype/internal/font"
	"github.com/stretchr/testify/assert"
)

func TestFaceSetCharSize(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)
	err := face.SetCharSize(50, 50, 96, 96)
	assert.Nil(t, err)
}

func TestFaceSetPixelSizes(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)
	err := face.SetPixelSizes(24, 0)
	assert.Nil(t, err)
}

func TestFaceRequestSize(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)
	err := face.RequestSize(SizeRequestRec{
		Type:           SIZE_REQUEST_TYPE_BBOX,
		Width:          50,
		Height:         50,
		HoriResolution: 96,
		VertResolution: 96,
	})
	assert.Nil(t, err)
}

func TestFaceSelectSize(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)
	err := face.SelectSize(1)
	assert.Error(t, err) // the font is not a bitmap font
}
