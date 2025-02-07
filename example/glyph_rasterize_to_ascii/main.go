package main

import (
	"fmt"

	"github.com/pekim/freetype-go"
	"github.com/pekim/freetype-go/internal/font"
)

// Show how initialize the library, load a font, rasterize a glyph, and dump
// an ascii representation of the rasterization to stdout.

func main() {
	// Initialize an instance of Library,
	lib, err := freetype.Init()
	if err != nil {
		panic(err)
	}

	// Load a Face from font data.
	face, err := lib.NewMemoryFace(font.DejaVuSansMono, 0)
	if err != nil {
		panic(err)
	}

	// Load the glyph for the codepoint 'A'.
	err = face.LoadGlyph(face.GetCharIndex('A'), freetype.LOAD_DEFAULT)
	if err != nil {
		panic(err)
	}

	// Render the glyph as a bitmap.
	err = face.RenderGlyph(freetype.RENDER_MODE_NORMAL)
	if err != nil {
		panic(err)
	}

	// Print the rasterized bitmap to stdout using the BufferVisualization debug method.
	fmt.Println(face.Rec().Glyph.Bitmap.BufferVisualization())
}
