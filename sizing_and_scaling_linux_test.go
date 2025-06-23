package freetype

import (
	_ "embed"
	"testing"

	"github.com/pekim/freetype/internal/font"
	"github.com/stretchr/testify/assert"
)

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
