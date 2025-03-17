package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/pekim/freetype"
	"github.com/pekim/freetype/internal/font"
)

func main() {
	// Initialize an instance of Library,
	lib, err := freetype.Init()
	if err != nil {
		panic(err)
	}

	// Load a Face from font data.
	face, err := lib.NewMemoryFace(font.RobotoVariable, 0)
	if err != nil {
		panic(err)
	}

	// Set the nominal height.
	err = face.SetPixelSizes(0, 32)
	if err != nil {
		panic(err)
	}

	lineHeight := int(face.Rec().Size.Rec().Metrics.Height / 64)

	mmVar, err := face.GetMMVar()
	if err != nil {
		panic(err)
	}
	weightAxis := mmVar.Axes()[0]
	if weightAxis.Name() != "Weight" {
		panic(fmt.Sprintf("%q != %q", weightAxis.Name(), "Weight"))
	}
	widthAxis := mmVar.Axes()[1]
	if widthAxis.Name() != "Width" {
		panic(fmt.Sprintf("%q != %q", widthAxis.Name(), "Width"))
	}

	image := image.NewRGBA(image.Rect(0, 0, 400, 300))

	// Draw glyphs in to image
	x := 10
	y := 30
	drawString(image, face, "Default weight", x, y)

	// width, narrowest
	y += lineHeight
	setVariation(face, weightAxis.Def, widthAxis.Minimum)
	drawString(image, face, "Default weight, narrow", x, y)

	// weight, lightest
	y += lineHeight
	setVariation(face, weightAxis.Minimum, widthAxis.Def)
	drawString(image, face, "Light weight", x, y)

	// weight, heaviest
	y += lineHeight
	setVariation(face, weightAxis.Maximum, widthAxis.Def)
	drawString(image, face, "Heavy weight", x, y)

	// Write the image to a file
	outputFile, err := os.Create("example/image_with_variable_font_text/image_with_variable_font_text.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(outputFile, image)
	if err != nil {
		panic(err)
	}
	err = outputFile.Close()
	if err != nil {
		panic(err)
	}
}

func drawGlyphAt(image *image.RGBA, face freetype.Face, codepoint rune, penX int, penY int) freetype.Pos {
	// Load the glyph.
	err := face.LoadGlyph(face.GetCharIndex(codepoint), freetype.LOAD_DEFAULT)
	if err != nil {
		panic(err)
	}

	// Render the glyph as a bitmap.
	err = face.RenderGlyph(freetype.RENDER_MODE_NORMAL)
	if err != nil {
		panic(err)
	}
	glyph := face.Rec().Glyph.Rec()

	y := penY - int(glyph.BitmapTop)
	bitmap := glyph.Bitmap
	for yGlyph := 0; yGlyph < int(bitmap.Rows); yGlyph++ {
		x := penX + int(glyph.BitmapLeft)
		for xGlyph := 0; xGlyph < int(bitmap.Width); xGlyph++ {
			imageIndex := (y * image.Stride) + (4 * x)
			bitmapIndex := (yGlyph * int(bitmap.Pitch)) + xGlyph
			alpha := bitmap.Buffer()[bitmapIndex]
			gray := 0xff - alpha
			image.Pix[imageIndex+0] = gray // R
			image.Pix[imageIndex+1] = gray // G
			image.Pix[imageIndex+2] = gray // B
			image.Pix[imageIndex+3] = 0xFF // A
			// image.Pix[imageIndex+0] = 0x00                         // R
			// image.Pix[imageIndex+1] = 0x00                         // G
			// image.Pix[imageIndex+2] = 0x00                         // B
			// image.Pix[imageIndex+3] = bitmap.Buffer()[bitmapIndex] // A
			x++
		}
		y++
	}

	return glyph.Advance.X
}

func drawString(image *image.RGBA, face freetype.Face, text string, x int, y int) {
	for _, r := range text {
		advance := drawGlyphAt(image, face, r, x, y)
		x += int(advance / 64)
	}
}

func setVariation(face freetype.Face, weight freetype.Fixed, width freetype.Fixed) {
	err := face.SetVarDesignCoordinates([]freetype.Fixed{weight, width})
	if err != nil {
		panic(err)
	}
}
