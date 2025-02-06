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
