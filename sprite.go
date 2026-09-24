package main

import "github.com/gdamore/tcell/v2"

type Sprite struct {
	Char rune
	X, Y int
}

func NewSprite(char rune, x, y int) *Sprite {
	return &Sprite{
		Char: char,
		X:    x,
		Y:    y,
	}
}

func Draw(x, y int, sprite rune, screen tcell.Screen, color *tcell.Color) {
	screen.SetContent(
		x,
		y,
		sprite,
		nil,
		tcell.StyleDefault.Foreground(*color),
	)
}
