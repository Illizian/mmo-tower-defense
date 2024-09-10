package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"mmo-tower-defense/pkg/entities"
	"mmo-tower-defense/pkg/maths"
	"mmo-tower-defense/pkg/renderer"
	"mmo-tower-defense/pkg/terminal"

	"time"

	"github.com/gliderlabs/ssh"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tick := time.NewTicker(200 * time.Millisecond)
	size := 20
	pip := maths.NewRandomVec2(0, size-1)

	var snakes []*entities.Snake

	ssh.Handle(func(s ssh.Session) {
		addr := s.RemoteAddr()
		fmt.Printf("[%s] Client Connected\n", addr)
		closed := make(chan bool, 1)

		s.Write([]byte(terminal.CursorHide))
		snek := entities.Snake{
			Color:       terminal.Green,
			Location:    maths.NewRandomVec2(0, size-1),
			Direction:   maths.East,
			Length:      3,
			Path:        make([]maths.Vec2, 0),
			Status:      entities.SNAKE_ALIVE,
			DeadCounter: 0,
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
					if snek.Status == entities.SNAKE_ALIVE {
						snek.Status = entities.SNAKE_DIEING
					}

					s.Write([]byte(terminal.CursorShow))
					s.Exit(0)
					closed <- true
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
				case <-closed:
					return
				case <-ctx.Done():
					s.Write([]byte(fmt.Sprintf("%s%sServer shutting down... Goodbye!", terminal.ClearScreen, terminal.ResetCursor)))
					s.Exit(0)
					return
				case <-framer.C:
					s.Write([]byte(fmt.Sprintf(terminal.ClearScreen + terminal.ResetCursor)))

					if snek.Status != entities.SNAKE_ALIVE {
						s.Write([]byte(fmt.Sprintf("You DED - Your score: %d", snek.Length-3)))
					}

					if snek.Status == entities.SNAKE_ALIVE || snek.Status == entities.SNAKE_DEAD {
						// @TODO: Should probably only render once, and share with everyone! Can we just pipe in a channel? From the renderer?
						s.Write([]byte(renderer.Render(snakes, pip, size)))
					}

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
					if snake.Status == entities.SNAKE_DEAD {
						continue
					}

					occupied[snake.Location] = true
					for _, tail := range snake.Path {
						occupied[tail] = true
					}
				}

				// Tick each snake with the generated occupied for collisions
				for s := range snakes {
					if snakes[s].Status == entities.SNAKE_DIEING {
						snakes[s].DeadCounter++
						if snakes[s].DeadCounter == 10 {
							snakes[s].Status = entities.SNAKE_DEAD
						}
						continue
					}

					if snakes[s].Status == entities.SNAKE_DEAD {
						continue
					}

					location := snakes[s].Tick(occupied, size)
					if location.Eq(pip) {
						snakes[s].Length += 1
						pip = maths.NewRandomVec2(0, size-1)
					}
				}

				break
			}
		}
	}(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Serving connections on :2048")
	go ssh.ListenAndServe(":2048", nil, ssh.HostKeyFile("./keys/id_rsa"))

	<-c

	fmt.Println("SSH server shutting down...")
}
