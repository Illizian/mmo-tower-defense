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
