package main

import (
	"image"
	"image/png"
	"os"

	"github.com/pekim/freetype-go"
	"github.com/pekim/freetype-go/internal/font"
)

func main() {
	// Initialize an instance of Library,
	lib, err := freetype.Init()
	if err != nil {
		panic(err)
	}

	// Load a Face from font data.
	face, err := lib.NewMemoryFace(font.DejaVuSans, 0)
	if err != nil {
		panic(err)
	}

	// Set the nominal height.
	err = face.SetPixelSizes(0, 16)
	if err != nil {
		panic(err)
	}

	image := image.NewRGBA(image.Rect(0, 0, 400, 300))

	// Draw glyphs in to image
	x := 10
	y := 30
	for _, r := range "FreeType" {
		advance := drawGlyphAt(image, face, r, x, y)
		x += int(advance / 64)
	}

	// Write the image to a file
	outputFile, err := os.Create("example/image_with_text/image_with_text.png")
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
