package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

var (
	width, height int
	cols          = width / res
	rows          = height / res
	res           = 20
	scale         = 2
	grid          [][]int
)

type Game struct {
	//generation int
	grid [][]int
}

var aqua = color.RGBA{R: 173, G: 216, B: 230}
var blue = color.RGBA{R: 173, G: 216, B: 230}

//Use this as 2D array
//func emptyGen(cols, rows int)  {
//	arr := make([][]int, cols)
//	for i := 0; i < len(arr); i++ {
//		arr[i] = make([]int, rows)
//	}
//
//	return
//}

//Drawing the game board
func Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			//x := i * res
			//y := j * res
			if grid[i][j] > 0 {
				for x := 0; x < scale; x++ {
					for y := 0; y < scale; y++ {
						screen.Set((i*scale)-1, (j*scale)-1, blue)

					}
				}
				//screen.Fill(color.RGBA{R: 173, G: 216, B: 230})
				//ebiten.NewImage(res,res,ebiten.FilterDefault)
			}
		}
	}

}

//Counting the neighbors
func countNeighbor(g *Game, x, y int) int {
	sum := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			//		cols = ( int(x + i) + cols) % cols
			//		rows = (int(y + j) + rows) % rows
			//		sum += grid[i-1][j-1]
			//		sum += grid[i][j-1]
			//		sum += grid[i+1][j-1]
			//		sum += grid[i+1][j]
			//		sum += grid[i+1][j+1]
			//		sum += grid[i][j+1]
			//		sum += grid[i-1][j+1]
			//		sum += grid[i-1][j]
			sum += g.grid[i][j]
		}
	}
	sum -= g.grid[x][y]
	return sum
}

var next = Draw

//Updating the game board
func (g *Game) Update(screen *ebiten.Image) error {
	//rand.Seed(time.Now().UnixNano())
	for x := 1; x < cols; x++ {
		for y := 1; y < rows; y++ {
			var state = grid[x][y]
			neighbors := countNeighbor()

			if state == 0 && neighbors == 3 {
				//need to make another grid
				grid[x][y] = 1
			} else if state == 1 && (neighbors < 2 || neighbors > 3) {
				//make another grid
				grid[x][y] = 0
			} else {
				//make another grid
				grid[x][y] = state
			}
		}
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 500
}

func main() {
	//e := emptyGen()
	//newState(e)
	ebiten.SetWindowSize(1040, 800)
	ebiten.SetWindowTitle("Conway's Game of Life")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
