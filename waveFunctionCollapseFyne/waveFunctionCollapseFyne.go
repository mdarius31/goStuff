package main

//FYNE VERSION
/* program inspiration:
 * https://www.youtube.com/watch?v=rI_y2GAlQFM&list=PLB3OFCROxZ41eaR2Q4Ls27WjnzVoDLT6D
 */

import (
	"bytes"
	"embed"
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

//go:embed all:tiles
var tiles embed.FS

type CellType int

const (
	NOTHING CellType = iota
	BLANK
	LEFT
	RIGHT
	UP
	DOWN
)

func loadImage(s string) *canvas.Image {
	b, _ := tiles.ReadFile(s)
	d, _, _ := image.Decode(bytes.NewReader(b))

	i := canvas.NewImageFromImage(d)
	i.FillMode = canvas.ImageFillOriginal
	i.ScaleMode = canvas.ImageScalePixels

	return i
}

func resetTableData(tableData [][]CellType, rows int, cols int) [][]CellType {
	for r := 0; r < rows; r++ {

		var row []CellType

		for c := 0; c < cols; c++ {
			row = append(row, NOTHING)
		}

		tableData = append(tableData, row)
	}

	return tableData
}

func main() {

	// nothing := loadImage("./tiles/blank.png")
	// blank := loadImage("./tiles/blank.png")
	// left := loadImage("./tiles/left.png")
	// right := loadImage("./tiles/right.png")
	// up := loadImage("./tiles/up.png")
	// down := loadImage("./tiles/down.png")

	rows := 10
	cols := 10

	// firstX := rand.IntN(rows)
	// firstY := rand.IntN(cols)

	var tableData [][]CellType

	tableData = resetTableData(tableData, rows, cols)

	a := app.New()

	w := a.NewWindow("Wave Function Collapse")

	w.SetPadded(false)
	w.Resize(fyne.NewSize(float32(50*rows), float32(50*cols)))
	// w.SetFixedSize(true)
	w.CenterOnScreen()

	list := widget.NewTable(
		func() (int, int) {
			return rows, cols
		},

		func() fyne.CanvasObject {
			return loadImage("tiles/blank.png")
		},

		func(id widget.TableCellID, o fyne.CanvasObject) {
		},
	)

	w.SetContent(list)

	w.ShowAndRun()
}
