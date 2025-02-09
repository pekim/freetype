package freetype

import (
	_ "embed"
	"testing"

	"github.com/pekim/freetype-go/internal/font"
	"github.com/stretchr/testify/assert"
)

func TestFaceRecFields(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)
	rec := face.Rec()

	assert.Equal(t, Long(3377), rec.NumGlyphs)
	assert.Equal(t, "DejaVu Sans Mono", rec.FamilyName())
	assert.Equal(t, "Book", rec.StyleName())
	assert.Equal(t, 5, len(rec.Charmaps()))
	assert.Equal(t, ENCODING_UNICODE, rec.Charmaps()[0].Encoding)
}

func TestLibraryNewFace(t *testing.T) {
	lib, _ := Init()

	// good font file
	face, err := lib.NewFace("internal/font/DejaVuSansMono.ttf", 0)
	assert.Nil(t, err)
	assert.NotNil(t, face.face)
	err = face.Done()
	assert.Nil(t, err)

	// no such file
	face, err = lib.NewFace("bad path", 0)
	assert.Error(t, err)

	// file exists but is not a font file
	face, err = lib.NewFace("library.go", 0)
	assert.Error(t, err)
}

func TestFaceDone(t *testing.T) {
	lib, _ := Init()
	err := lib.Done()
	assert.Nil(t, err)
}

func TestFaceReference(t *testing.T) {
	lib, _ := Init()
	face, err := lib.NewMemoryFace(font.DejaVuSansMono, 0)
	assert.Nil(t, err)
	err = face.Reference()
	assert.Nil(t, err)
}

func TestLibraryNewMemoryFace(t *testing.T) {
	lib, _ := Init()

	// good font data
	face, err := lib.NewMemoryFace(font.DejaVuSansMono, 0)
	assert.Nil(t, err)
	assert.NotNil(t, face.face)

	// bad font data
	face, err = lib.NewMemoryFace(font.DejaVuSansMono[1:], 0)
	assert.Error(t, err)
}

func TestFaceProperies(t *testing.T) {
	lib, _ := Init()
	true_ := true
	number := 1
	weights := [5]byte{1, 2, 3, 4, 5}

	// supported properties - set
	{
		face, err := lib.NewMemoryFace(font.DejaVuSansMono, 0)
		assert.Nil(t, err)
		err = face.Properties(
			ParameterTagStemDarkening(&true_),
			ParameterTagLCDFilterWeights(&weights),
			ParameterTagRandomSeed(&number),
		)
		assert.Nil(t, err)
	}

	// supported property - reset
	{
		face, err := lib.NewMemoryFace(font.DejaVuSansMono, 0)
		assert.Nil(t, err)
		err = face.Properties(
			ParameterTagStemDarkening(nil),
		)
		assert.Nil(t, err)
	}

	// unsupported property
	{
		face, err := lib.NewMemoryFace(font.DejaVuSansMono, 0)
		assert.Nil(t, err)
		err = face.Properties(
			ParameterTagIgnoreTypoGraphicFamily(&true_),
		)
		assert.Error(t, err)
	}
}
