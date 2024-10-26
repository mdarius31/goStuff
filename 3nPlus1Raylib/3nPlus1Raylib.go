package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func gen3nPlus1(num int64) []int64 {
	var out []int64

	for num != 1 {

		out = append(out, num)
		even := num%2 == 0
		if even {
			num = num / 2
		} else {
			num = (num * 3) + 1
		}
	}

	out = append(out, 1)

	return out
}

func main() {

	fmt.Println(gen3nPlus1(10))

	rl.SetConfigFlags(rl.FlagWindowUndecorated)
	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
