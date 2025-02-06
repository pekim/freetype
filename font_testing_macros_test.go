package freetype

import (
	_ "embed"
	"testing"

	"github.com/pekim/freetype-go/internal/font"
	"github.com/stretchr/testify/assert"
)

func TestFaceFontTestingMacros(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	assert.True(t, face.HasHorizontal())
	assert.False(t, face.HasVertical())
	assert.False(t, face.HasKerning())
	assert.False(t, face.HasFixedSizes())
	assert.True(t, face.HasGlyphNames())
	assert.False(t, face.HasColor())
	assert.False(t, face.HasMultipleMasters())
	assert.False(t, face.HasSVG())
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
