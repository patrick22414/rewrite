package system

import (
	"fmt"
	"math"
	"time"

	"github.com/patrick22414/rewrite/render"
)

type BinaryTree struct{}

func (BinaryTree) Axioms() []rune {
	return []rune{'G'}
}

func (BinaryTree) Rule(r rune) []rune {
	switch r {
	case 'F':
		return []rune{'F', 'F'}
	case 'G':
		return []rune{'F', '[', '+', 'G', ']', '-', 'G'}
	case '+', '-', '[', ']':
		return []rune{r}
	default:
		panic(fmt.Sprintf("unknown rune: %v", r))
	}
}

var BinaryTreeRenderOptions = render.RenderOptions{
	InitX:         640,
	InitY:         700,
	InitDirection: math.Pi,
	InitMagnitude: 4,
	TurnAngle:     math.Pi / 4,
	Interval:      8 * time.Millisecond,
}
