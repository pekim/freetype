package freetype

import (
	"modernc.org/libc"
	"modernc.org/libfreetype"
)

// Crunching fixed numbers and vectors.

// MulDiv computes (a*b)/c with maximum accuracy, using a 64-bit intermediate integer whenever necessary.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_muldiv
func MulDiv(a Long, b Long, c Long) Long {
	return libfreetype.XFT_MulDiv(nil, a, b, c)
}

// MulFix computes (a*b)/0x10000 with maximum accuracy.
// Its main use is to multiply a given value by a 16.16 fixed-point factor.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_mulfix
func MulFix(a Long, b Long) Long {
	return libfreetype.XFT_MulFix(nil, a, b)
}

// DivFix computes (a*0x10000)/b with maximum accuracy.
// Its main use is to divide a given value by a 16.16 fixed-point factor.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_divfix
func DivFix(a Long, b Long) Long {
	return libfreetype.XFT_DivFix(nil, a, b)
}

// RoundFix rounds a 16.16 fixed number.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_roundfix
func RoundFix(a Long) Long {
	return libfreetype.XFT_RoundFix(nil, a)
}

// CeilFix computes the smallest following integer of a 16.16 fixed number.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_ceilfix
func CeilFix(a Long) Long {
	return libfreetype.XFT_CeilFix(nil, a)
}

// FloorFix computes the largest previous integer of a 16.16 fixed number.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_floorfix
func FloorFix(a Long) Long {
	return libfreetype.XFT_FloorFix(nil, a)
}

// VectorTransform transforms a single vector through a 2x2 matrix.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_vector_transform
func VectorTransform(vector Vector, matrix Matrix) Vector {
	outVector := vector
	libfreetype.XFT_Vector_Transform(nil, toUintptr(&outVector), toUintptr(&matrix))
	return outVector
}

// MatrixMultiply returns the matrix operation a*b.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_matrix_multiply
func MatrixMultiply(a Matrix, b Matrix) Matrix {
	outB := b
	libfreetype.XFT_Matrix_Multiply(nil, toUintptr(&a), toUintptr(&outB))
	return outB
}

// MatrixInvert inverts a 2x2 matrix. Returns an error if it can't be inverted.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_matrix_invert
func MatrixInvert(matrix Matrix) (Matrix, error) {
	outMatrix := matrix
	err := libfreetype.XFT_Matrix_Invert(nil, toUintptr(&outMatrix))
	return outMatrix, newError(err, "failed to invert matrix")
}

// Angle is used to model angle values in FreeType.
// Note that the angle is a 16.16 fixed-point value expressed in degrees.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_angle
type Angle = libfreetype.TFT_Angle

// AnglePI is the angle pi expressed in Angle units.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_angle_pi
const AnglePI = Angle(180 << 16)

// Angle2PI is the angle 2*pi expressed in Angle units.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_angle_2pi
const Angle2PI = Angle(AnglePI * 2)

// AnglePI2 is the angle pi/2 expressed in Angle units.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_angle_pi2
const AnglePI2 = Angle(AnglePI / 2)

// AnglePI4 is the angle pi/4 expressed in Angle units.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_angle_pi4
const AnglePI4 = Angle(AnglePI / 4)

// The trignometric functions require an instance of libc.TLS.
// In any case Go's math package is probably just as good as these functions.
//
// FT_Sin
// FT_Cos
// FT_Tan
// FT_Atan2

// AngleDiff returns the difference between two angles. The result is always constrained to the ]-PI..PI] interval.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_angle_diff
func AngleDiff(angle1 Angle, angle2 Angle) Angle {
	return libfreetype.XFT_Angle_Diff(nil, angle1, angle2)
}

// VectorUnit returns the unit vector corresponding to a given angle.
// After the call, the value of vec.x will be cos(angle), and the value of vec.y will be sin(angle).
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_vector_unit
func VectorUnit(angle Angle) Vector {
	var vec Vector
	libfreetype.XFT_Vector_Unit(nil, toUintptr(&vec), angle)
	return vec
}

// VectorRotate rotates a vector by a given angle.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_vector_rotate
func VectorRotate(tls *libc.TLS, vec Vector, angle Angle) Vector {
	outVec := vec
	libfreetype.XFT_Vector_Rotate(tls, toUintptr(&outVec), angle)
	return outVec
}

// VectorLength returns the length of a given vector.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_vector_length
func VectorLength(tls *libc.TLS, vec Vector) Fixed {
	return libfreetype.XFT_Vector_Length(tls, toUintptr(&vec))
}

// VectorPolarize computes both the length and angle of a given vector.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_vector_polarize
func VectorPolarize(tls *libc.TLS, vec Vector) (Fixed, Angle) {
	var length Fixed
	var angle Angle
	libfreetype.XFT_Vector_Polarize(tls, toUintptr(&vec), toUintptr(&length), toUintptr(&angle))
	return length, angle
}

// VectorFromPolar computes vector coordinates from a length and angle.
//
// https://freetype.org/freetype2/docs/reference/ft2-computations.html#ft_vector_from_polar
func VectorFromPolar(tls *libc.TLS, length Fixed, angle Angle) Vector {
	var outVec Vector
	libfreetype.XFT_Vector_From_Polar(tls, toUintptr(&outVec), length, angle)
	return outVec
}
