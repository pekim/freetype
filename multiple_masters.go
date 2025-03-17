package freetype

import (
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libfreetype"
)

// How to manage Multiple Masters fonts.

const (
	// The maximum number of Multiple Masters axes.
	T1_MAX_MM_AXIS = 4
	// The maximum number of Multiple Masters designs.
	T1_MAX_MM_DESIGNS = 16
	// The maximum number of elements in a design map.
	T1_MAX_MM_MAP_POINTS = 20
)

// A structure to model a given axis in design space for Multiple Masters fonts.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_mm_axis
type MMAxis struct {
	name    uintptr
	Minimum Long
	Maximum Long
}

func (mma MMAxis) Name() string {
	return libc.GoString(mma.name)
}

// A structure to model the axes and space of a Multiple Masters font.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_multi_master
type MultiMaster struct {
	NumAxis    UInt
	NumDesigns UInt
	Axis       [T1_MAX_MM_AXIS]MMAxis
}

// A structure to model a given axis in design space for Multiple Masters, TrueType GX,
// and OpenType variation fonts.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_var_axis
type VarAxis struct {
	name uintptr

	Minimum Fixed
	Def     Fixed
	Maximum Fixed

	Tag   ULong
	Strid UInt
}

func (va VarAxis) Name() string {
	return libc.GoString(va.name)
}

// A structure to model a named instance in a TrueType GX or OpenType variation font.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_var_named_style
type VarNamedStyle struct {
	Coords *Fixed
	Strid  UInt
	Psid   UInt /* since 2.7.1 */
}

// A structure to model the axes and space of an Adobe MM, TrueType GX, or OpenType variation font.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_mm_var
type MMVar struct {
	numAxis        UInt
	NumDesigns     UInt
	numNamedstyles UInt
	axis           *VarAxis
	namedstyle     *VarNamedStyle
}

func (mmvar MMVar) Axes() []VarAxis {
	return unsafe.Slice(mmvar.axis, mmvar.numAxis)
}

func (mmvar MMVar) NamedStyles() []VarNamedStyle {
	return unsafe.Slice(mmvar.namedstyle, mmvar.numNamedstyles)
}

// Retrieve a variation descriptor of a given Adobe MM font.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_get_multi_master
func (face Face) GetMultiMaster() (MultiMaster, error) {
	var master MultiMaster
	err := libfreetype.XFT_Get_Multi_Master(face.tls, face.face, toUintptr(&master))
	return master, newError(err, "failed to get multi master")
}

// Retrieve a variation descriptor for a given font.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_get_mm_var
func (face Face) GetMMVar() (*MMVar, error) {
	var ptrMaster *MMVar
	err := libfreetype.XFT_Get_MM_Var(face.tls, face.face, toUintptr(&ptrMaster))
	return ptrMaster, newError(err, "failed to get multi master variation")
}

// Free the memory allocated by GetMMVar.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_done_mm_var
func (library Library) DoneMMVar(master *MMVar) error {
	err := libfreetype.XFT_Done_MM_Var(library.tls, library.library, toUintptr(master))
	return newError(err, "failed to free a multi master variation")
}

// For Adobe MM fonts, choose an interpolated font design through design coordinates.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_set_mm_design_coordinates
func (face Face) SetMMDesignCoordinates(coords []Long) error {
	err := libfreetype.XFT_Set_MM_Design_Coordinates(face.tls, face.face, UInt(len(coords)), toUintptr(&coords[0]))
	return newError(err, "failed to set multi master design coordinates")
}

// Choose an interpolated font design through design coordinates.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_set_var_design_coordinates
func (face Face) SetVarDesignCoordinates(coords []Fixed) error {
	err := libfreetype.XFT_Set_Var_Design_Coordinates(face.tls, face.face, UInt(len(coords)), toUintptr(&coords[0]))
	return newError(err, "failed to set multi master variation design coordinates")
}

