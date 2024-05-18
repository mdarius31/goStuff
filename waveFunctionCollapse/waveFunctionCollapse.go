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

type WFCData [][]Cell

const (
	NOTHING CellType = iota
	BLANK
	LEFT
	RIGHT
	UP
	DOWN
)

var allTypes = []CellType{BLANK,
	LEFT,
	RIGHT,
	UP,
	DOWN,
}

type Direction int

const (
	DUP Direction = iota
	DDWON
	DLEFT
	DRIGHT
)

var allTypesAllDirs = map[Direction][]CellType{
	DUP:    allTypes,
	DDWON:  allTypes,
	DLEFT:  allTypes,
	DRIGHT: allTypes,
}

// var blankAllDirs = map[Direction][]CellType{
// 	DUP:    {BLANK},
// 	DDWON:  {BLANK},
// 	DLEFT:  {BLANK},
// 	DRIGHT: {BLANK},
// }

var possibilies = map[CellType]map[Direction][]CellType{
	NOTHING: allTypesAllDirs,
	BLANK:   allTypesAllDirs,
	UP: {
		DUP: {
			LEFT,
			RIGHT,
			DOWN,
		},
		DDWON: allTypes,
		DLEFT: {
			DOWN,
			RIGHT,
			UP,
		},
		DRIGHT: {
			LEFT,
			DOWN,
			UP,
		},
	},
	DOWN: {
		DUP: allTypes,
		DDWON: {
			LEFT,
			RIGHT,
			UP,
		},
		DLEFT: {
			DOWN,
			RIGHT,
			UP,
		},
		DRIGHT: {
			LEFT,
			DOWN,
			UP,
		},
	},
	LEFT: {
		DUP: {
			LEFT,
			RIGHT,
			DOWN,
		},
		DDWON: {
			LEFT,
			RIGHT,
			UP,
		},
		DLEFT: {
			DOWN,
			RIGHT,
			UP,
		},
		DRIGHT: allTypes,
	},
	RIGHT: {
		DUP: {
			DOWN,
			LEFT,
			RIGHT,
		},
		DDWON: {
			LEFT,
			RIGHT,
			UP,
		},
		DLEFT: allTypes,
		DRIGHT: {
			LEFT,
			DOWN,
			UP,
		},
	},
}

//go:embed all:tiles/demo-tracks
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

			WFCData[row+1][col] = Cell{collapsed: false, possible: possibilies[cellType][DDWON]}
		}
	}
	// UP
	if row >= rows {
		if !WFCData[row-1][col].collapsed {

			WFCData[row-1][col] = Cell{collapsed: false, possible: possibilies[cellType][DUP]}
		}
	}
	// RIGHT
	if col < cols {
		if !WFCData[row][col+1].collapsed {

			WFCData[row][col+1] = Cell{collapsed: false, possible: possibilies[cellType][DRIGHT]}
		}
	}
	// LEFT
	if col >= cols {
		if !WFCData[row][col-1].collapsed {

			WFCData[row][col-1] = Cell{collapsed: false, possible: possibilies[cellType][DLEFT]}
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
				l := len(cell.possible) - 1
				pos := 0

				if l <= 0 {
					l = 1
				} else {
					pos = rand.IntN(l)
				}

				WFCData = setCellType(WFCData, r, c, cell.possible[pos])
			}
		}
	}

	return WFCData
}

const scale float32 = 0.145

func main() {

	rows := (10)
	cols := (10)

	//resolution of photos
	width := (600)
	height := (600)

	var WFCData WFCData

	rl.SetConfigFlags(rl.FlagWindowUndecorated)
	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(int32(float32(rows*height)*scale), int32(float32(cols*width)*scale), "Wave Function Collapse")
	closeWindow := false

	tiles := map[CellType]rl.Texture2D{
		BLANK: loadTexture("tiles/demo-tracks/blank.png"),
		LEFT:  loadTexture("tiles/demo-tracks/left.png"),
		RIGHT: loadTexture("tiles/demo-tracks/right.png"),
		UP:    loadTexture("tiles/demo-tracks/up.png"),
		DOWN:  loadTexture("tiles/demo-tracks/left.png"),
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
