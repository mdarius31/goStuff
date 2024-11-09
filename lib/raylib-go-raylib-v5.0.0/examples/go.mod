module examples

go 1.21

replace raylib-go/raylib => ../raylib

replace raylib-go/raygui => ../raygui

replace raylib-go/easings => ../easings

replace raylib-go/physics => ../physics

replace purego => ../../../purego

require (
	raylib-go/easings v0.0.0-00010101000000-000000000000
	raylib-go/physics v0.0.0-00010101000000-000000000000
	raylib-go/raygui v0.0.0-00010101000000-000000000000
	raylib-go/raylib v0.0.0-20231118125650-a1c890e8cbfc
	github.com/jakecoffman/cp v1.2.1
	github.com/neguse/go-box2d-lite v0.0.0-20170921151050-5d8ed9b7272b
)

require (
	purego v0.5.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
)
