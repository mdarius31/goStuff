module 3nPlus1Raylib

go 1.22.2

replace threeNPlusOne => ../lib/threeNPlusOne

replace raylib => ../lib/raylib-go-raylib-v5.0.0/raylib

require (
	raylib v0.0.0-00010101000000-000000000000
	threeNPlusOne v0.0.0-00010101000000-000000000000
)

require golang.org/x/sys v0.14.0 // indirect
