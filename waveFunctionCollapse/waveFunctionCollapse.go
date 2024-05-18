package main

/* program inspiration:
 * https://www.youtube.com/watch?v=rI_y2GAlQFM&list=PLB3OFCROxZ41eaR2Q4Ls27WjnzVoDLT6D
 */

import (
	"embed"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CellType int

type Cell struct {
	collapsed bool
	possible  []CellType
}

var allTypes = []CellType{BLANK,
	LEFT,
	RIGHT,
	UP,
	DOWN,
}

type WFCData [][]Cell

const (
	NOTHING CellType = iota
	BLANK
	LEFT
	RIGHT
	UP
	DOWN
)

var possibilies = map[CellType][]CellType{
	NOTHING: allTypes,
	BLANK:   allTypes,
	LEFT: {
		BLANK,
		RIGHT,
		UP,
		DOWN,
	},
	RIGHT: {
		BLANK,
		LEFT,
		UP,
		DOWN,
	},
	UP: {
		BLANK,
		LEFT,
		RIGHT,
		DOWN,
	},
	DOWN: {
		BLANK,
		LEFT,
		RIGHT,
		UP,
	},
}

//go:embed all:tiles/demo
var tiles embed.FS

func loadTexture(s string) rl.Texture2D {
	bytes, _ := tiles.ReadFile(s)

	img := rl.LoadImageFromMemory(".png", bytes, int32(len(bytes)))
	texture := rl.LoadTextureFromImage(img)

	return texture

}

func resetWFCData(rows, cols int32) WFCData {
	var WFCData WFCData
	for r := int32(0); r < rows; r++ {
		var row []Cell

		for c := int32(0); c < cols; c++ {
			row = append(row, Cell{collapsed: false, possible: allTypes})
		}
		WFCData = append(WFCData, row)
	}

	return WFCData
}

func setCellType(WFCData WFCData, row, col int, cellType CellType) WFCData {
	if WFCData[row][col].collapsed {
		return WFCData
	}

	WFCData[row][col] = Cell{collapsed: true, possible: []CellType{cellType}}

	rows := len(WFCData) - 1
	cols := len(WFCData[0]) - 1
	// DOWN
	if row < rows {
		if !WFCData[row+1][col].collapsed {

			WFCData[row+1][col] = Cell{collapsed: false, possible: possibilies[cellType]}
		}
	}
	// UP
	if row >= rows {
		if !WFCData[row-1][col].collapsed {

			WFCData[row-1][col] = Cell{collapsed: false, possible: possibilies[cellType]}
		}
	}
	// RIGHT
	if col < cols {
		if !WFCData[row][col+1].collapsed {

			WFCData[row][col+1] = Cell{collapsed: false, possible: possibilies[cellType]}
		}
	}
	// LEFT
	if col >= cols {
		if !WFCData[row][col-1].collapsed {

			WFCData[row][col-1] = Cell{collapsed: false, possible: possibilies[cellType]}
		}
	}
	return WFCData
}

func newRandomImg(rows, cols int32) WFCData {
	var WFCData WFCData = resetWFCData(rows, cols)

	// random pos
	rr := rand.IntN(int(rows))
	rc := rand.IntN(int(cols))

	// making sure we dont draw a NOTHING
	WFCData = setCellType(WFCData, rr, rc, CellType(rand.IntN(4)+1))

	for r := 0; r < int(rows); r++ {
		for c := 0; c < int(cols); c++ {
			cell := &WFCData[r][c]
			if !cell.collapsed {
				WFCData = setCellType(WFCData, r, c, cell.possible[rand.IntN(len(cell.possible))])
			}
		}
	}

	return WFCData
}

const scale float32 = 0.25

func main() {

	rows := (100)
	cols := (100)

	width := (50)
	height := (50)

	var WFCData WFCData

	rl.SetConfigFlags(rl.FlagWindowUndecorated)
	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(int32(float32(rows*height)*scale), int32(float32(cols*width)*scale), "Wave Function Collapse")
	closeWindow := false

	tiles := map[CellType]rl.Texture2D{
		BLANK: loadTexture("tiles/demo/blank.png"),
		LEFT:  loadTexture("tiles/demo/left.png"),
		RIGHT: loadTexture("tiles/demo/right.png"),
		UP:    loadTexture("tiles/demo/up.png"),
		DOWN:  loadTexture("tiles/demo/left.png"),
	}

	actions := map[uint]interface{}{
		rl.KeyQ: func() {
			closeWindow = true
		},
		rl.KeyR: func() {
			WFCData = newRandomImg(int32(rows), int32(cols))
		},
	}
	actions[rl.KeyR].(func())()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() && !closeWindow {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if action := actions[uint(rl.GetKeyPressed())]; action != nil {
			action.(func())()
		}

		for r := (0); r < (len(WFCData)); r++ {
			for c := (0); c < (len(WFCData[r])); c++ {

				x := float32(c*height) * scale
				y := float32(r*width) * scale

				if WFCData[r][c].collapsed {
					possible := &WFCData[r][c].possible
					r := rand.IntN(len(*possible))

					rl.DrawTextureEx(tiles[(*possible)[r]], rl.NewVector2(x, y), 0.0, scale, rl.White)
				}

			}
		}

		rl.EndDrawing()

	}

	rl.CloseWindow()
}
