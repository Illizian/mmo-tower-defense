package renderer

import (
	"mmo-tower-defense/pkg/entities"
	"mmo-tower-defense/pkg/maths"
	"mmo-tower-defense/pkg/terminal"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleRender(t *testing.T) {
	assert.True(t, true, "True is true!")

	size := 3
	pip := maths.NewVec2(0, 0)
	snakes := []*entities.Snake{
		{
			Color:       terminal.Green,
			Location:    maths.NewVec2(1, 1),
			Direction:   maths.East,
			Path:        make([]maths.Vec2, 0),
			Status:      entities.SNAKE_ALIVE,
			DeadCounter: 0,
		},
	}

	// ┏━━━┓
	// ┃•  ┃
	// ┃ ► ┃
	// ┃   ┃
	// ┗━━━┛
	actual := Render(snakes, pip, size)
	expected := []string{
		BORDER_TL, BORDER_HL, BORDER_HL, BORDER_HL, BORDER_TR, "\n\r",
		BORDER_VL, terminal.Yellow, "•", terminal.Reset, " ", " ", BORDER_VL, "\n\r",
		BORDER_VL, " ", terminal.Green, "►", terminal.Reset, " ", BORDER_VL, "\n\r",
		BORDER_VL, " ", " ", " ", BORDER_VL, "\n\r",
		BORDER_BL, BORDER_HL, BORDER_HL, BORDER_HL, BORDER_BR, "\n\r",
	}

	assert.Equal(t, strings.Join(expected, ""), actual)
}
