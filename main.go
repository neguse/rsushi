package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/ebiten/emoji"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const Delta = 1.0 / 60.0

type Sushi struct {
	X, Y     float64 // 0 ~ 1
	TimeInit float64
	TimeLeft float64
	R        float64
	DR       float64
	Scale    float64
}

func (s *Sushi) Update() {
	s.R += Delta * s.DR
	s.TimeLeft -= Delta
}

func (s *Sushi) Dead() bool {
	return s.TimeLeft < 0
}

func (g *Game) DrawSushi(screen *ebiten.Image, s *Sushi) {
	geom := &ebiten.GeoM{}
	geom.Reset()
	geom.Translate(-float64(g.SushiImage.Bounds().Dx())/2, -float64(g.SushiImage.Bounds().Dy())/2)
	geom.Scale(1/float64(g.SushiImage.Bounds().Dx()), 1/float64(g.SushiImage.Bounds().Dy()))
	scale := math.Sin(math.Pi*((s.TimeInit-s.TimeLeft)/s.TimeInit)) * s.Scale
	geom.Scale(scale, scale)
	geom.Rotate(s.R)
	geom.Translate(float64(g.ScreenWidth)*s.X, float64(g.ScreenHeight)*s.Y)
	screen.DrawImage(g.SushiImage, &ebiten.DrawImageOptions{
		GeoM: *geom,
	})
}

// Game implements ebiten.Game interface.
type Game struct {
	ScreenWidth, ScreenHeight int
	Sushis                    []*Sushi
	Tick                      int
	SushiImage                *ebiten.Image
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.Tick++
	sushis := []*Sushi{}
	for _, s := range g.Sushis {
		s.Update()
		if !s.Dead() {
			sushis = append(sushis, s)
		}
	}
	spawn := false
	if ebiten.CurrentFPS() > 60 {
		spawn = true
	} else if ebiten.CurrentFPS() > 30 && g.Tick%2 == 0 {
		spawn = true
	} else if ebiten.CurrentFPS() > 15 && g.Tick%4 == 0 {
		spawn = true
	}
	if spawn {
		t := rand.Float64()*4.0 + 2
		sushis = append(sushis, &Sushi{
			X:        rand.Float64(),
			Y:        rand.Float64(),
			TimeInit: t,
			TimeLeft: t,
			R:        rand.Float64() * math.Pi * 2,
			DR:       rand.Float64() * math.Pi * 2,
			Scale:    rand.Float64()*120 + 80,
		})
	}
	g.Sushis = sushis
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	for _, s := range g.Sushis {
		g.DrawSushi(screen, s)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Sushis:%3d FPS:%3.2f", len(g.Sushis), ebiten.CurrentFPS()))
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenHeight
}

func main() {
	game := &Game{
		ScreenWidth:  640,
		ScreenHeight: 480,
		SushiImage:   emoji.Image("🍣"),
	}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
