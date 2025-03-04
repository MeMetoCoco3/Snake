package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREENWIDTH = 400
	SCREENHEIGHT

	RECTSIZE = 20
)

type Player struct {
	Position      rl.Vector2
	Direction     rl.Vector2
	NextDirection rl.Vector2
	Body          []rl.Vector2
}

// Position refers to position on matrix
type Candy struct {
	Position rl.Vector2
	Eaten    bool
}

var (
	PLAYERCOLOR = rl.Color{79, 119, 45, 255}
	Directions  = []rl.Vector2{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
)

func main() {
	rl.InitWindow(SCREENWIDTH, SCREENHEIGHT, "Snake")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// field := FieldInit(30)
	center := rl.Vector2{X: SCREENWIDTH/2 + RECTSIZE/2, Y: SCREENHEIGHT/2 + RECTSIZE/2}
	var playerPosition []rl.Vector2
	playerPosition = append(playerPosition, center)
	p := Player{
		Position:      center,
		Body:          playerPosition,
		Direction:     Directions[0],
		NextDirection: Directions[0],
	}

	BorderBox := rl.Rectangle{30, 30, SCREENWIDTH - 60, SCREENHEIGHT - 60}

	i := GenerateRandomPosition()
	j := GenerateRandomPosition()
	candy := Candy{Position: rl.Vector2{X: float32(i), Y: float32(j)}}

	for !rl.WindowShouldClose() {
		p.GetInput()
		p.UpdatePosition()
		candy.Eaten = false
		if EqualVectors(p.Position, candy.Position) {
			p.Grow()
			i := GenerateRandomPosition()
			j := GenerateRandomPosition()
			candy.Position = rl.Vector2{X: float32(i), Y: float32(j)}
			candy.Eaten = true
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		DrawRectangleLinesRec(BorderBox, rl.Red)

		if !candy.Eaten {
			rl.DrawRectangle(int32(candy.Position.X), int32(candy.Position.Y), RECTSIZE, RECTSIZE, rl.Red)
		}
		for i := range p.Body {
			rect := rl.Rectangle{p.Body[i].X, p.Body[i].Y, RECTSIZE, RECTSIZE}
			rl.DrawRectangleRec(rect, PLAYERCOLOR)
		}
		rl.EndDrawing()
	}
}

func (p *Player) GetInput() {
	if rl.IsKeyDown(rl.KeyUp) && p.Direction.Y == 0 {
		p.NextDirection = Directions[0]
	}
	if rl.IsKeyDown(rl.KeyRight) && p.Direction.X == 0 {
		p.NextDirection = Directions[1]
	}
	if rl.IsKeyDown(rl.KeyDown) && p.Direction.Y == 0 {
		p.NextDirection = Directions[2]
	}
	if rl.IsKeyDown(rl.KeyLeft) && p.Direction.X == 0 {
		p.NextDirection = Directions[3]
	}
}

func (p *Player) UpdatePosition() {
	if p.Position.X < RECTSIZE/2 || p.Position.X >= SCREENWIDTH-RECTSIZE/2 || p.Position.Y < RECTSIZE/2 || p.Position.Y >= SCREENHEIGHT-RECTSIZE/2 {
		os.Exit(1)
	}

	if len(p.Body) > 2 {
		for i := 2; i < len(p.Body); i++ {
			if EqualVectors(p.Body[0], p.Body[i]) {
				os.Exit(1)
			}
		}
	}

	if int(p.Position.X)%RECTSIZE == RECTSIZE/2 &&
		int(p.Position.Y)%RECTSIZE == RECTSIZE/2 {

		if p.NextDirection != p.Direction {
			p.Direction = p.NextDirection
		}
		p.Body = append([]rl.Vector2{p.Position}, p.Body[:len(p.Body)-1]...)
	}
	p.Position = Adition(p.Position, p.Direction)
}

func (p *Player) Grow() {
	tail := p.Body[len(p.Body)-1]
	p.Body = append(p.Body, tail)
}
