package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAABBContainsVector(t *testing.T) {
	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(0, 0)))
	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(1, 0)))
	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(2, 0)))
	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(3, 0)))

	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(0, 1)))
	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(1, 1)))
	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(2, 1)))
	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(3, 1)))

	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(0, 2)))
	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(1, 2)))
	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(2, 2)))
	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(3, 2)))

	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(0, 3)))
	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(1, 3)))
	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(2, 3)))
	assert.True(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(3, 3)))

	assert.False(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(-1, -1)))
	assert.False(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(1, 5)))
	assert.False(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(2, -1)))
	assert.False(t, NewAABB(NewVec2(0, 0), NewVec2(3, 3)).Contains(NewVec2(5, 5)))
}
