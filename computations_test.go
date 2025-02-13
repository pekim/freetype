package freetype

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"modernc.org/libc"
)

func TestMulDiv(t *testing.T) {
	assert.Equal(t, Long(4), MulDiv(2, 6, 3))
}

func TestMulFix(t *testing.T) {
	assert.Equal(t, Long(0x20000*0x30000/0x10000), MulFix(0x20000, 0x30000))
}

func TestDivFix(t *testing.T) {
	assert.Equal(t, Long(3*0x10000/2), DivFix(3, 2))

}

func TestRoundFix(t *testing.T) {
	assert.Equal(t, Fixed(4<<16), RoundFix(4<<16+3))
	assert.Equal(t, Fixed(4<<16), RoundFix(4<<16-3))
}

func TestCeilFix(t *testing.T) {
	assert.Equal(t, Fixed(4<<16), RoundFix(4<<16-3))
}

func TestFloorFix(t *testing.T) {
	assert.Equal(t, Fixed(4<<16), RoundFix(4<<16+3))
}

func TestVectorTransform(t *testing.T) {
	assert.Equal(t, Vector{20, 80}, VectorTransform(Vector{10, 20}, Matrix{0, 0x10000, 0x20000, 0x30000}))
}

func TestMatrixMultiply(t *testing.T) {
	assert.Equal(t, Matrix{0x20000, 0x30000, 0x80000, 0xB0000}, MatrixMultiply(
		Matrix{0x00000, 0x10000, 0x20000, 0x30000},
		Matrix{0x10000, 0x10000, 0x20000, 0x30000},
	))
}

func TestMatrixInvert(t *testing.T) {
	inverted, err := MatrixInvert(Matrix{0x00000, 0x10000, 0x20000, 0x30000})
	assert.Nil(t, err)
	assert.Equal(t, Matrix{-0x18000, 0x08000, 0x10000, 0x00000}, inverted)

}

func TestAngleDiff(t *testing.T) {
	assert.Equal(t, AnglePI, AngleDiff(Angle2PI, AnglePI))
	assert.Equal(t, AnglePI, AngleDiff(AnglePI, Angle2PI))
}

func TestVectorUnit(t *testing.T) {
	assert.Equal(t, Vector{0x10000, 0}, VectorUnit(Angle2PI))
}

func TestVectorRotate(t *testing.T) {
	assert.Equal(t, Vector{-0x10000, 0}, VectorRotate(libc.NewTLS(), Vector{0x10000, 0}, AnglePI))
}

func TestVectorLength(t *testing.T) {
	assert.Equal(t, Fixed(0x10000), VectorLength(libc.NewTLS(), Vector{0x10000, 0}))
}

func TestVectorPolarize(t *testing.T) {
	length, angle := VectorPolarize(libc.NewTLS(), Vector{0, 0x10000})
	assert.Equal(t, Fixed(0x10000), length)
	assert.Equal(t, AnglePI2, angle)
}

func TestVectorFrom_Polar(t *testing.T) {
	assert.Equal(t, Vector{0, 0x10000}, VectorFromPolar(libc.NewTLS(), Fixed(0x10000), AnglePI2))
}
