package system

import "fmt"

type Algae struct{}

func (Algae) Axioms() []rune {
	return []rune{'A', 'B'}
}

func (Algae) Rule(r rune) []rune {
	switch r {
	case 'A':
		return []rune{'A', 'B'}
	case 'B':
		return []rune{'A'}
	default:
		panic(fmt.Sprintf("unknown rune: %v", r))
	}
}
