package system

type System interface {
	Axioms() []rune
	Rule(r rune) []rune
}

func Rewrite(sys System) <-chan []rune {
	out := make(chan []rune)

	go func() {
		runes := sys.Axioms()

		out <- runes

		for {
			newRunes := make([]rune, 0, cap(runes)*2)
			for _, r := range runes {
				newRunes = append(newRunes, sys.Rule(r)...)
			}
			out <- newRunes
			runes = newRunes
		}
	}()

	return out
}
