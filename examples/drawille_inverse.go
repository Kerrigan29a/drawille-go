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

const N int = 12

func main() {

	fmt.Println("x>=0 && y>=0 && !inverse")
	g := drawille.NewCanvas()
	g.Inverse = false
	for x := 0; x < N; x++ {
		for y := 0; y < N-x; y++ {
			g.Set(x, y)
		}
	}
	fmt.Print(g.String())

	fmt.Println("x>=0 && y>=0 && inverse")
	g.Clear()
	g.Inverse = true
	for x := 0; x < N; x++ {
		for y := 0; y < N-x; y++ {
			g.Set(x, y)
		}
	}
	fmt.Print(g.String())

	fmt.Println("x<0 && y>=0 && !inverse")
	g.Clear()
	g.Inverse = false
	for x := -N; x < 0; x++ {
		for y := 0; y <= N+x; y++ {
			g.Set(x, y)
		}
	}
	fmt.Print(g.String())

	fmt.Println("x<0 && y>=0 && inverse")
	g.Clear()
	g.Inverse = true
	for x := -N; x < 0; x++ {
		for y := 0; y <= N+x; y++ {
			g.Set(x, y)
		}
	}
	fmt.Print(g.String())

	fmt.Println("x>=0 && y<0 && !inverse")
	g.Clear()
	g.Inverse = false
	for x := 0; x < N; x++ {
		for y := -N + x; y < 0; y++ {
			g.Set(x, y)
		}
	}
	fmt.Print(g.String())

	fmt.Println("x>=0 && y<0 && inverse")
	g.Clear()
	g.Inverse = true
	for x := 0; x < N; x++ {
		for y := -N + x; y < 0; y++ {
			g.Set(x, y)
		}
	}
	fmt.Print(g.String())

	fmt.Println("x<0 && y<0 && !inverse")
	g.Clear()
	g.Inverse = false
	for x := -N; x < 0; x++ {
		for y := -N - x - 1; y < 0; y++ {
			g.Set(x, y)
		}
	}
	fmt.Print(g.String())

	fmt.Println("x<0 && y<0 && inverse")
	g.Clear()
	g.Inverse = true
	for x := -N; x < 0; x++ {
		for y := -N - x - 1; y < 0; y++ {
			g.Set(x, y)
		}
	}
	fmt.Print(g.String())
}
