package drawille

import (
	"fmt"
	"math"
)

// Braille chars start at 0x2800
var brailleStartOrdinal = 0x2800

func internalPosition(n, base int) int {
	if n >= 0 {
		return n % base
	}
	result := n % base
	if result == 0 {
		return -base
	}
	return result
}

func getDot(y, x int, inverse bool) int {
	y = internalPosition(y, 4)
	x = internalPosition(x, 2)

	/* x>=0 && y>=0 && !inverse */
	if y == 0 && x == 0 && !inverse {
		return 0x1
	}
	if y == 1 && x == 0 && !inverse {
		return 0x2
	}
	if y == 2 && x == 0 && !inverse {
		return 0x4
	}
	if y == 3 && x == 0 && !inverse {
		return 0x40
	}
	if y == 0 && x == 1 && !inverse {
		return 0x8
	}
	if y == 1 && x == 1 && !inverse {
		return 0x10
	}
	if y == 2 && x == 1 && !inverse {
		return 0x20
	}
	if y == 3 && x == 1 && !inverse {
		return 0x80
	}

	/* x>=0 && y>=0 && inverse */
	if y == 0 && x == 0 && inverse {
		return 0x40
	}
	if y == 1 && x == 0 && inverse {
		return 0x4
	}
	if y == 2 && x == 0 && inverse {
		return 0x2
	}
	if y == 3 && x == 0 && inverse {
		return 0x1
	}
	if y == 0 && x == 1 && inverse {
		return 0x80
	}
	if y == 1 && x == 1 && inverse {
		return 0x20
	}
	if y == 2 && x == 1 && inverse {
		return 0x10
	}
	if y == 3 && x == 1 && inverse {
		return 0x8
	}

	/* x<0 && y<0 && !inverse */
	if y == -1 && x == -1 && !inverse {
		return 0x80
	}
	if y == -2 && x == -1 && !inverse {
		return 0x20
	}
	if y == -3 && x == -1 && !inverse {
		return 0x10
	}
	if y == -4 && x == -1 && !inverse {
		return 0x8
	}
	if y == -1 && x == -2 && !inverse {
		return 0x40
	}
	if y == -2 && x == -2 && !inverse {
		return 0x4
	}
	if y == -3 && x == -2 && !inverse {
		return 0x2
	}
	if y == -4 && x == -2 && !inverse {
		return 0x1
	}

	/* x<0 && y<0 && inverse */
	if y == -1 && x == -1 && inverse {
		return 0x8
	}
	if y == -2 && x == -1 && inverse {
		return 0x10
	}
	if y == -3 && x == -1 && inverse {
		return 0x20
	}
	if y == -4 && x == -1 && inverse {
		return 0x80
	}
	if y == -1 && x == -2 && inverse {
		return 0x1
	}
	if y == -2 && x == -2 && inverse {
		return 0x2
	}
	if y == -3 && x == -2 && inverse {
		return 0x4
	}
	if y == -4 && x == -2 && inverse {
		return 0x40
	}

	/* x>=0 && y<0 && !inverse */
	if y == -1 && x == 0 && !inverse {
		return 0x40
	}
	if y == -2 && x == 0 && !inverse {
		return 0x4
	}
	if y == -3 && x == 0 && !inverse {
		return 0x2
	}
	if y == -4 && x == 0 && !inverse {
		return 0x1
	}
	if y == -1 && x == 1 && !inverse {
		return 0x80
	}
	if y == -2 && x == 1 && !inverse {
		return 0x20
	}
	if y == -3 && x == 1 && !inverse {
		return 0x10
	}
	if y == -4 && x == 1 && !inverse {
		return 0x8
	}

	/* x>=0 && y<0 && inverse */
	if y == -1 && x == 0 && inverse {
		return 0x1
	}
	if y == -2 && x == 0 && inverse {
		return 0x2
	}
	if y == -3 && x == 0 && inverse {
		return 0x4
	}
	if y == -4 && x == 0 && inverse {
		return 0x40
	}
	if y == -1 && x == 1 && inverse {
		return 0x8
	}
	if y == -2 && x == 1 && inverse {
		return 0x10
	}
	if y == -3 && x == 1 && inverse {
		return 0x20
	}
	if y == -4 && x == 1 && inverse {
		return 0x80
	}

	/* x<0 && y>=0 && !inverse */
	if y == 0 && x == -1 && !inverse {
		return 0x8
	}
	if y == 1 && x == -1 && !inverse {
		return 0x10
	}
	if y == 2 && x == -1 && !inverse {
		return 0x20
	}
	if y == 3 && x == -1 && !inverse {
		return 0x80
	}
	if y == 0 && x == -2 && !inverse {
		return 0x1
	}
	if y == 1 && x == -2 && !inverse {
		return 0x2
	}
	if y == 2 && x == -2 && !inverse {
		return 0x4
	}
	if y == 3 && x == -2 && !inverse {
		return 0x40
	}

	/* x<0 && y>=0 && inverse */
	if y == 0 && x == -1 && inverse {
		return 0x80
	}
	if y == 1 && x == -1 && inverse {
		return 0x20
	}
	if y == 2 && x == -1 && inverse {
		return 0x10
	}
	if y == 3 && x == -1 && inverse {
		return 0x8
	}
	if y == 0 && x == -2 && inverse {
		return 0x40
	}
	if y == 1 && x == -2 && inverse {
		return 0x4
	}
	if y == 2 && x == -2 && inverse {
		return 0x2
	}
	if y == 3 && x == -2 && inverse {
		return 0x1
	}

	panic(fmt.Sprintf("Unknown values: y=%d x=%d inverse=%t", y, x, inverse))
}

