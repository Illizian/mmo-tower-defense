package main

import (
	"fmt"
	"mmo-tower-defense/pkg/terminal"
	"time"
)

type Entity interface {
	GetPos() (int, int)
	Tick()
	Label() string
}

type Spawner struct {
	Row    int
	Col    int
	Level  int
	Health int
}

func (s Spawner) GetPos() (int, int) {
	return s.Row, s.Col
}

func (s Spawner) Tick() {
	fmt.Println("Spawner#Tick")
}

func (s Spawner) Label() string {
	return terminal.Magenta + "●" + terminal.Reset
}

type Creep struct {
	Row    int
	Col    int
	Health int
}

func (c Creep) GetPos() (int, int) {
	return c.Row, c.Col
}

func (c Creep) Tick() {
	fmt.Println("Creep#Tick")
}

func (c Creep) Label() string {
	return terminal.Red + "*" + terminal.Reset
}

type Shrine struct {
	Row    int
	Col    int
	Health int
}

func (s Shrine) GetPos() (int, int) {
	return s.Row, s.Col
}

func (s Shrine) Tick() {
	fmt.Println("Shrine#Tick")
}

func (s Shrine) Label() string {
	return terminal.Cyan + "@" + terminal.Reset
}

// type Cell struct {
// 	Type     int
// 	Entities []Entity
// }
//
// func (c Cell) Label() string {
// 	if c.Type == -1 {
// 		return "░"
// 	}
//
// 	if len(c.Entities) > 0 {
// 		return c.Entities[0].Label()
// 	}
//
// 	return "x"
// }

// func getGridSize() (int, int, int, int) {
// 	return 0, 10, 0, 40
// minRow := 0
// maxRow := 0
// minCol := 0
// maxCol := 0
// for _, cell := range cells {
// 	if cell.Row < minRow {
// 		minRow = cell.Row
// 	}
//
// 	if cell.Row > maxRow {
// 		maxRow = cell.Row
// 	}
//
// 	if cell.Col < minCol {
// 		minCol = cell.Col
// 	}
//
// 	if cell.Col > maxCol {
// 		maxCol = cell.Col
// 	}
// }
// return minRow, maxRow, minCol, maxCol
// }

func outOfBounds(entity Entity, startRow, endRow, startCol, endCol int) bool {
	row, col := entity.GetPos()

	return false &&
		(row >= startRow && row <= endRow) &&
		(col >= startCol && col <= endCol)

}

func indexFromPos(row, col, size int) int {
	return row*size + col
}

func render(entities []Entity, startRow, endRow, startCol, endCol int) string {
	width := endCol - startCol
	height := endRow - startRow
	cells := make([]string, width*height)

	var output string

	for _, entity := range entities {
		if !outOfBounds(entity, startRow, endRow, startCol, endCol) {
			row, col := entity.GetPos()

			cells[col+row*width] = entity.Label()
		}
	}

	for i, cell := range cells {
		if i%width == 0 {
			output += "\n"
		}

		if cell == "" {
			output += " "
		} else {
			output += cell
		}
	}

	return output
}

func main() {
	fmt.Println("Launching...")

	// minRow, maxRow, minCol, maxCol := getGridSize()
	var entities []Entity
	// board := make([]Cell, size*size)
	tick := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	entities = append(entities, Spawner{Row: 2, Col: 4, Level: 1, Health: 100})
	entities = append(entities, Creep{Row: 5, Col: 4, Health: 100})
	entities = append(entities, Shrine{Row: 8, Col: 4, Health: 100})

	defer func() {
		fmt.Printf(terminal.Reset + terminal.ClearScreen + terminal.ResetCursor + terminal.CursorShow)
		fmt.Println("Exiting...")
	}()

	fmt.Print(terminal.CursorHide)

	go func() {
		for {
			select {
			case <-tick.C:
				// for cellIndex := range board {
				// 	for entityIndex := range board[cellIndex].Entities {
				// 		board[cellIndex].Entities[entityIndex].Tick()
				// 	}
				// }

				fmt.Printf(terminal.ClearScreen + terminal.ResetCursor)
				fmt.Printf(render(entities, 0, 10, 0, 10))
				// 			fmt.Print(render + "\n")
				// 			fmt.Printf("{%d}:{%d}, {%d}:{%d}\n", minRow, maxRow, minCol, maxCol)
				// 			fmt.Println(cells)
			}
		}
	}()

	<-done
}
