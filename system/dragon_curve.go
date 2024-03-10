package system

import "fmt"

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