func externalPosition(n, base int) int {
	return int(math.Floor(float64(n) / float64(base)))
}

// Convert x,y to cols, rows
func getPos(x, y int) (int, int) {
	c := externalPosition(x, 2)
	r := externalPosition(y, 4)
	return c, r
}

type Canvas struct {
	LineEnding string
	Inverse    bool
	chars      map[int]map[int]int
}

// Make a new canvas
func NewCanvas() Canvas {
	c := Canvas{LineEnding: "\n", Inverse: false}
	c.Clear()
	return c
}

func (c Canvas) MaxY() int {
	max := 0
	for k, _ := range c.chars {
		if k > max {
			max = k
		}
	}
	return max * 4
}

func (c Canvas) MinY() int {
	min := 0
	for k, _ := range c.chars {
		if k < min {
			min = k
		}
	}
	return min * 4
}

func (c Canvas) MaxX() int {
	max := 0
	for _, v := range c.chars {
		for k, _ := range v {
			if k > max {
				max = k
			}
		}
	}
	return max * 2
}

func (c Canvas) MinX() int {
	min := 0
	for _, v := range c.chars {
		for k, _ := range v {
			if k < min {
				min = k
			}
		}
	}
	return min * 2
}

// Clear all pixels
func (c *Canvas) Clear() {
	c.chars = make(map[int]map[int]int)
}

// Set a pixel of c
func (c *Canvas) Set(x, y int) {
	col, row := getPos(x, y)
	if m := c.chars[row]; m == nil {
		c.chars[row] = make(map[int]int)
	}
	val := c.chars[row][col]
	mapv := getDot(y, x, c.Inverse)
	c.chars[row][col] = val | mapv
}

// Unset a pixel of c
func (c *Canvas) UnSet(x, y int) {
	col, row := getPos(x, y)
	if m := c.chars[row]; m == nil {
		c.chars[row] = make(map[int]int)
	}
	c.chars[row][col] &^= getDot(y, x, c.Inverse)
}

// Toggle a point
func (c *Canvas) Toggle(x, y int) {
	col, row := getPos(x, y)
	if m := c.chars[row]; m == nil {
		c.chars[row] = make(map[int]int)
	}
	c.chars[row][col] ^= getDot(y, x, c.Inverse)
}

