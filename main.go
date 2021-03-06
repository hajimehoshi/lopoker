package main

import (
	"Loowootoo/lopoker/game"
	"Loowootoo/lopoker/ui2d"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const winWidth, winHeight int = 800, 600

var newGame *game.Game
var ui *ui2d.UI2d

func init() {
	rand.Seed(time.Now().UnixNano())
}

func update(screen *ebiten.Image) error {
	newGame.GameLoop()
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	ui.Draw(screen, newGame)
	return nil
}

func main() {
	newGame = game.NewGame(0)
	newGame.Shuffle()
	ui = ui2d.NewUI2d()
	err := ebiten.Run(update, winWidth, winHeight, 1, "LoPoker !!!")
	if err != nil {
		panic(err)
	}
}
