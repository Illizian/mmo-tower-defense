package maths

import "math/rand"

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

func NewRandomVec2(min, max int) Vec2 {
	return Vec2{
		X: rand.Intn(max-min+1) + min,
		Y: rand.Intn(max-min+1) + min,
	}
}
