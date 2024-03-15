package game

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/patrick22414/rewrite/render"
	"github.com/patrick22414/rewrite/system"
)

const (
	SW, SH = 1280, 720
)

var (
	whiteImage = ebiten.NewImage(3, 3)

	// whiteSubImage is an internal sub image of whiteImage.
	// Use whiteSubImage at DrawTriangles instead of whiteImage in order to avoid bleeding edges.
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

var aa = false

type Game struct {
	i                 int // iteration count
	runes             []rune
	rewrite           <-chan []rune
	renderer          render.PathRenderer
	path              vector.Path
	cancelAsyncRender chan struct{}
}

func (g *Game) StartAsyncRender() {
	// fmt.Printf("Started async render of runes len %d\n", len(g.runes))
	if g.cancelAsyncRender != nil {
		close(g.cancelAsyncRender) // cancel previous render
	}

	g.cancelAsyncRender = make(chan struct{})

	var ticker *time.Ticker
	switch ren := g.renderer.(type) {
	case (*render.DefaultPathRenderer):
		ticker = time.NewTicker(ren.Interval)
	default:
		ticker = time.NewTicker(10 * time.Millisecond)
	}
	defer ticker.Stop()

	for nextPath := range g.renderer.AsyncRender(g.runes, g.cancelAsyncRender) {
		select {
		case <-ticker.C:
			g.path = nextPath
		case <-g.cancelAsyncRender:
			return
		}
	}

	g.cancelAsyncRender = nil // end of render, no longer cancellable
	// fmt.Printf("End of async render of runes len %d\n", len(g.runes))
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		if g.cancelAsyncRender != nil {
			close(g.cancelAsyncRender) // cancel async render
			g.cancelAsyncRender = nil
			g.path = g.renderer.Render(g.runes)
			return nil
		}

		g.i++
		g.runes = <-g.rewrite

		// Sync render.
		// g.path = g.renderer.Render(g.runes)

		// Async render.
		go g.StartAsyncRender()
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyR) {
		go g.StartAsyncRender()
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyA) {
		aa = !aa
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("Game Ended")
	}

	// fmt.Println(string(g.runes))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	stokeOptions := &vector.StrokeOptions{
		Width:    1,
		LineCap:  vector.LineCapButt,
		LineJoin: vector.LineJoinMiter,
	}
	vertices, indices := g.path.AppendVerticesAndIndicesForStroke(nil, nil, stokeOptions)

	for i := range vertices {
		vertices[i].SrcX = 1
		vertices[i].SrcY = 1
	}

	triOptions := &ebiten.DrawTrianglesOptions{
		AntiAlias: aa,
	}
	screen.DrawTriangles(vertices, indices, whiteSubImage, triOptions)

	msg := fmt.Sprintf("TPS: %0.2f | FPS: %0.2f | goroutines: %d", ebiten.ActualTPS(), ebiten.ActualFPS(), runtime.NumGoroutine())
	msg += fmt.Sprintf("\nIteration: %d", g.i)
	msg += "\n [Space] Rewrite the system and play growth animation"
	msg += "\n     [R] Replay the current system growth animation"
	msg += "\n     [A] Toggle anti-alias"
	msg += "\n   [Esc] End game"
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SW, SH
}

func Start(sys system.System, rend render.PathRenderer) {
	ebiten.SetWindowSize(SW, SH)
	ebiten.SetWindowTitle("Rewrite — L-systems in Go")

	rewrite := system.Rewrite(sys)
	runes := <-rewrite
	path := rend.Render(runes)

	game := &Game{
		runes:    runes,
		rewrite:  rewrite,
		renderer: rend,
		path:     path,
	}

	whiteImage.Fill(color.White)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
