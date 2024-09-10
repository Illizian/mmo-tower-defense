package entities

import (
	"fmt"
	"mmo-tower-defense/pkg/maths"
	"mmo-tower-defense/pkg/terminal"
)

type Snake struct {
	Color     string
	Location  maths.Vec2
	Direction maths.Vec2
	Length    int
	Path      []maths.Vec2
	Alive     bool
}

func (s Snake) GetColor() string {
	if s.Alive {
		return s.Color
	}

	return terminal.Red
}

func (s Snake) GetHead() string {
	switch s.Direction {
	case maths.North:
		return "▲"
	case maths.South:
		return "▼"
	case maths.East:
		return "►"
	case maths.West:
		return "◄"
	}

	return "@"
}

func (s Snake) GetTail() string {
	return "*"
}

func (s *Snake) Tick(occupied map[maths.Vec2]bool, size int) maths.Vec2 {
	if s.Alive == false {
		// @TODO: This feels "wrong", find a better way, should this even return a maths.Vec2?
		return maths.Vec2{X: -1, Y: -1}
	}

	destination := maths.Vec2{
		X: (s.Location.X + s.Direction.X%size + size) % size,
		Y: (s.Location.Y + s.Direction.Y%size + size) % size,
	}

	if occupied[destination] {
		s.Alive = false
	} else {
		s.Path = append([]maths.Vec2{s.Location}, s.Path[:min(len(s.Path), s.Length-1)]...)
		s.Location = destination
	}

	return s.Location
}

func (s Snake) Debug() string {
	return fmt.Sprintf("Snake{location: %+v, direction: %+v, length: %d, path: %+v, alive: %t}", s.Location, s.Direction, s.Length, s.Path, s.Alive)
}
