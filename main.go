package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 641
	screenHeight = 481
	tileSize     = 32
	rows         = screenHeight / tileSize
	cols         = screenWidth / tileSize
)

type Game struct {
	matrix [rows][cols]bool
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		gridX := mx / tileSize
		gridY := my / tileSize

		if gridX >= 0 && gridX < cols && gridY >= 0 && gridY < rows {
			g.matrix[gridY][gridX] = true
		}
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		mx, my := ebiten.CursorPosition()
		gridX := mx / tileSize
		gridY := my / tileSize

		if gridX >= 0 && gridX < cols && gridY >= 0 && gridY < rows {
			g.matrix[gridY][gridX] = false
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 20, 255})

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			clr := color.RGBA{255, 255, 255, 255}
			if g.matrix[y][x] {
				vector.FillRect(
					screen,
					float32(x*tileSize),
					float32(y*tileSize),
					float32(tileSize),
					float32(tileSize),
					clr,
					false,
				)
			} else {
				vector.StrokeRect(
					screen,
					float32(x*tileSize),
					float32(y*tileSize),
					float32(tileSize),
					float32(tileSize),
					1, // Liniendicke
					clr,
					false,
				)
			}

		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pathfinding Demo")

	game := &Game{
		matrix: [rows][cols]bool{},
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
