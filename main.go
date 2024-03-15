package main

import (
	"github.com/patrick22414/rewrite/game"
	"github.com/patrick22414/rewrite/render"
	"github.com/patrick22414/rewrite/system"
)

func main() {
	game.Start(
		system.BarnsleyFern{},
		&render.DefaultPathRenderer{
			RenderOptions: system.BarnsleyFernRenderOptions,
		},
	)
}
