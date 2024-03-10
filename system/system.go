package system

type System interface {
	Axioms() []rune
	Rule(r rune) []rune
}

func Rewrite(sys System) (output chan []rune) {
	output = make(chan []rune, 1)

	go func() {
		runes := sys.Axioms()

		output <- runes

		for {
			newRunes := make([]rune, 0, cap(runes)*2)
			for _, r := range runes {
				newRunes = append(newRunes, sys.Rule(r)...)
			}
			output <- newRunes
			runes = newRunes
		}
	}()

	return
}
