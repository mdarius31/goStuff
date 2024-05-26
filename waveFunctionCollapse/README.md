### Wave Function Collapse

R -> regenerate

Q -> Exit

### All Instructions were tested on Debian 12 Bookworm

### Dependencies

`libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev libwayland-dev libxkbcommon-dev`

`go mod tidy`

#### Build

`go build`

##### Compile For Windows

`CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" $ file waveFunctionCollapse.exe`

###### The Build result for windows was tested using wine-9.9 (Staging)
