//go:build linux

package freetype

import (
	"testing"

	"github.com/pekim/freetype/internal/font"
	"github.com/stretchr/testify/assert"
	"modernc.org/libc"
)

func TestLibrarySetSVGHooks(t *testing.T) {
	initFuncCalled := false
	var initFunc SVGLibInitFunc = func(_ *libc.TLS, _ uintptr) FTError {
		initFuncCalled = true
		return Err_Ok
	}

	freeFuncCalled := false
	var freeFunc SVGLibFreeFunc = func(_ *libc.TLS, _ uintptr) {
		freeFuncCalled = true
	}

	renderFuncCalled := false
	var renderFunc SVGLibRenderFunc = func(_ *libc.TLS, _ GlyphSlot, _ uintptr) FTError {
		renderFuncCalled = true
		return Err_Ok
	}

	presetSlotFuncCalled := 0
	var presetSlotFunc SVGLibPresetSlotFunc = func(_ *libc.TLS, _ GlyphSlot, _ Bool, _ uintptr) FTError {
		presetSlotFuncCalled++
		return Err_Ok
	}

	lib, err := Init()
	assert.NoError(t, err)

	err = lib.SetSVGHooks(SVGRendererHooks{
		InitSVG:    initFunc,
		FreeSvg:    freeFunc,
		RenderSVG:  renderFunc,
		PresetSlot: presetSlotFunc,
	})
	assert.NoError(t, err)

	face, err := lib.NewMemoryFace(font.NotoColorEmoji, 0)
	assert.NoError(t, err)

	err = face.SetPixelSizes(0, 32)
	assert.NoError(t, err)

	err = face.LoadGlyph(face.GetCharIndex('😀'), LOAD_COLOR)
	assert.NoError(t, err)

	err = face.RenderGlyph(RENDER_MODE_NORMAL)
	assert.NoError(t, err)

	err = lib.Done()
	assert.NoError(t, err)

	assert.True(t, initFuncCalled)
	assert.True(t, freeFuncCalled)
	assert.True(t, renderFuncCalled)
	assert.Equal(t, 2, presetSlotFuncCalled)
}
