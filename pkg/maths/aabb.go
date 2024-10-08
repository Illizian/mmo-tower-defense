package maths

type AABB struct {
	A Vec2
	B Vec2
}

func (aabb AABB) Contains(vec Vec2) bool {
	return (vec.X >= aabb.A.X && vec.X <= aabb.B.X &&
		vec.Y >= aabb.A.Y && vec.Y <= aabb.B.Y)
}

func (aabb AABB) Len() int {
	return (aabb.B.X - aabb.A.X) * (aabb.B.Y - aabb.A.Y)
}

func NewAABB(a, b Vec2) AABB {
	return AABB{
		A: a,
		B: b,
	}
}
