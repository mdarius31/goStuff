package main

/* program inspiration:
 * https://www.youtube.com/watch?v=rI_y2GAlQFM&list=PLB3OFCROxZ41eaR2Q4Ls27WjnzVoDLT6D
 */

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CellType int

const (
	NOTHING CellType = iota
	BLANK
	LEFT
	RIGHT
	UP
	DOWN
)

//go:embed all:tiles/demo
var tiles embed.FS

func loadTexture(s string) rl.Texture2D {
	bytes, _ := tiles.ReadFile(s)

	img := rl.LoadImageFromMemory(".png", bytes, int32(len(bytes)))
	texture := rl.LoadTextureFromImage(img)

	return texture

}

func main() {
	rows := int32(10)
	cols := int32(10)

	width := int32(50)
	height := int32(50)

	var WFCData [][]CellType

	for r := int32(0); r < rows; r++ {
		var row []CellType

		for c := int32(0); c < cols; c++ {
			row = append(row, UP)
		}
		WFCData = append(WFCData, row)
	}

	rl.SetConfigFlags(rl.FlagWindowUndecorated)
	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(rows*height, cols*width, "Wave Function Collapse")

	blank := loadTexture("tiles/demo/blank.png")
	left := loadTexture("tiles/demo/left.png")
	right := loadTexture("tiles/demo/right.png")
	up := loadTexture("tiles/demo/up.png")
	down := loadTexture("tiles/demo/left.png")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for r := int32(0); r < int32(len(WFCData)); r++ {
			for c := int32(0); c < int32(len(WFCData[r])); c++ {

				x := r * width
				y := c * height
				switch WFCData[r][c] {
				case NOTHING:
					rl.DrawTexture(blank, x, y, rl.White)

				case BLANK:
					rl.DrawTexture(blank, x, y, rl.White)

				case DOWN:
					rl.DrawTexture(down, x, y, rl.White)

				case UP:
					rl.DrawTexture(up, x, y, rl.White)

				case RIGHT:
					rl.DrawTexture(right, x, y, rl.White)

				case LEFT:
					rl.DrawTexture(left, x, y, rl.White)

				default:
					rl.DrawTexture(blank, x, y, rl.White)

				}
			}
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
