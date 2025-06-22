//go:build linux

package freetype

func (rec GlyphSlotRec) SVGDocument() *SVGDocumentRec {
	return fromUintptr[SVGDocumentRec](rec.Other)
}
