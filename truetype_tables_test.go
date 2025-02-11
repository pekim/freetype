package freetype

import (
	"testing"

	"github.com/pekim/freetype-go/internal/font"
	"github.com/stretchr/testify/assert"
)

func TestFaceSfntTable(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	headPointer, err := face.GetSFNTTable(SFNT_HEAD)
	assert.Nil(t, err)
	head := (*TT_Header)(headPointer)
	assert.Equal(t, Long(0x5F0F3CF5), head.MagicNumber)

	os2Pointer, err := face.GetSFNTTable(SFNT_OS2)
	assert.Nil(t, err)
	os2 := (*TT_OS2)(os2Pointer)
	assert.Equal(t, Short(1233), os2.XAvgCharWidth)
}
