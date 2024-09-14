package maths

import (
	"math"
	"math/rand"
)

type Vec2 struct {
	X int
	Y int
}

func (v Vec2) Eq(other Vec2) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v Vec2) ToInt(size int) int {
	return v.X + v.Y*size
}

func NewVec2(x, y int) Vec2 {
	return Vec2{X: x, Y: y}
}

func NewVec2FromInt(i, size int) Vec2 {
	x := i % size
	y := int(math.Floor(float64(i) / float64(size)))

	return NewVec2(x, y)
}

func NewRandomVec2(min, max int) Vec2 {
	return Vec2{
		X: rand.Intn(max-min+1) + min,
		Y: rand.Intn(max-min+1) + min,
	}
}
