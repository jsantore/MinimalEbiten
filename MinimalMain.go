package main

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/png"
	"log"
)

//go:embed assets/*
var EmbeddedAssets embed.FS

const (
	GameWidth   = 1000
	GameHeight  = 540
	PlayerSpeed = 2
)

type Sprite struct {
	pict *ebiten.Image
	xloc int
	yloc int
	dX   int
	dY   int
}

type Game struct {
	player  Sprite
	score   int
	drawOps ebiten.DrawImageOptions
}

func (g *Game) Update() error {
	processPlayerInput(g)
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	g.drawOps.GeoM.Reset()
	g.drawOps.GeoM.Translate(float64(g.player.xloc), float64(g.player.yloc))
	screen.DrawImage(g.player.pict, &g.drawOps)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GameWidth, GameHeight
}

func main() {
	ebiten.SetWindowSize(GameWidth, GameHeight)
	ebiten.SetWindowTitle("Minimal Game")
	simpleGame := Game{score: 0}
	simpleGame.player = Sprite{
		pict: loadPNGImageFromEmbedded("f1-ship1-3.png"),
		xloc: 200,
		yloc: 300,
		dX:   0,
		dY:   0,
	}
	if err := ebiten.RunGame(&simpleGame); err != nil {
		log.Fatal("Oh no! something terrible happened and the game crashed", err)
	}
}

func loadPNGImageFromEmbedded(name string) *ebiten.Image {
	pictNames, err := EmbeddedAssets.ReadDir("assets")
	if err != nil {
		log.Fatal("failed to read embedded dir ", pictNames, " ", err)
	}
	embeddedFile, err := EmbeddedAssets.Open("assets/" + name)
	if err != nil {
		log.Fatal("failed to load embedded image ", embeddedFile, err)
	}
	rawImage, err := png.Decode(embeddedFile)
	if err != nil {
		log.Fatal("failed to load embedded image ", name, err)
	}
	gameImage := ebiten.NewImageFromImage(rawImage)
	return gameImage
}

func processPlayerInput(theGame *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		theGame.player.dY = -PlayerSpeed
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		theGame.player.dY = PlayerSpeed
	} else if inpututil.IsKeyJustReleased(ebiten.KeyUp) || inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		theGame.player.dY = 0
	}
	theGame.player.yloc += theGame.player.dY
	if theGame.player.yloc <= 0 {
		theGame.player.dY = 0
		theGame.player.yloc = 0
	} else if theGame.player.yloc+theGame.player.pict.Bounds().Size().Y > GameHeight {
		theGame.player.dY = 0
		theGame.player.yloc = GameHeight - theGame.player.pict.Bounds().Size().Y
	}
}
