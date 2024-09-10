package renderer

import (
	"fmt"
	"mmo-tower-defense/pkg/entities"
	"mmo-tower-defense/pkg/maths"
	"mmo-tower-defense/pkg/terminal"
)

func Render(snakes []*entities.Snake, pip maths.Vec2, size int) string {
	cells := make([]string, size*size)
	var output string

	for _, snake := range snakes {
		if snake.Status == entities.SNAKE_DEAD {
			continue
		}

		cells[snake.Location.ToInt(size)] = fmt.Sprintf(
			"%s%s%s",
			snake.GetColor(),
			snake.GetHead(),
			terminal.Reset,
		)

		for _, path := range snake.Path {
			cells[path.ToInt(size)] = fmt.Sprintf(
				"%s%s%s",
				snake.GetColor(),
				snake.GetTail(),
				terminal.Reset,
			)
		}
	}

	cells[pip.ToInt(size)] = fmt.Sprintf("%s•%s", terminal.Yellow, terminal.Reset)

	for i, cell := range cells {
		if i != 0 && i%size == 0 {
			output += "\n\r"
		}

		if cell == "" {
			output += " "
			continue
		}

		output += cell
	}

	return output
}

// const (
// 	BORDER_TOP_LEFT   = "┏"
// 	BORDER_VERTICAL   = "┃"
// 	BORDER_TOP_RIGHT  = "┓"
// 	BORDER_HORIZONTAL = "━"
// )

// func Render(snakes []*entities.Snake, pip maths.Vec2, size int) string {
// 	board := make([]string, size*size)
//
// 	for _, snake := range snakes {
// 		board[snake.Location.ToInt(size)] = fmt.Sprintf(
// 			"%s%s%s",
// 			snake.GetColor(),
// 			snake.GetHead(),
// 			terminal.Reset,
// 		)
//
// 		for _, path := range snake.Path {
// 			board[path.ToInt(size)] = fmt.Sprintf(
// 				"%s%s%s",
// 				snake.GetColor(),
// 				snake.GetTail(),
// 				terminal.Reset,
// 			)
// 		}
// 	}
//
// 	board[pip.ToInt(size)] = fmt.Sprintf("%s•%s", terminal.Yellow, terminal.Reset)
//
// 	// Create a board with a border
// 	var output string
// 	for i := 0; i < (size+2)*(size+2); i++ {
// 		pos := maths.NewVec2FromInt(i, size)
// 		x := pos.X
// 		y := pos.Y
//
// 		if x == 0 && y == 0 {
// 			output += BORDER_TOP_LEFT
// 			continue
// 		}
//}
// 		if x == size+1 && y == size+1 {
// 			output += BORDER_TOP_LEFT
// 			continue
// 		}
//
// 		if x == 0 || x == size+1 {
// 			output += BORDER_VERTICAL
// 			continue
// 		}
//
// 		if y == 0 || y == size+1 {
// 			output += BORDER_HORIZONTAL
// 			continue
// 		}
//
// 		if i != 0 && i%(size+2) == 0 {
// 			output += "\n\r"
// 		}
//
// 		cell := board[maths.NewVec2(x-1, y-1).ToInt(size)]
// 		if cell == "" {
// 			output += " "
// 			continue
//
//
// 		output += cell
// 	}
//
// 	return output
// }
