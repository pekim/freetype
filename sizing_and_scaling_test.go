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

func TestFaceGetTransformSetTransform(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	matrix := Matrix{XX: 1, XY: 2, YX: 3, YY: 4}
	vector := Vector{X: 5, Y: 6}
	face.SetTransform(&matrix, &vector)

	matrix2, vector2 := face.GetTransform()
	assert.Equal(t, matrix, matrix2)
	assert.Equal(t, vector, vector2)
}
