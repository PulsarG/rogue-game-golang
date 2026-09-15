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

func Draw(p *Player, screen tcell.Screen) {
	screen.SetContent(
		p.X,
		p.Y,
		p.Sprite,
		nil,
		tcell.StyleDefault,
	)
}
