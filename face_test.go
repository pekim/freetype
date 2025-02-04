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

func TestFaceGetTransformSetTransform(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	matrix := Matrix{xx: 1, xy: 2, yx: 3, yy: 4}
	vector := Vector{x: 5, y: 6}
	face.SetTransform(&matrix, &vector)

	matrix2, vector2 := face.GetTransform()
	assert.Equal(t, matrix, matrix2)
	assert.Equal(t, vector, vector2)
}

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
		Face:       face,
		Encoding:   ENCODING_UNICODE,
		PlatformID: 0,
		EncodingID: 1,
	})
	assert.Error(t, err)
}

func TestFaceGetCharmapIndex(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	// No idea whether 5 really is the expected index.
	index := GetCharmapIndex(CharMapRec{
		Face:       face,
		Encoding:   ENCODING_UNICODE,
		PlatformID: 0,
		EncodingID: 1,
	})
	assert.Equal(t, Int(5), index)
}

func TestFaceGetIndex(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	assert.Equal(t, UInt(0x3), face.GetCharIndex(' '))
	assert.Equal(t, UInt(0x44), face.GetCharIndex('a'))
}
