// An implementation of Conway's Game of Life in Go.
package src

import (
	"bytes"
	"math/rand"
)

// Field represents a two-dimensional field of cells.
type Field struct {
	states [][]bool
	width  int
	height int
}

// NewField returns an empty field of the specified width and height.
func NewField(w, h int) *Field {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return &Field{states: s, width: w, height: h}
}

// Set sets the state of the specified cell to the given value.
func (field *Field) Set(x, y int, b bool) {
	field.states[y][x] = b
}

// Alive reports whether the specified cell is alive.
// If the x or y coordinates are outside the field boundaries they are wrapped
// toroidally. For instance, an x value of -1 is treated as width-1.
func (field *Field) Alive(x, y int) bool {
	x += field.width
	x %= field.width
	y += field.height
	y %= field.height
	return field.states[y][x]
}

// Next returns the state of the specified cell at the next time step.
func (field *Field) Next(x, y int) bool {
	// Count the adjacent cells that are alive.
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && field.Alive(x+i, y+j) {
				alive++
			}
		}
	}
	// Return next state according to the game rules:
	//   exactly 3 neighbors: on,
	//   exactly 2 neighbors: maintain current state,
	//   otherwise: off.
	return alive == 3 || alive == 2 && field.Alive(x, y)
}

// Life stores the state of a round of Conway's Game of Life.
type Life struct {
	a, b   *Field
	width  int
	height int
}

// NewLife returns a new Life game state with a random initial state.
func NewLife(width, height int) *Life {
	a := NewField(width, height)
	for i := 0; i < (width * height / 4); i++ {
		a.Set(rand.Intn(width), rand.Intn(height), true)
	}
	return &Life{
		a:     a, b: NewField(width, height),
		width: width, height: height,
	}
}

// Step advances the game by one instant, recomputing and updating all cells.
func (life *Life) Step() {
	// Update the state of the next field (b) from the current field (a).
	for y := 0; y < life.height; y++ {
		for x := 0; x < life.width; x++ {
			life.b.Set(x, y, life.a.Next(x, y))
		}
	}
	// Swap fields a and b.
	life.a, life.b = life.b, life.a
}

// String returns the game board as a string.
func (life *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < life.height; y++ {
		for x := 0; x < life.width; x++ {
			b := "  "
			if life.a.Alive(x, y) {
				b = "\u2588\u2588"
			}
			buf.WriteString(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
