//go:build linux

package freetype

import (
	"fmt"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libfreetype"
)

// OT-SVG API between FreeType and an external SVG rendering library.

// SVGLibInitFunc is a callback that is called when the first OT-SVG glyph is rendered in the
// lifetime of an FT_Library object.
// In a typical implementation, one would want to allocate a structure and point the data_pointer
// to it and perform any library initializations that might be needed.
//
// https://freetype.org/freetype2/docs/reference/ft2-svg_fonts.html#svg_lib_init_func
type SVGLibInitFunc func(tls *libc.TLS, state uintptr) FTError

// SVGLibFreeFunc is a callback that is called when the ot-svg module is being freed.
// It is only called if the init hook was called earlier.
// This means that neither the init nor the free hook is called if no OT-SVG glyph is rendered.
//
// https://freetype.org/freetype2/docs/reference/ft2-svg_fonts.html#svg_lib_free_func
type SVGLibFreeFunc func(tls *libc.TLS, state uintptr)

// SVGLibRenderFunc is a callback that is called to render an OT-SVG glyph.
// This callback hook is called right after the preset hook SVG_Lib_Preset_Slot_Func has been
// called with cache set to TRUE.
// The data necessary to render is available through the handle FT_SVG_Document,
// which is set in the other field of FT_GlyphSlotRec.
//
// https://freetype.org/freetype2/docs/reference/ft2-svg_fonts.html#svg_lib_render_func
type SVGLibRenderFunc func(tls *libc.TLS, slot GlyphSlot, state uintptr) FTError

// SVGLibPresetSlotFunc is a callback that is called to preset the glyph slot. It is called from two places.
//
//   - When FT_Load_Glyph needs to preset the glyph slot.
//   - Right before the svg module calls the render callback hook.
//
// https://freetype.org/freetype2/docs/reference/ft2-svg_fonts.html#svg_lib_preset_slot_func
type SVGLibPresetSlotFunc func(tls *libc.TLS, slot GlyphSlot, cache Bool, state uintptr) FTError

func init() {
	assertSameSize(SVGRendererHooks{}, libfreetype.TSVG_RendererHooks{})
}

// SVGRendererHooks is a structure that stores the four hooks needed to render OT-SVG glyphs properly.
// The structure is publicly used to set the hooks via the svg-hooks driver property.
//
// https://freetype.org/freetype2/docs/reference/ft2-svg_fonts.html#svg_rendererhooks
type SVGRendererHooks struct {
	InitSVG    SVGLibInitFunc
	FreeSvg    SVGLibFreeFunc
	RenderSVG  SVGLibRenderFunc
	PresetSlot SVGLibPresetSlotFunc
}

// SetSVGHooks sets the four hooks needed to render OT-SVG glyphs properly.
func (lib Library) SetSVGHooks(hooks SVGRendererHooks) error {
	hooks_ := libfreetype.TSVG_RendererHooks{
		Finit_svg:    __ccgo_fp(hooks.InitSVG),
		Ffree_svg:    __ccgo_fp(hooks.FreeSvg),
		Frender_svg:  __ccgo_fp(hooks.RenderSVG),
		Fpreset_slot: __ccgo_fp(hooks.PresetSlot),
	}
	err := lib.PropertySet("ot-svg", "svg-hooks", toUintptr(&hooks_))
	if err != nil {
		return fmt.Errorf("failed to set svg hooks for library : %w", err)
	}
	return nil
}

func init() {
	assertSameSize(SVGDocumentRec{}, libfreetype.TFT_SVG_DocumentRec{})
}

// SVGDocumentRec is a structure that models one SVG document.
//
// https://freetype.org/freetype2/docs/reference/ft2-svg_fonts.html#ft_svg_documentrec
type SVGDocumentRec struct {
	svgDocument       *Byte
	svgDocumentLength ULong

	Metrics    SizeMetrics
	UnitsPerEM UShort

	StartGlyphID UShort
	EndGlyphID   UShort

	Transform Matrix
	Delta     Vector
}

func (rec SVGDocumentRec) Document() []byte {
	return unsafe.Slice(rec.svgDocument, rec.svgDocumentLength)
}

// SVGDocument is a handle to an SVGDocumentRec object.
//
// https://freetype.org/freetype2/docs/reference/ft2-svg_fonts.html#ft_svg_document
type SVGDocument libfreetype.TFT_SVG_Document

// Rec returns a pointer to the SVGDocumentRec that is referenced by the SVGDocument.
func (sd SVGDocument) Rec() *SVGDocumentRec {
	return fromUintptr[SVGDocumentRec](uintptr(sd))
}
