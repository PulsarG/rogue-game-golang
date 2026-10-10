package main

import "github.com/gdamore/tcell/v2"

type Player struct {
	X, Y        int
	Sprite      rune
	ColorSprite tcell.Color
	MaxHp       int
	CurrentHp   int
	Dextery     int
	Strength    int
	Weapon      *Weapon
	FightWith   *Enemy
}

func PlayerInit(r *Room) *Player {
	return &Player{
		X:           r.CenterX,
		Y:           r.CenterY,
		Sprite:      '@',
		CurrentHp:   10,
		MaxHp:       10,
		Strength:    3,
		Dextery:     3,
		ColorSprite: tcell.ColorDefault,
		Weapon:      initWeapon(),
	}
}

func (p *Player) SetCoordinat(x, y int) {
	p.X = x
	p.Y = y
}
