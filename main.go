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
)

type Game struct {
	generation int
	grid       [][]int
}

//func (g *Game) Error() string {
//	panic("implement me")
//}

var grid [40][50]int
var aqua = color.RGBA{R: 173, G: 216, B: 230}
var blue = color.RGBA{R: 173, G: 216, B: 230}

//Use this as 2D array
func emptyGen() *Game {
	arr := make([][]int, cols)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, rows)
	}

	return &Game{grid: arr, generation: 1}
}

//func setup()  {
//	//grid := emptyGen(cols, rows)
//	for i:=0;i <cols; i++ {
//		for j:= 0; j <rows; j++ {
//			grid[i][j] = rand.Intn(2)
//		}
//	}
//}

//func newState(g *Game)  {
//	rand.Seed(time.Now().UnixNano())
//	for x := 0; x < res; x++ {
//		for y := 0; y < res; y++ {
//			if rand.Intn(15) == 1 {
//				g.grid[x][y] = 1
//			}
//		}
//	}
//}

func (g *Game) Update(screen *ebiten.Image) error {
	//rand.Seed(time.Now().UnixNano())
	for x := 1; x < cols; x++ {
		for y := 1; y < rows; y++ {
			state := grid[x][y]
			neighbors := countNeighbor(x, y, g.grid)

			if state == 0 && neighbors == 3 {
				//need to make another grid
			} else if state == 1 && (neighbors < 2 || neighbors > 3) {
				//make another grid
			} else {
				//make another grid
			}

			//sum +=grid[x-1][y-1]
			//sum +=grid[x][y-1]
			//sum +=grid[x+1][y-1]
			//sum +=grid[x+1][y]
			//sum +=grid[x+1][y+1]
			//sum +=grid[x-1][y]
			//sum +=grid[x][y+1]
			//sum +=grid[x-1][y]
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
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
