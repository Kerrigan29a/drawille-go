package main

import (
	"fmt"

	drawille "github.com/Kerrigan29a/drawille-go"
)

func main() {

	fmt.Println("2 parallel lines")
	g := drawille.NewCanvas()
	g.DrawLine(0, 0, 50, 0)
	g.DrawLine(0, 20, 20, 20)
	fmt.Println(g.String())

	fmt.Println("A polygon of size 3")
	g.Clear()
	g.DrawPolygon(0, 0, 3, 30)
	fmt.Println(g.String())

	fmt.Println("A polygon of size 4")
	g.Clear()
	g.DrawPolygon(0, 0, 4, 30)
	fmt.Println(g.String())

	fmt.Println("A polygon of size 5")
	g.Clear()
	g.DrawPolygon(0, 0, 5, 30)
	fmt.Println(g.String())

	fmt.Println("A polygon of size 6")
	g.Clear()
	g.DrawPolygon(0, 0, 6, 30)
	fmt.Println(g.String())

	fmt.Println("A circle")
	g.Clear()
	g.DrawCircle(0, 0, 20)
	fmt.Println(g.String())

	fmt.Println("A filled circle")
	g.Clear()
	g.DrawCircle(0, 0, 20)
	g.Fill(0, 0)
	fmt.Println(g.String())

	fmt.Println("A donuts")
	g.Clear()
	g.DrawCircle(0, 0, 10)
	g.DrawCircle(0, 0, 20)
	g.Fill(10, 10)
	fmt.Println(g.String())

	fmt.Println("A Bézier quadratic curve")
	g.Clear()
	g.DrawBézier2(10, 75, 250, -50, 150, 140)
	fmt.Println(g.String())

	fmt.Println("A Bézier cubic curve")
	g.Clear()
	g.DrawBézier3(10, 100, 350, 25, -150, 25, 190, 75)
	fmt.Println(g.String())

}
