package game

import (
	"Loowootoo/lopoker/genlib"
	"Loowootoo/lopoker/ui2d/assets/pcard"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	Player        *Player
	CardSet       *Cards
	SmokeAnim     *genlib.Sprite
	Message       string
	GameStatus    MainGameState
	GameSubStatus int
	MsgCounter    *genlib.TimeCounter
	GameWin       int
	WinStr        string
}

func NewGame(credit int) *Game {
	player := NewPlayer(credit)
	cardSet := NewCardSet()
	smokeAnim := genlib.NewSprite()
	smokeAnim.AddAnimFrameFromBytes("default", pcard.SmokePNG, 300, 15, ebiten.FilterDefault)
	smokeAnim.CenterCoordonnates = true
	smokeAnim.Pos = genlib.Vector{576, 153, 100}
	smokeAnim.Speed = 1
	smokeAnim.Direction = genlib.Vector{0, 0, -1}
	smokeAnim.Start()
	msgCounter := genlib.NewCounter(500)
	return &Game{player, cardSet, smokeAnim, "", GameDEMO, 0, msgCounter, 0, ""}
}

func (g *Game) AddBet(bet int) {
	g.Player.Bet += bet
}

func (g *Game) Shuffle() {
	g.CardSet.DeckCard()
}

func (g *Game) Deal() {
	g.Player.ResetHeld()
	for i := 0; i < len(g.Player.Hand); i++ {
		g.Player.Hand[i] = g.CardSet.GetNextCard()
	}
	g.Player.sortHand()
	g.Player.ResetBackCard()
}

func (g *Game) DealWithHeld() {
	for i := 0; i < 5; i++ {
		if g.Player.Held[i] == false {
			g.Player.Hand[i] = g.CardSet.GetNextCard()
			g.Player.BackCard[i] = true
		}
	}
	g.Player.sortHand()
}

func (g *Game) ShowPlayerCard() {
	g.Player.ShowPlayerCard()
	g.Player.ShowPlayerSortCard()
}

var testCount float64

func (g *Game) Run() {
	testCount++
	if testCount > 60 {
		g.Deal()
		g.Message = g.Player.CheckWin()
		testCount = 0
	}
}

type MainGameState int

const (
	GameDEMO MainGameState = iota
	GameSTART
	GameBET
	GamePLAY
	GameCHECK
	GameWIN
	GameLOSE
	GameACCOUNT
)

var GameMessage = [...]string{
	"GAME DEMO",
	"GAME START",
	"GAME BET",
	"GAME PLAY",
	"GAME CHECK",
	"GAME WIN",
	"GAME LOSE",
	"GAME ACCOUNT",
}

func (g *Game) GameLoop() {
	if g.IsCreditKey() {
		g.Player.Credit += 100
	}
	switch g.GameStatus {
	case GameDEMO:
		g.DemoProc()
	case GameSTART:
		g.GameStartProc()
	case GameBET:
		g.GameBetProc()
	case GamePLAY:
		g.GamePlayProc()
	case GameCHECK:
		g.GameCheckProc()
	case GameWIN:
		g.GameWinProc()
	case GameLOSE:
		g.GameLoseProc()
	case GameACCOUNT:
		g.GameAccountProc()
	}
}
