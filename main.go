package main

import (
	"log"
	"math"

	"github.com/ebiten/emoji"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type Game struct {
	ScreenWidth, ScreenHeight int
	R                         float64
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.R += 0.1
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	sushi := emoji.Image("🍣")
	geom := &ebiten.GeoM{}
	geom.Reset()
	geom.Translate(-float64(sushi.Bounds().Dx())/2, -float64(sushi.Bounds().Dy())/2)
	s := 1 + math.Cos(g.R)*0.2
	geom.Scale(s, s)
	geom.Rotate(g.R)
	geom.Translate(float64(g.ScreenWidth)/2, float64(g.ScreenHeight)/2)
	screen.DrawImage(sushi, &ebiten.DrawImageOptions{
		GeoM: *geom,
	})
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenHeight
}

func main() {
	game := &Game{
		ScreenWidth:  320,
		ScreenHeight: 240,
		R:            0,
	}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