// Get the design coordinates of the currently selected interpolated font.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_get_var_design_coordinates
func (face Face) GetVarDesignCoordinates(numCoords UInt) ([]Fixed, error) {
	coords := make([]Fixed, numCoords)
	err := libfreetype.XFT_Get_Var_Design_Coordinates(face.tls, face.face, numCoords, toUintptr(&coords[0]))
	return coords, newError(err, "failed to get multi master variation design coordinates")
}

// Choose an interpolated font design through normalized blend coordinates.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_set_mm_blend_coordinates
func (face Face) SetMMBlendCoordinates(coords []Fixed) error {
	err := libfreetype.XFT_Set_MM_Blend_Coordinates(face.tls, face.face, UInt(len(coords)), toUintptr(&coords[0]))
	return newError(err, "failed to set multi master blend coordinates")
}

// Get the normalized blend coordinates of the currently selected interpolated font.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_get_mm_blend_coordinates
func (face Face) GetMMBlendCoordinates(numCoords UInt) ([]Fixed, error) {
	coords := make([]Fixed, numCoords)
	err := libfreetype.XFT_Get_MM_Blend_Coordinates(face.tls, face.face, numCoords, toUintptr(&coords[0]))
	return coords, newError(err, "failed to get multi master blend coordinates")
}

// This is another name of SetMMBlendCoordinates.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_set_var_blend_coordinates
func (face Face) SetVarBlendCoordinates(coords []Fixed) error {
	return face.SetMMBlendCoordinates(coords)
}

// This is another name of GetMMBlendCoordinates.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_get_var_blend_coordinates
func (face Face) GetVarBlendCoordinates(numCoords UInt) ([]Fixed, error) {
	return face.GetMMBlendCoordinates(numCoords)
}

// For Adobe MM fonts, choose an interpolated font design by directly setting the weight vector.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_set_mm_weightvector
func (face Face) SetMMWeightVector(weightVector []Fixed) error {
	err := libfreetype.XFT_Set_MM_WeightVector(face.tls, face.face, UInt(len(weightVector)), toUintptr(&weightVector[0]))
	return newError(err, "failed to set multi master weight vector")
}

// For Adobe MM fonts, retrieve the current weight vector of the font.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_get_mm_weightvector
func (face Face) GetMMWeightVector(length UInt) (UInt, []Fixed, error) {
	weightVector := make([]Fixed, length)
	err := libfreetype.XFT_Get_MM_WeightVector(face.tls, face.face, toUintptr(&length), toUintptr(&weightVector[0]))
	return length, weightVector, newError(err, "failed to get multi master weight vector")
}

// The variation axis should not be exposed to user interfaces.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_var_axis_flag_xxx
const FT_VAR_AXIS_FLAG_HIDDEN = UInt(1)

// Get the ‘flags’ field of an OpenType Variation Axis Record.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_get_var_axis_flags
func (face Face) GetVarAxisFlags(master *MMVar, axisIndex UInt) (UInt, error) {
	var flags UInt
	err := libfreetype.XFT_Get_Var_Axis_Flags(face.tls, toUintptr(master), axisIndex, toUintptr(&flags))
	return flags, newError(err, "failed to get multi master flags")
}

// Set or change the current named instance..
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_set_named_instance
func (face Face) SetNamedInstance(instanceIndex UInt) error {
	err := libfreetype.XFT_Set_Named_Instance(face.tls, face.face, instanceIndex)
	return newError(err, "failed to set named instance")
}

// Retrieve the index of the default named instance, to be used with SetNamedInstance.
//
// https://freetype.org/freetype2/docs/reference/ft2-multiple_masters.html#ft_get_default_named_instance
// func (face Face) GetDefaultNamedInstance() (UInt, error) {
// 	var instanceIndex UInt
// 	err := libfreetype.XFT_Get_Default_Named_Instance(face.tls, face.face, toUintptr(&instanceIndex))
// 	return instanceIndex, newError(err, "failed to get default named instance flags")
// }
