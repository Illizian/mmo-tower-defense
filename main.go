package main

import (
	"context"
	"fmt"

	"mmo-tower-defense/pkg/entities"
	"mmo-tower-defense/pkg/maths"
	"mmo-tower-defense/pkg/renderer"
	"mmo-tower-defense/pkg/terminal"

	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool)
	input := make(chan []byte)
	tick := time.NewTicker(200 * time.Millisecond)
	size := 40

	pip := maths.NewRandomVec2(0, size-1)

	var snakes []entities.Snake
	snakes = append(snakes, entities.Snake{
		Color:     terminal.Green,
		Location:  maths.NewRandomVec2(0, size),
		Direction: maths.East,
		Length:    3,
		Path:      make([]maths.Vec2, 0),
		Alive:     true,
	})

	defer func() {
		fmt.Printf(terminal.Reset + terminal.ClearScreen + terminal.ResetCursor + terminal.CursorShow)
		fmt.Println("Thanks for playing...")
	}()

	fmt.Print(terminal.CursorHide)

	go renderer.StdIn(ctx, input)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case char := <-input:
				switch char[0] {
				case 3: // Ctrl-c
					done <- true
				case 119: // W
					snakes[0].Direction = maths.North
				case 97: // A
					snakes[0].Direction = maths.West
				case 115: // S
					snakes[0].Direction = maths.South
				case 100: // D
					snakes[0].Direction = maths.East
				}
				break
			case <-tick.C:
				// create an map of occupied tiles
				occupied := make(map[maths.Vec2]bool)
				for _, snake := range snakes {
					occupied[snake.Location] = true

					for _, tail := range snake.Path {
						occupied[tail] = true
					}
				}

				// Tick each snake with the generated occupied for collisions
				for s := range snakes {
					location := snakes[s].Tick(occupied, size)
					if location.Eq(pip) {
						snakes[s].Length += 1
						pip = maths.NewRandomVec2(0, size-1)
					}
				}

				fmt.Printf(terminal.ClearScreen + terminal.ResetCursor)
				fmt.Println(renderer.Render(snakes, pip, size))

				break
			}
		}
	}(ctx)

	<-done
	cancel()
}
