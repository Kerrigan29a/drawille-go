package main

import (
	"fmt"

	drawille "github.com/Kerrigan29a/drawille-go"
)

/*
C:\Users\VirtualUser\go\src\github.com\Kerrigan29a\drawille-go (master -> origin)
λ go run examples\test_drawille.go
⠙⢿⣿⣿⣿⣿
⠀⠀⠙⢿⣿⣿
⠀⠀⠀⠀⠙⢿
⠀⠀⠀⠀⢉⣽
⠀⠀⢉⣽⣿⣿
⢉⣽⣿⣿⣿⣿
*/
func main() {
	g := drawille.NewCanvas()
	for i := 0; i < 12; i++ {
		for j := 0; j <= i; j++ {
			g.Set(i, j)
		}
	}
	fmt.Print(g.String())

	g.Clear()
	g.Inverse = true
	for i := 0; i < 12; i++ {
		for j := 0; j <= i; j++ {
			g.Set(i, j)
		}
	}
	fmt.Print(g.String())
}
