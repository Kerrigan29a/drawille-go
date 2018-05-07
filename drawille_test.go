package drawille

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	c := NewCanvas()
	c.Set(0, 0)

	assert.True(t, c.Get(0, 0))
}

func TestUnsetEmpty(t *testing.T) {
	c := NewCanvas()
	c.Set(1, 1)
	c.UnSet(1, 1)

	assert.False(t, c.Get(0, 0))
	assert.False(t, c.Get(0, 1))
	assert.False(t, c.Get(1, 0))
	assert.False(t, c.Get(1, 1))
}

func TestUnsetNonempty(t *testing.T) {
	c := NewCanvas()
	c.Set(0, 0)
	c.Set(0, 1)
	c.UnSet(0, 1)
	assert.True(t, c.Get(0, 0))
}

func TestClear(t *testing.T) {
	c := NewCanvas()
	c.Set(1, 1)
	c.Clear()
	assert.False(t, c.Get(1, 1))
}

func TestToggle(t *testing.T) {
	c := NewCanvas()
	c.Toggle(0, 0)
	assert.True(t, c.Get(0, 0))
	c.Toggle(0, 0)
	assert.False(t, c.Get(0, 0))
}

func TestSetText(t *testing.T) {
	c := NewCanvas()
	c.SetText(0, 0, "asdf")
	assert.Equal(t, "asdf\n", c.Frame(0, 0, 7, 3))
}

func TestFrame(t *testing.T) {
	c := NewCanvas()
	assert.Equal(t, "\u2800\n", c.Frame(0, 0, 1, 3))
	c.Set(0, 0)
	assert.Equal(t, "⠁\n", c.Frame(0, 0, 1, 3))
}

// test_max_min_limits could not be ported

func TestGet(t *testing.T) {
	c := NewCanvas()
	assert.False(t, c.Get(0, 0))
	c.Set(0, 0)
	assert.True(t, c.Get(0, 0))
	assert.False(t, c.Get(0, 1))
	assert.False(t, c.Get(1, 0))
	assert.False(t, c.Get(1, 1))
}

func TestDrawLineSinglePixel(t *testing.T) {
	c := NewCanvas()
	c.DrawLine(0, 0, 0, 0)
	assert.True(t, c.Get(0, 0))
}

func TestDrawLineRow(t *testing.T) {
	c := NewCanvas()
	c.DrawLine(0, 0, 1, 0)
	assert.True(t, c.Get(0, 0))
	assert.True(t, c.Get(1, 0))
}

func TestDrawLineColumn(t *testing.T) {
	c := NewCanvas()
	c.DrawLine(0, 0, 0, 1)
	assert.True(t, c.Get(0, 0))
	assert.True(t, c.Get(0, 1))
}

func TestDrawLineDiagonal(t *testing.T) {
	c := NewCanvas()
	c.DrawLine(0, 0, 1, 1)
	assert.True(t, c.Get(0, 0))
	assert.True(t, c.Get(1, 1))
}

func TestDrawLineDiagonalNonSquare(t *testing.T) {
	c := NewCanvas()
	c.DrawLine(0, 0, 2, 3)
	assert.True(t, c.Get(0, 0))
	assert.True(t, c.Get(1, 1))
	assert.True(t, c.Get(1, 2))
	assert.True(t, c.Get(2, 3))
}

func TestGetDot(t *testing.T) {
	assert.Equal(t, getDot(0, 0, false), 0x1)
	assert.Equal(t, getDot(1, 0, false), 0x2)
	assert.Equal(t, getDot(2, 0, false), 0x4)
	assert.Equal(t, getDot(3, 0, false), 0x40)
	assert.Equal(t, getDot(0, 1, false), 0x8)
	assert.Equal(t, getDot(1, 1, false), 0x10)
	assert.Equal(t, getDot(2, 1, false), 0x20)
	assert.Equal(t, getDot(3, 1, false), 0x80)

	assert.Equal(t, getDot(4, 2, false), 0x1)
	assert.Equal(t, getDot(5, 2, false), 0x2)
	assert.Equal(t, getDot(6, 2, false), 0x4)
	assert.Equal(t, getDot(7, 2, false), 0x40)
	assert.Equal(t, getDot(4, 3, false), 0x8)
	assert.Equal(t, getDot(5, 3, false), 0x10)
	assert.Equal(t, getDot(6, 3, false), 0x20)
	assert.Equal(t, getDot(7, 3, false), 0x80)

	assert.Equal(t, getDot(0, 0, true), 0x40)
	assert.Equal(t, getDot(1, 0, true), 0x4)
	assert.Equal(t, getDot(2, 0, true), 0x2)
	assert.Equal(t, getDot(3, 0, true), 0x1)
	assert.Equal(t, getDot(0, 1, true), 0x80)
	assert.Equal(t, getDot(1, 1, true), 0x20)
	assert.Equal(t, getDot(2, 1, true), 0x10)
	assert.Equal(t, getDot(3, 1, true), 0x8)

	assert.Equal(t, getDot(4, 2, true), 0x40)
	assert.Equal(t, getDot(5, 2, true), 0x4)
	assert.Equal(t, getDot(6, 2, true), 0x2)
	assert.Equal(t, getDot(7, 2, true), 0x1)
	assert.Equal(t, getDot(4, 3, true), 0x80)
	assert.Equal(t, getDot(5, 3, true), 0x20)
	assert.Equal(t, getDot(6, 3, true), 0x10)
	assert.Equal(t, getDot(7, 3, true), 0x8)
}
