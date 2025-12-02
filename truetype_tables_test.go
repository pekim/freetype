package freetype

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pekim/freetype/internal/font"
)

func TestFaceSfntTable(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	headPointer, err := face.GetSfntTable(SFNT_HEAD)
	assert.Nil(t, err)
	head := (*TT_Header)(headPointer)
	assert.Equal(t, Long(0x5F0F3CF5), head.MagicNumber)

	os2Pointer, err := face.GetSfntTable(SFNT_OS2)
	assert.Nil(t, err)
	os2 := (*TT_OS2)(os2Pointer)
	assert.Equal(t, Short(1233), os2.XAvgCharWidth)
}

func TestFaceLoadSfntTable(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)
	tag := imageTag('O', 'S', '/', '2')

	// Get the table's length.
	var length ULong
	err := face.LoadSfntTable(tag, 0, nil, &length)
	assert.Nil(t, err)

	// Get the table.
	buffer := make([]byte, length)
	err = face.LoadSfntTable(tag, 0, buffer, &length)
	assert.Nil(t, err)

	// LoadSFNTTable gets the raw font table, which is not the same as the TT_OS2 type.
	// All of the data is big-endian.
	//
	// Verify the xAvgCharWidth field is the same value as asserted in the TestFaceSfntTable test.
	var xAvgCharWidth int16
	_, err = binary.Decode(buffer[2:], binary.BigEndian, &xAvgCharWidth)
	assert.Nil(t, err)
	assert.Equal(t, int16(1233), xAvgCharWidth)
}

func TestFaceSfntTableInfo(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.DejaVuSansMono, 0)

	// Get number of tables.
	count, err := face.SfntTableInfo(0, nil)
	assert.Nil(t, err)
	assert.Equal(t, ULong(18), count)

	// Get info for a specific table.
	var tag ULong
	length, err := face.SfntTableInfo(4, &tag)
	assert.Nil(t, err)
	assert.Equal(t, ULong(86), length)
	assert.Equal(t, ULong(imageTag('O', 'S', '/', '2')), tag)
}
