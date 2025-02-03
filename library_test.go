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