// Set text to the given coordinates
func (c *Canvas) SetText(x, y int, text string) {
	col, row := getPos(x, y)
	if m := c.chars[row]; m == nil {
		c.chars[row] = make(map[int]int)
	}
	for i, char := range text {
		c.chars[row][col+i] = int(char) - brailleStartOrdinal
	}
}

// Get pixel at the given coordinates
func (c Canvas) Get(x, y int) bool {
	dot := getDot(y, x, c.Inverse)
	col, row := getPos(x, y)
	char := c.chars[row][col]
	return (char & dot) != 0
}

// Get character at the given screen coordinates
func (c Canvas) GetScreenCharacter(x, y int) rune {
	return rune(c.chars[y][x] + brailleStartOrdinal)
}

// Get character for the given pixel
func (c Canvas) GetCharacter(x, y int) rune {
	return c.GetScreenCharacter(x/4, y/4)
}

// Retrieve the rows from a given view
func (c Canvas) Rows(minX, minY, maxX, maxY int) []string {
	minRow, maxRow := minY/4, maxY/4
	minCol, maxCol := minX/2, maxX/2

	txts := make([]string, 0)
	if c.Inverse {
		for row := maxRow; row >= minRow; row-- {
			txts = append(txts, c.line(row, minCol, maxCol))
		}
	} else {
		for row := minRow; row <= maxRow; row++ {
			txts = append(txts, c.line(row, minCol, maxCol))
		}
	}
	return txts
}

func (c Canvas) line(row, minCol, maxCol int) string {
	txt := []rune{}
	for col := minCol; col <= maxCol; col++ {
		char := c.chars[row][col]
		txt = append(txt, rune(char+brailleStartOrdinal))
	}
	return string(txt)
}

// Retrieve a string representation of the frame at the given parameters
func (c Canvas) Frame(minX, minY, maxX, maxY int) string {
	var txt string
	for _, row := range c.Rows(minX, minY, maxX, maxY) {
		txt += row
		txt += c.LineEnding
	}
	return txt
}

func (c Canvas) String() string {
	return c.Frame(c.MinX(), c.MinY(), c.MaxX(), c.MaxY())
}

/*
Fill implements the Flood fill algorithm.

From: https://rosettacode.org/wiki/Bitmap/Flood_fill#Go
*/
func (c *Canvas) Fill(x, y int) {
	target := c.Get(x, y)
	tasks := [][2]int{[2]int{x, y}}
	for {
		if len(tasks) == 0 {
			break
		}
		var task [2]int
		task, tasks = tasks[0], tasks[1:]
		tx := task[0]
		ty := task[1]
		current := c.Get(tx, ty)
		if current == target {
			c.Set(tx, ty)
			if c.Get(tx, ty) {
				tasks = append(tasks, [2]int{tx - 1, ty})
			}
			if c.Get(tx, ty) {
				tasks = append(tasks, [2]int{tx + 1, ty})
			}
			if c.Get(tx, ty) {
				tasks = append(tasks, [2]int{tx, ty - 1})
			}
			if c.Get(tx, ty) {
				tasks = append(tasks, [2]int{tx, ty + 1})
			}
		}
	}
}

