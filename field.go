package main

import rl "github.com/gen2brain/raylib-go/raylib"

func DrawRectangleLinesRec(rec rl.Rectangle, col rl.Color) {
	rl.DrawRectangleLines(int32(rec.X), int32(rec.Y), int32(rec.Width), int32(rec.Height), col)
}

// Gives back a n*n matrix initialized with rl.Rectangle
func FieldInit(n int) [][]rl.Rectangle {
	field := make([][]rl.Rectangle, 30)
	for i := range field {
		field[i] = make([]rl.Rectangle, 30)
	}

	for i := range field {
		for j := range field[i] {
			field[i][j] = rl.Rectangle{
				X:      float32(j * RECTSIZE),
				Y:      float32(i * RECTSIZE),
				Width:  float32(RECTSIZE),
				Height: float32(RECTSIZE),
			}
		}
	}
	return field
}
