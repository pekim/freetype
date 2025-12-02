package freetype

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pekim/freetype/internal/font"
)

func TestGetMMVar(t *testing.T) {
	lib, _ := Init()
	face, _ := lib.NewMemoryFace(font.RobotoVariable, 0)

	mmVar, err := face.GetMMVar()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(mmVar.Axes()))

	expectedNames := []string{"Weight", "Width"}
	expectedTags := []Tag{imageTag('w', 'g', 'h', 't'), imageTag('w', 'd', 't', 'h')}
	for i, axis := range mmVar.Axes() {
		assert.Equal(t, expectedNames[i], axis.Name())
		assert.Equal(t, expectedTags[i], Tag(axis.Tag))
	}
}
