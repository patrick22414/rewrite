# Rewrite

[L-system](https://en.wikipedia.org/wiki/L-system) implemented in Go with [Ebitengine](https://github.com/hajimehoshi/ebiten).

## Demo

### Barnsley Fern
https://github.com/patrick22414/rewrite/assets/32917395/e20cef6a-da43-4068-8189-9f3157a20d80

### Binary Tree
https://github.com/patrick22414/rewrite/assets/32917395/ea0245ea-b9ae-4a01-bd02-308c36333a8c

### Dragon Curve
https://github.com/patrick22414/rewrite/assets/32917395/87489bc7-0081-4c14-bd14-d0e1f71d7b2e

### Koch Curve
https://github.com/patrick22414/rewrite/assets/32917395/d40818ba-f136-4fcc-aedd-64889818f453

### Sierpinski Triangle
https://github.com/patrick22414/rewrite/assets/32917395/9f844997-2665-401f-b805-8ef2576492fc

## Code

- `system/` L-system interface and rewriting logic (see below), and a few examples.
- `render/` Renderer inferface with [turtle graphics](https://en.wikipedia.org/wiki/Turtle_graphics).
- `game/` Game engine logic.
- `main.go` Entry point. Select examples here.

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
