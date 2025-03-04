package main

import (
	"log"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Adition(a, b rl.Vector2) rl.Vector2 {
	return rl.Vector2{X: a.X + b.X, Y: a.Y + b.Y}
}

func EqualVectors(v1, v2 rl.Vector2) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

func GetRandom(max int) int {
	return rand.Intn(max)
}

func GenerateRandomPosition() int {
	maxArg := (SCREENWIDTH - RECTSIZE*2) / RECTSIZE
	i := GetRandom(maxArg)
	log.Println(i)
	if i <= 1 {
		i = 2
	} else if i >= maxArg-1 {
		i = maxArg - 3
	}
	return i*RECTSIZE + RECTSIZE/2
}
