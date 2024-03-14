package system

import (
	"fmt"
	"math"
	"time"

	"github.com/patrick22414/rewrite/render"
)

type SierpinskiTriangle struct {
	RenderOptions render.RenderOptions
}

func (SierpinskiTriangle) Axioms() []rune {
	return []rune{'F', '-', 'G', '-', 'G'}
}

func (SierpinskiTriangle) Rule(r rune) []rune {
	switch r {
	case 'F':
		return []rune{'F', '-', 'G', '+', 'F', '+', 'G', '-', 'F'}
	case 'G':
		return []rune{'G', 'G'}
	case '+', '-':
		return []rune{r}
	default:
		panic(fmt.Sprintf("unknown rune: %v", r))
	}
}

var SierpinskiTriangleRenderOptions = render.RenderOptions{
	InitX:         20,
	InitY:         700,
	InitDirection: math.Pi / 2,
	InitMagnitude: 15,
	TurnAngle:     -math.Pi * 2 / 3,
	Interval:      5 * time.Millisecond,
}
