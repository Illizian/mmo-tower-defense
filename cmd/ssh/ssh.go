package main

import (
	"context"
	"fmt"
	"io"

	"mmo-tower-defense/pkg/entities"
	"mmo-tower-defense/pkg/maths"
	"mmo-tower-defense/pkg/renderer"
	"mmo-tower-defense/pkg/terminal"

	"time"

	"github.com/gliderlabs/ssh"
)

func main() {
	done := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tick := time.NewTicker(200 * time.Millisecond)
	size := 20
	pip := maths.NewRandomVec2(0, size-1)

	var snakes []*entities.Snake

	ssh.Handle(func(s ssh.Session) {
		addr := s.RemoteAddr()
		fmt.Printf("[%s] Client Connected\n", addr)

		s.Write([]byte(terminal.CursorHide))
		snek := entities.Snake{
			Color:     terminal.Green,
			Location:  maths.NewRandomVec2(0, size),
			Direction: maths.East,
			Length:    3,
			Path:      make([]maths.Vec2, 0),
			Alive:     true,
		}

		snakes = append(snakes, &snek)

		framer := time.NewTicker(50 * time.Millisecond)
		reader, writer := io.Pipe()

		go func() {
			buf := make([]byte, 256)
			for {
				_, err := reader.Read(buf)
				if err == io.EOF {
					fmt.Printf("[%s] Error: Received EOF from client stdin\n", addr)
					break
				}

				if err != nil {
					fmt.Printf("[%s] Error: reading from ssh client stdin\n", addr)
				}

				switch buf[0] {
				case 3: // Ctrl-c
					fmt.Printf("[%s] Client Disconnected\n", addr)
					snek.Alive = false
					s.Write([]byte(terminal.CursorShow))
					s.Exit(0)
					return
				case 119: // W
					snek.Direction = maths.North
				case 97: // A
					snek.Direction = maths.West
				case 115: // S
					snek.Direction = maths.South
				case 100: // D
					snek.Direction = maths.East
				}
			}
		}()

		go func() {
			for {
				select {
				case <-ctx.Done():
					s.Write([]byte(fmt.Sprintf("%s%sServer shutting down... Goodbye!", terminal.ClearScreen, terminal.ResetCursor)))
					s.Exit(0)
					return
				case <-framer.C:
					s.Write([]byte(fmt.Sprintf(terminal.ClearScreen + terminal.ResetCursor)))
					// @TODO: Should probably only render once, and share with everyone! Can we just pipe in a channel? From the renderer?
					s.Write([]byte(renderer.Render(snakes, pip, size)))
					break
				}
			}
		}()

		io.Copy(writer, s)
	})

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
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

				// fmt.Printf(terminal.ClearScreen + terminal.ResetCursor)
				// fmt.Println(renderer.Render(snakes, pip, size))

				break
			}
		}
	}(ctx)

	fmt.Println("Serving connections on :2048")
	go ssh.ListenAndServe(":2048", nil, ssh.HostKeyFile("./keys/id_rsa"))

	<-done

	fmt.Println("SSH server shutting down...")
}