/*
DrawLine plots a line using the Bresenham's algorithm

From: https://rosettacode.org/wiki/Bitmap/Bresenham%27s_line_algorithm#Go
*/
func (c *Canvas) DrawLine(x0, y0, x1, y1 int) {
	/* Implemented straight from WP pseudocode */
	dx := x1 - x0
	if dx < 0 {
		dx = -dx
	}
	dy := y1 - y0
	if dy < 0 {
		dy = -dy
	}
	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	for {
		c.Set(x0, y0)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

/*
DrawCircle plots a circle with center x, y and radius r using the Bresenham's algorithm.

Limiting behavior:
r < 0 plots no pixels.
r = 0 plots a single pixel at x, y.
r = 1 plots four pixels in a diamond shape around the center pixel at x, y.

From: https://rosettacode.org/wiki/Bitmap/Midpoint_circle_algorithm#Go
*/
func (c *Canvas) DrawCircle(x, y, r int) {
	if r < 0 {
		return
	}
	/* Bresenham algorithm */
	x1, y1, err := -r, 0, 2-2*r
	for {
		c.Set(x-x1, y+y1)
		c.Set(x-y1, y-x1)
		c.Set(x+x1, y-y1)
		c.Set(x+y1, y+x1)
		r = err
		if r > x1 {
			x1++
			err += x1*2 + 1
		}
		if r <= y1 {
			y1++
			err += y1*2 + 1
		}
		if x1 >= 0 {
			break
		}
	}
}

const b2Seg = 20

/*
DrawBezier2 draws a quadratic Bézier curve.

From: https://rosettacode.org/wiki/Bitmap/B%C3%A9zier_curves/Quadratic#Go
*/
func (c *Canvas) DrawBezier2(x1, y1, x2, y2, x3, y3 int) {
	var px, py [b2Seg + 1]int
	fx1, fy1 := float64(x1), float64(y1)
	fx2, fy2 := float64(x2), float64(y2)
	fx3, fy3 := float64(x3), float64(y3)
	for i := range px {
		c := float64(i) / b2Seg
		a := 1 - c
		a, b, c := a*a, 2*c*a, c*c
		px[i] = int(a*fx1 + b*fx2 + c*fx3)
		py[i] = int(a*fy1 + b*fy2 + c*fy3)
	}
	x0, y0 := px[0], py[0]
	for i := 1; i <= b2Seg; i++ {
		x1, y1 := px[i], py[i]
		c.DrawLine(x0, y0, x1, y1)
		x0, y0 = x1, y1
	}
}

const b3Seg = 30

/*
DrawBezier3 draws a cubic Bézier curve.

From: https://rosettacode.org/wiki/Bitmap/B%C3%A9zier_curves/Cubic#Go
*/
func (c *Canvas) DrawBezier3(x1, y1, x2, y2, x3, y3, x4, y4 int) {
	var px, py [b3Seg + 1]int
	fx1, fy1 := float64(x1), float64(y1)
	fx2, fy2 := float64(x2), float64(y2)
	fx3, fy3 := float64(x3), float64(y3)
	fx4, fy4 := float64(x4), float64(y4)
	for i := range px {
		d := float64(i) / b3Seg
		a := 1 - d
		b, c := a*a, d*d
		a, b, c, d = a*b, 3*b*d, 3*a*c, c*d
		px[i] = int(a*fx1 + b*fx2 + c*fx3 + d*fx4)
		py[i] = int(a*fy1 + b*fy2 + c*fy3 + d*fy4)
	}
	x0, y0 := px[0], py[0]
	for i := 1; i <= b3Seg; i++ {
		x1, y1 := px[i], py[i]
		c.DrawLine(x0, y0, x1, y1)
		x0, y0 = x1, y1
	}
}

/*
DrawPolygon plots a polygon of N sides inside the circumscribed circle of radius R.
*/
func (c *Canvas) DrawPolygon(centerX, centerY, sides, radius int) {
	xs, ys := PolygonVertices(centerX, centerY, sides, radius)
	px := xs[sides-1]
	py := ys[sides-1]
	for i := 0; i < sides; i++ {
		x := xs[i]
		y := ys[i]
		c.DrawLine(px, py, x, y)
		px = x
		py = y
	}
}

/*
PolygonVertices calculates the vertices of a polygon of N sides inside the circumscribed circle of radius R.
*/
func PolygonVertices(centerX, centerY, sides, radius int) ([]int, []int) {
	degree := 360 / float64(sides)
	x0 := float64(centerX)
	y0 := float64(centerY)
	r := float64(radius)
	xs := make([]int, sides)
	ys := make([]int, sides)
	for i := 0; i < sides; i++ {
		a := radians(float64(i) * degree)
		xs = append(xs, round((x0 + (math.Cos(a) * r))))
		ys = append(ys, round((y0 + (math.Sin(a) * r))))
	}
	return xs, ys
}

func radians(d float64) float64 {
	return d * (math.Pi / 180)
}

func round(x float64) int {
	return int(x + 0.5)
}
