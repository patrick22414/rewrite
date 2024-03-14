package system

import (
	"fmt"
	"math"
	"time"

	"github.com/patrick22414/rewrite/render"
)

type DragonCurve struct{}

func (DragonCurve) Axioms() []rune {
	return []rune{'F'}
}

func (DragonCurve) Rule(r rune) []rune {
	switch r {
	case 'F':
		return []rune{'F', '+', 'G'}
	case 'G':
		return []rune{'F', '-', 'G'}
	case '+', '-':
		return []rune{r}
	default:
		panic(fmt.Sprintf("unknown rune: %v", r))
	}
}

var DragonCurveRenderOptions = render.RenderOptions{
	InitX:         800,
	InitY:         480,
	InitDirection: math.Pi / 2,
	InitMagnitude: 15,
	TurnAngle:     math.Pi / 2,
	Interval:      10 * time.Millisecond,
}
