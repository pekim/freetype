package freetype

import (
	"testing"

	"github.com/pekim/freetype-go/internal/font"
	"github.com/stretchr/testify/assert"
)

func TestFaceReference(t *testing.T) {
	lib, _ := Init()
	face, err := lib.NewMemoryFace(font.DejaVuSansMono, 0)
	assert.Nil(t, err)
	err = face.Reference()
	assert.Nil(t, err)
}

func TestFaceProperies(t *testing.T) {
	lib, _ := Init()
	true_ := true
	number := 1
	// weights := [5]byte{1, 2, 3, 4, 5}

	// supported properties - set
	{
		face, err := lib.NewMemoryFace(font.DejaVuSansMono, 0)
		assert.Nil(t, err)
		err = face.Properties(
			ParamTagStemDarkening(&true_),
			// ParamTagLCDFilterWeights(&weights),
			ParamTagRandomSeed(&number),
		)
		assert.Nil(t, err)
	}

	// supported property - reset
	{
		face, err := lib.NewMemoryFace(font.DejaVuSansMono, 0)
		assert.Nil(t, err)
		err = face.Properties(
			ParamTagStemDarkening(nil),
		)
		assert.Nil(t, err)
	}

	// unsupported property
	{
		face, err := lib.NewMemoryFace(font.DejaVuSansMono, 0)
		assert.Nil(t, err)
		err = face.Properties(
			ParamTagIgnoreTypoGraphicFamily(&true_),
		)
		assert.Error(t, err)
	}
}

func TestFaceFontTestingMacros(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	assert.True(t, face.HasHorizontal())
	assert.False(t, face.HasVertical())
	assert.False(t, face.HasKerning())
	assert.False(t, face.HadFixedSizes())
	assert.True(t, face.HasGlyphNames())
	assert.False(t, face.HasColor())
	assert.False(t, face.HasMultipleMasters())
	assert.False(t, face.HaseSVG())
	assert.False(t, face.HasSbix())
	assert.False(t, face.HasSbixOverlay())
	assert.True(t, face.IsSFNT())
	assert.True(t, face.IsScalable())
	assert.True(t, face.IsFixedWidth())
	assert.False(t, face.IsCIDKeyed())
	assert.False(t, face.IsTricky())
	assert.False(t, face.IsNamedInstance())
	assert.False(t, face.IsVariation())
}

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

func TestFaceSelectSize(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)
	err := face.SelectSize(1)
	assert.Error(t, err) // the font is not a bitmap font
}
