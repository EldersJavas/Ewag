package main

import (
	game "github.com/EldersJavas/Ewag/Game"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

func main() {
	g:=game.NewGame()
	g.Name="God"
	ebiten.SetMaxTPS(30)
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle(g.Name)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
