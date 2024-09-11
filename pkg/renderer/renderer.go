package renderer

import (
	"fmt"
	"mmo-tower-defense/pkg/entities"
	"mmo-tower-defense/pkg/maths"
	"mmo-tower-defense/pkg/terminal"
)

const (
	BORDER_TL = "┏"
	BORDER_BL = "┗"
	BORDER_TR = "┓"
	BORDER_BR = "┛"
	BORDER_VL = "┃"
	BORDER_HL = "━"
)

func Render(snakes []*entities.Snake, pip maths.Vec2, size int) string {
	board := make([]string, size*size)

	for _, snake := range snakes {
		if snake.Status == entities.SNAKE_DEAD {
			continue
		}

		board[snake.Location.ToInt(size)] = fmt.Sprintf(
			"%s%s%s",
			snake.GetColor(),
			snake.GetHead(),
			terminal.Reset,
		)

		for _, path := range snake.Path {
			board[path.ToInt(size)] = fmt.Sprintf(
				"%s%s%s",
				snake.GetColor(),
				snake.GetTail(),
				terminal.Reset,
			)
		}
	}

	board[pip.ToInt(size)] = fmt.Sprintf("%s•%s", terminal.Yellow, terminal.Reset)

	// Create a board with a border
	var output string
	for i := 0; i < (size+2)*(size+2); i++ {
		pos := maths.NewVec2FromInt(i, size+2)
		x := pos.X
		y := pos.Y

		if x == 0 && y == 0 {
			output += BORDER_TL
			continue
		}

		if x == size+1 && y == 0 {
			output += BORDER_TR
			output += "\n\r"
			continue
		}

		if x == 0 && y == size+1 {
			output += BORDER_BL
			continue
		}

		if x == size+1 && y == size+1 {
			output += BORDER_BR
			output += "\n\r"
			continue
		}

		if x == 0 || x == size+1 {
			output += BORDER_VL
			if x == size+1 {
				output += "\n\r"
			}
			continue
		}

		if y == 0 || y == size+1 {
			output += BORDER_HL
			continue
		}

		cell := board[maths.NewVec2(x-1, y-1).ToInt(size)]
		if cell == "" {
			output += " "
			continue
		}

		output += cell
	}

	return output
}
