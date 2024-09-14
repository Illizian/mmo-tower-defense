package entities

import (
	"fmt"
	"mmo-tower-defense/pkg/maths"
	"mmo-tower-defense/pkg/terminal"
)

type SnakeStatus int

const (
	SNAKE_ALIVE  = 0
	SNAKE_DIEING = 1
	SNAKE_DEAD   = 2
)

type Snake struct {
	Label       string
	Color       string
	Location    maths.Vec2
	Direction   maths.Vec2
	Length      int
	Path        []maths.Vec2
	Status      SnakeStatus
	DeadCounter int
}

func (s Snake) GetColor() string {
	if s.Status == SNAKE_ALIVE {
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
	if s.Status == SNAKE_DIEING || s.Status == SNAKE_DEAD {
		// @TODO: This feels "wrong", find a better way, should this even return a maths.Vec2?
		return maths.Vec2{X: -1, Y: -1}
	}

	destination := maths.Vec2{
		X: (s.Location.X + s.Direction.X%size + size) % size,
		Y: (s.Location.Y + s.Direction.Y%size + size) % size,
	}

	if occupied[destination] {
		s.Status = SNAKE_DIEING
	} else {
		s.Path = append([]maths.Vec2{s.Location}, s.Path[:min(len(s.Path), s.Length-1)]...)
		s.Location = destination
	}

	return s.Location
}

func (s Snake) Debug() string {
	return fmt.Sprintf("Snake{label: %s, location: %+v, direction: %+v, length: %d, path: %+v, status: %v}", s.Label, s.Location, s.Direction, s.Length, s.Path, s.Status)
}
