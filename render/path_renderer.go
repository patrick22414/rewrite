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
func (ren *DefaultPathRenderer) RenderStep(r rune, t Turtle, s TurtleStack, p vector.Path) (Turtle, TurtleStack, vector.Path) {
	switch r {
	case 'X':
		// do nothing
	case 'F', 'G':
		t = t.Move()
		p.LineTo(float32(t.x), float32(t.y))
	case '+':
		t = t.Turn(ren.TurnAngle)
	case '-':
		t = t.Turn(-ren.TurnAngle)
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

func (ren *DefaultPathRenderer) Render(runes []rune) vector.Path {
	t := Turtle{
		x: ren.InitX,
		y: ren.InitY,
		a: ren.InitDirection,
		d: ren.InitMagnitude,
	}
	s := make(TurtleStack, 0)
	var p vector.Path
	p.MoveTo(float32(t.x), float32(t.y))

	for _, r := range runes {
		t, s, p = ren.RenderStep(r, t, s, p)
	}

	return p
}

func (ren *DefaultPathRenderer) AsyncRender(runes []rune, cancel <-chan struct{}) <-chan vector.Path {
	out := make(chan vector.Path)

	go func() {
		defer close(out)

		t := Turtle{
			x: ren.InitX,
			y: ren.InitY,
			a: ren.InitDirection,
			d: ren.InitMagnitude,
		}
		s := make(TurtleStack, 0)
		var p vector.Path
		p.MoveTo(float32(t.x), float32(t.y))
		for _, r := range runes {
			select {
			case out <- p:
				t, s, p = ren.RenderStep(r, t, s, p)
			case <-cancel:
				return
			}
		}
	}()

	return out
}
