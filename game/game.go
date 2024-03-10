package game

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/patrick22414/rewrite/system"
)

const (
	SW, SH = 640, 480
)

var (
	whiteImage = ebiten.NewImage(3, 3)

	// whiteSubImage is an internal sub image of whiteImage.
	// Use whiteSubImage at DrawTriangles instead of whiteImage in order to avoid bleeding edges.
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func DrawRunes(screen *ebiten.Image, runes []rune) {
	turtle := &Turtle{
		x: 100,
		y: 100,
		a: 0,
	}

	var path vector.Path
	path.MoveTo(float32(turtle.x), float32(turtle.y))
	for _, r := range runes {
		switch r {
		case 'F', 'G':
			x, y := turtle.Move(Forward)
			path.LineTo(float32(x), float32(y))
		case '+':
			turtle.Turn(Left)
		case '-':
			turtle.Turn(Right)
		default:
			panic(fmt.Sprintf("unknown rune to draw: %v", r))
		}
	}
	path.MoveTo(0, 0)
	path.Close()

	stokeOptions := &vector.StrokeOptions{
		Width:    3,
		LineCap:  vector.LineCapButt,
		LineJoin: vector.LineJoinMiter,
	}
	vertices, indices := path.AppendVerticesAndIndicesForStroke(nil, nil, stokeOptions)

	for i := range vertices {
		vertices[i].SrcX = 1
		vertices[i].SrcY = 1
		vertices[i].ColorA = 1
	}

	triOptions := &ebiten.DrawTrianglesOptions{}
	screen.DrawTriangles(vertices, indices, whiteSubImage, triOptions)
}

type Game struct {
	runes   []rune
	rewrite chan []rune
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		g.runes = <-g.rewrite
		fmt.Println(string(g.runes))
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	// fmt.Printf("g.runes: %v\n", g.runes)
	DrawRunes(screen, g.runes)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SW, SH
}

func Start() {
	ebiten.SetWindowSize(SW, SH)
	ebiten.SetWindowTitle("Rewrite — L-systems in Go")

	sys := system.DragonCurve{}
	rewrite := system.Rewrite(sys)

	game := &Game{
		runes:   <-rewrite,
		rewrite: rewrite,
	}

	whiteImage.Fill(color.White)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
