package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"strconv"
)

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, strconv.FormatFloat(ebiten.CurrentTPS(), 'f', 2, 64))
}