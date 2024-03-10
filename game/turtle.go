package game

import (
	"math"
)

type Turtle struct {
	x, y float64
	a    float64 // radian angle
}

func (t *Turtle) Move(distance float64) (newX, newY float64) {
	sin, cos := math.Sincos(t.a)
	t.x += distance * sin
	t.y += distance * cos
	return t.x, t.y
}

func (t *Turtle) Turn(angle float64) {
	t.a += angle
}

const Forward = 16
const Left = +math.Pi / 2
const Right = -math.Pi / 2
