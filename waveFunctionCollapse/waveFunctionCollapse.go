package main

/* program inspiration:
 * https://www.youtube.com/watch?v=rI_y2GAlQFM&list=PLB3OFCROxZ41eaR2Q4Ls27WjnzVoDLT6D
 */

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello, World!")

	w.SetContent(widget.NewLabel("Hello, World!"))
	w.ShowAndRun()
}
