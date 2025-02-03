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
	major, minor, patch := lib.Version()
	assert.Equal(t, 2, major)
	assert.Greater(t, minor, 0)
	assert.GreaterOrEqual(t, patch, 0)
}
