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
	weights := [5]byte{1, 2, 3, 4, 5}

	// supported properties - set
	{
		face, err := lib.NewMemoryFace(font.DejaVuSansMono, 0)
		assert.Nil(t, err)
		err = face.Properties(
			ParamTagStemDarkening(&true_),
			ParamTagLCDFilterWeights(&weights),
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

	matrix := Matrix{XX: 1, XY: 2, YX: 3, YY: 4}
	vector := Vector{X: 5, Y: 6}
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

	charmaps := face.Rec().Charmaps()
	index := GetCharmapIndex(charmaps[2])
	assert.Equal(t, Int(2), index)
}

func TestFaceGetIndex(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	assert.Equal(t, UInt(0x3), face.GetCharIndex(' '))
	assert.Equal(t, UInt(0x44), face.GetCharIndex('a'))
}

func TestFaceGetFirstCharGetNextChar(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	// first char
	charCode, glyphIndex := face.GetFirstChar()
	assert.Equal(t, ULong(' '), charCode)
	assert.Equal(t, UInt(3), glyphIndex)

	// second char
	charCode, glyphIndex = face.GetNextChar(charCode)
	assert.Equal(t, ULong('!'), charCode)
	assert.Equal(t, UInt(4), glyphIndex)

	// remaining chars
	c := 2
	for glyphIndex != 0 {
		charCode, glyphIndex = face.GetNextChar(charCode)
		c++
	}
	c-- // discount the last one, as it wasn't a char
	assert.Equal(t, 3322, c)
}

func TestFaceLoadChar(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	// face.SetPixelSizes(0, 20)
	err := face.LoadChar('A', LOAD_DEFAULT)
	assert.Nil(t, err)
	err = face.RenderGlyph(RENDER_MODE_NORMAL)
	assert.Nil(t, err)
	assertGlyphRecFieldsForUppercaseA(t, face.Rec().Glyph)
}

func TestFaceLoadGlyph(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	err := face.LoadGlyph(face.GetCharIndex('A'), LOAD_DEFAULT)
	assert.Nil(t, err)
	err = face.RenderGlyph(RENDER_MODE_NORMAL)
	assert.Nil(t, err)
	assertGlyphRecFieldsForUppercaseA(t, face.Rec().Glyph)
}

func assertGlyphRecFieldsForUppercaseA(t *testing.T, rec *GlyphSlotRec) {
	t.Helper()
	assert.Equal(t, GLYPH_FORMAT_BITMAP, rec.Format)
	assert.Equal(t, Pos(1159), rec.Metrics.Width)
	assert.Equal(t, expectedBitmapForA, rec.Bitmap.Buffer())
	assert.Equal(t, bitmapVisualizationA, rec.Bitmap.BufferVisualization())
}

var expectedBitmapForA = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1B, 0x54, 0x54, 0x54, 0x32, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x89, 0xFF, 0xFF, 0xFF, 0xCC, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xD8, 0xFF, 0xFF, 0xFF, 0xFF, 0x1B, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x25, 0xFF, 0xFF, 0xEA, 0xFF, 0xFF, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x74, 0xFF, 0xFF, 0x6B, 0xFE, 0xFF, 0xB7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xC3, 0xFF, 0xFC, 0x11, 0xCE, 0xFF, 0xF9, 0x0D, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x13, 0xFC, 0xFF, 0xC6, 0x00, 0x87, 0xFF, 0xFF, 0x54, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x5F, 0xFF, 0xFF, 0x7E, 0x00, 0x3F, 0xFF, 0xFF, 0xA2, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0xAE, 0xFF, 0xFF, 0x37, 0x00, 0x05, 0xF2, 0xFF, 0xED, 0x03, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x07, 0xF4, 0xFF, 0xED, 0x02, 0x00, 0x00, 0xB1, 0xFF, 0xFF, 0x3F, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x4A, 0xFF, 0xFF, 0xA8, 0x00, 0x00, 0x00, 0x69, 0xFF, 0xFF, 0x8D, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x99, 0xFF, 0xFF, 0x61, 0x00, 0x00, 0x00, 0x22, 0xFF, 0xFF, 0xDC, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x01, 0xE5, 0xFF, 0xFF, 0x1A, 0x00, 0x00, 0x00, 0x00, 0xDA, 0xFF, 0xFF, 0x2A, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x36, 0xFF, 0xFF, 0xD2, 0x00, 0x00, 0x00, 0x00, 0x00, 0x93, 0xFF, 0xFF, 0x79, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x84, 0xFF, 0xFF, 0x8B, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4C, 0xFF, 0xFF, 0xC7, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0xD2, 0xFF, 0xFF, 0xBE, 0x9C, 0x9C, 0x9C, 0x9C, 0x9C, 0xA6, 0xFF, 0xFF, 0xFD, 0x17, 0x00, 0x00,
	0x00, 0x00, 0x20, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x64, 0x00, 0x00,
	0x00, 0x00, 0x6F, 0xFF, 0xFF, 0xF6, 0xEC, 0xEC, 0xEC, 0xEC, 0xEC, 0xEC, 0xEC, 0xF1, 0xFF, 0xFF, 0xB2, 0x00, 0x00,
	0x00, 0x00, 0xBD, 0xFF, 0xFF, 0x5E, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x22, 0xFF, 0xFF, 0xF6, 0x0A, 0x00,
	0x00, 0x10, 0xFB, 0xFF, 0xFF, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xDA, 0xFF, 0xFF, 0x4F, 0x00,
	0x00, 0x5A, 0xFF, 0xFF, 0xD0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x92, 0xFF, 0xFF, 0x9D, 0x00,
	0x00, 0xA8, 0xFF, 0xFF, 0x89, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x49, 0xFF, 0xFF, 0xE9, 0x02,
	0x05, 0xF1, 0xFF, 0xFF, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09, 0xF7, 0xFF, 0xFF, 0x3A,
	0x45, 0xFF, 0xFF, 0xF4, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xB8, 0xFF, 0xFF, 0x88,
}

//go:embed bitmap_visualization_A
var bitmapVisualizationA string

func TestFaceGetKerning(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSans, 0)

	leftGlyph := face.GetCharIndex('V')
	rightGlyph := face.GetCharIndex('A')
	kerning, err := face.GetKerning(leftGlyph, rightGlyph, KERNING_UNSCALED)
	assert.Nil(t, err)
	assert.Equal(t, Pos(-131), kerning.X)
	assert.Equal(t, Pos(0), kerning.Y)
}
