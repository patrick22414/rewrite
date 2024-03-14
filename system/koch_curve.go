package system

import (
	"fmt"
	"math"
	"time"

	"github.com/patrick22414/rewrite/render"
)

type KochCurve struct{}

func (KochCurve) Axioms() []rune {
	return []rune{'F'}
}

func (KochCurve) Rule(r rune) []rune {
	switch r {
	case 'F':
		return []rune{'F', '+', 'F', '-', 'F', '-', 'F', '+', 'F'}
	case '+', '-':
		return []rune{r}
	default:
		panic(fmt.Sprintf("unknown rune: %v", r))
	}
}

var KochCurveRenderOptions = render.RenderOptions{
	InitX:         20,
	InitY:         700,
	InitDirection: math.Pi / 2,
	InitMagnitude: 15,
	TurnAngle:     math.Pi / 2,
	Interval:      10 * time.Millisecond,
}
