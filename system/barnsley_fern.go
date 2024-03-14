package system

import (
	"fmt"
	"math"
	"time"

	"github.com/patrick22414/rewrite/render"
)

type BarnsleyFern struct{}

func (BarnsleyFern) Axioms() []rune {
	return []rune{'X'}
}

func (BarnsleyFern) Rule(r rune) []rune {
	switch r {
	case 'X':
		return []rune("F+[[X]-X]-F[-FX]+X")
	case 'F':
		return []rune("FF")
	case '+', '-', '[', ']':
		return []rune{r}
	default:
		panic(fmt.Sprintf("unknown rune: %v", r))
	}
}

var BarnsleyFernRenderOptions = render.RenderOptions{
	InitX:         60,
	InitY:         700,
	InitDirection: (180 - 25) * math.Pi / 180,
	InitMagnitude: 4,
	TurnAngle:     25 * math.Pi / 180,
	Interval:      2 * time.Millisecond,
}
