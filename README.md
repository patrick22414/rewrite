# Rewrite

[L-system](https://en.wikipedia.org/wiki/L-system) implemented in Go with [Ebitengine](https://github.com/hajimehoshi/ebiten).

## Example

## Code

- `system/` L-system interface and rewriting logic (see below), and a few examples.
- `game/` Render L-system strings with [turtle graphics](https://en.wikipedia.org/wiki/Turtle_graphics)

```go
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
```
