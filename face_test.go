package freetype

import (
	"testing"

	"github.com/pekim/freetype-go/internal/font"
	"github.com/stretchr/testify/assert"
)

func TestFaceReference(t *testing.T) {
	lib, _ := Init()
	defer func() { _ = lib.Done() }()

	face, err := lib.NewMemoryFace(font.DejaVuSansMono, 0)
	assert.Nil(t, err)
	err = face.Reference()
	assert.Nil(t, err)
}

func TestFaceProperies(t *testing.T) {
	lib, _ := Init()
	defer func() { _ = lib.Done() }()

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
