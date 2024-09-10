package renderer

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/term"
)

func StdIn(ctx context.Context, input chan []byte) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	rune := make([]byte, 1)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, err = os.Stdin.Read(rune)
			if err != nil {
				fmt.Println(err)
				return
			}

			input <- rune
		}
	}
}
