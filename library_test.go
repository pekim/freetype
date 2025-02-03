package freetype

import (
	"testing"

	"github.com/pekim/freetype-go/internal/font"
	"github.com/stretchr/testify/assert"
)

func TestLibraryInitDone(t *testing.T) {
	lib, err := Init()
	assert.Nil(t, err)
	assert.NotNil(t, lib.library)

	err = lib.Done()
	assert.Nil(t, err)
}

func TestLibraryVersion(t *testing.T) {
	lib, _ := Init()
	defer func() { _ = lib.Done() }()

	major, minor, patch := lib.Version()
	assert.Equal(t, 2, major)
	assert.Greater(t, minor, 0)
	assert.GreaterOrEqual(t, patch, 0)
}

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
