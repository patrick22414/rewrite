package system

import "fmt"

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
