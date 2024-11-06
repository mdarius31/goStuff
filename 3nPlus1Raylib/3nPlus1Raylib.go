package main

import (
	"fmt"
	"math"

	t "threeNPlusOne"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	title := "3nPlus1"
	closeWindow := false

	fmt.Println(t.GenThreeNPlusOne(1))
	fmt.Println(t.GenThreeNPlusOne(10))

	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagVsyncHint)
	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(800, 450, title)

	rl.SetExitKey(rl.KeyNull)

	text := ""

	var val int32 = 0

	inputFocus := false

	focusable := []*bool{&inputFocus}

	recH := float32(30)
	rec := rl.NewRectangle(0, 200, float32(rl.GetScreenWidth()), recH)

	updateRec := func() {

		rec.Width = float32(rl.GetScreenWidth())
		rec.Y = float32(rl.GetScreenHeight()) - recH
	}
	updateRec()

	actions := map[uint]func(){
		rl.KeyQ: func() {
			closeWindow = true
		},
		rl.KeyEscape: func() {
			for _, focus := range focusable {
				*focus = false
			}
		},
	}

	defer rl.CloseWindow()

	for !rl.WindowShouldClose() && !closeWindow {
		if rl.IsWindowResized() {
			updateRec()
		}

		if rl.IsMouseButtonPressed(0) {
			actions[rl.KeyEscape]()
		}
		if action := actions[uint(rl.GetKeyPressed())]; action != nil {
			action()
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if rg.ValueBox(rec, text, &val, 0, math.MaxInt, inputFocus) {
			actions[rl.KeyEscape]()
			inputFocus = true
		}

		rl.EndDrawing()
	}

}
