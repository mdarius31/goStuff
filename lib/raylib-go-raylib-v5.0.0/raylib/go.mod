module github.com/gen2brain/raylib-go/raylib

go 1.21

replace purego => ../../purego

require (
	golang.org/x/sys v0.14.0
	purego v0.0.0-00010101000000-000000000000
)

require github.com/ebitengine/purego v0.8.1 // indirect
