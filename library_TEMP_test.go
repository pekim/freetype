package freetype

import (
	"testing"

	"github.com/pekim/freetype-go/internal/font"
	"github.com/stretchr/testify/assert"
)

func TestLibraryNewFace(t *testing.T) {
	lib, _ := Init()
	defer func() { _ = lib.Done() }()

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

func TestLibraryNewMemoryFace(t *testing.T) {
	lib, _ := Init()
	defer func() { _ = lib.Done() }()

	// good font data
	face, err := lib.NewMemoryFace(font.DejaVuSansMono, 0)
	assert.Nil(t, err)
	assert.NotNil(t, face.face)

	// bad font data
	face, err = lib.NewMemoryFace(font.DejaVuSansMono[1:], 0)
	assert.Error(t, err)
}
