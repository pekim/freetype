package freetype

import (
	"testing"

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
	assert.Equal(t, 12, minor)
	assert.Equal(t, 1, patch)
}
