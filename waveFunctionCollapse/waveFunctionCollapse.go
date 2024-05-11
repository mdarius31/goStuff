package main

/* program inspiration:
 * https://www.youtube.com/watch?v=rI_y2GAlQFM&list=PLB3OFCROxZ41eaR2Q4Ls27WjnzVoDLT6D
 */

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
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

type CellData struct {
	cellType CellType
	row      int
	col      int
}

func loadImage(s string) *canvas.Image {
	i := canvas.NewImageFromFile(s)
	i.FillMode = canvas.ImageFillOriginal
	i.ScaleMode = canvas.ImageScalePixels

	return i
}

func main() {

	// nothing := loadImage("./tiles/demo/blank.png")
	// blank := loadImage("./tiles/demo/blank.png")
	// left := loadImage("./tiles/demo/left.png")
	// right := loadImage("./tiles/demo/right.png")
	// up := loadImage("./tiles/demo/up.png")
	// down := loadImage("./tiles/demo/down.png")

	rows := 10
	cols := 10

	var tableData [][]CellData

	for r := 0; r < rows; r++ {

		var row []CellData

		for c := 0; c < cols; c++ {
			row = append(row, CellData{cellType: NOTHING, row: r, col: c})
		}

		tableData = append(tableData, row)
	}

	a := app.New()
	w := a.NewWindow("Wave Function Collapse")

	list := widget.NewTable(
		func() (int, int) {
			return rows, cols
		},

		func() fyne.CanvasObject {
			return loadImage("./tiles/demo/blank.png")
		},

		func(id widget.TableCellID, o fyne.CanvasObject) {

		},
	)

	w.SetContent(list)

	w.ShowAndRun()
}
