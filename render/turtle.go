package render

import (
	"math"
)

type Turtle struct {
	x, y float64
	a    float64 // radian direction
	d    float64 // magnitude
}

func (t Turtle) Move() Turtle {
	sin, cos := math.Sincos(t.a)
	t.x += t.d * sin
	t.y += t.d * cos
	return t
}

func (t Turtle) Turn(angle float64) Turtle {
	t.a += angle
	return t
}

func (t Turtle) Scale(scale float64) Turtle {
	t.d *= scale
	return t
}

type TurtleStack []Turtle

func (s *TurtleStack) Push(v Turtle) {
	*s = append(*s, v)
}

func (s *TurtleStack) Pop() Turtle {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
