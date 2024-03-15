package render

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2/vector"
)

type PathRenderer interface {
	Render(runes []rune) vector.Path
	AsyncRender(runes []rune, cancel <-chan struct{}) <-chan vector.Path
}

type RenderOptions struct {
	InitX, InitY  float64
	InitDirection float64
	InitMagnitude float64
	TurnAngle     float64
	Interval      time.Duration // async render interval
}

type DefaultPathRenderer struct {
	RenderOptions
}

// Render a single rune.
func (rend *DefaultPathRenderer) RenderStep(r rune, t Turtle, s TurtleStack, p vector.Path) (Turtle, TurtleStack, vector.Path) {
	switch r {
	case 'X':
		// do nothing
	case 'F', 'G':
		t = t.Move()
		p.LineTo(float32(t.x), float32(t.y))
	case '+':
		t = t.Turn(rend.TurnAngle)
	case '-':
		t = t.Turn(-rend.TurnAngle)
	case '[':
		s.Push(t)
	case ']':
		t = s.Pop()
		p.MoveTo(float32(t.x), float32(t.y))
	default:
		panic(fmt.Sprintf("unknown rune to render: %v", r))
	}
	return t, s, p
}

func (rend *DefaultPathRenderer) Render(runes []rune) vector.Path {
	t := Turtle{
		x: rend.InitX,
		y: rend.InitY,
		a: rend.InitDirection,
		d: rend.InitMagnitude,
	}
	s := make(TurtleStack, 0)
	var p vector.Path
	p.MoveTo(float32(t.x), float32(t.y))

	for _, r := range runes {
		t, s, p = rend.RenderStep(r, t, s, p)
	}

	return p
}

func (rend *DefaultPathRenderer) AsyncRender(runes []rune, cancel <-chan struct{}) <-chan vector.Path {
	out := make(chan vector.Path)

	go func() {
		defer close(out)

		t := Turtle{
			x: rend.InitX,
			y: rend.InitY,
			a: rend.InitDirection,
			d: rend.InitMagnitude,
		}
		s := make(TurtleStack, 0)
		var p vector.Path
		p.MoveTo(float32(t.x), float32(t.y))
		for _, r := range runes {
			select {
			case out <- p:
				t, s, p = rend.RenderStep(r, t, s, p)
			case <-cancel:
				return
			}
		}
	}()

	return out
}
