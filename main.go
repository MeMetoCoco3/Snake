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
	Frame         float32
}

// Position refers to position on matrix
type Candy struct {
	Position rl.Vector2
	Eaten    bool
}

var (
	PLAYERCOLOR     = rl.Color{6, 214, 160, 255}
	GRIDCOLOR       = rl.Color{7, 59, 76, 80}
	BORDERCOLOR     = rl.Color{7, 59, 76, 255}
	CANDYCOLOR      = rl.Color{239, 45, 94, 255}
	BACKGROUNDCOLOR = rl.Color{255, 209, 102, 255}
	Directions      = []rl.Vector2{
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
		rl.ClearBackground(BACKGROUNDCOLOR)
		DrawRectangleLinesRec(BorderBox, BORDERCOLOR)
		for i := 0; i < SCREENHEIGHT; i = i + RECTSIZE {
			rl.DrawLine(int32(i), 0, int32(i), SCREENHEIGHT, GRIDCOLOR)
			rl.DrawLine(0, int32(i), SCREENWIDTH, int32(i), GRIDCOLOR)
		}

		if !candy.Eaten {
			rl.DrawRectangle(int32(candy.Position.X), int32(candy.Position.Y), RECTSIZE, RECTSIZE, CANDYCOLOR)
		}
		for i := range p.Body {
			if i == 0 {
				rect := rl.Rectangle{p.Position.X - p.Direction.X, p.Position.Y - p.Direction.Y, RECTSIZE, RECTSIZE}
				rl.DrawRectangleRec(rect, PLAYERCOLOR)
				if len(p.Body) == 1 {
					continue
				}
				rect = rl.Rectangle{p.Body[i].X, p.Body[i].Y, RECTSIZE, RECTSIZE}
				rl.DrawRectangleRec(rect, PLAYERCOLOR)
				continue
			}
			var dx float32
			var dy float32
			if p.Body[i].X == p.Body[i-1].X {
				dx = 0
				if p.Body[i].Y < p.Body[i-1].Y {
					dy = 1
				} else {
					dy = -1
				}
			} else if p.Body[i].X > p.Body[i-1].X {
				dx = -1
				dy = 0
			} else {
				dx = 1
				dy = 0
			}
			rect := rl.Rectangle{
				p.Body[i].X + (dx * p.Frame), p.Body[i].Y + (dy * p.Frame),
				RECTSIZE, RECTSIZE,
			}

			rl.DrawRectangleRec(rect, PLAYERCOLOR)
			if i != len(p.Body)-1 {
				rect = rl.Rectangle{
					p.Body[i].X, p.Body[i].Y,
					RECTSIZE, RECTSIZE,
				}
				rl.DrawRectangleRec(rect, PLAYERCOLOR)
			}

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
	p.Frame++
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
		p.Frame = 0
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
