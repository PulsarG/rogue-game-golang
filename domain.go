package main

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

// Игровая сессия;
type Session struct {
	CurrentLvl                          int
	Mapp                                [100][40]rune
	Player                              *Player
	Rooms                               [9]Room
	StartRoomIdx, EndRoomIdx            int
	Enemys                              []*Enemy
	PlayerStatusDoing, EnemyStatusDoing string
	CurrentState                        GameState
	EndRoom                             *EndRoom
	// возможно данные для статистики
}

type GameState int

const (
	StateMainMenu GameState = iota
	StatePlay
	StateGameMenu
	StateInventory
	StateGameOver
	StateStartNewGame
	StateNextLvl
)

type EndRoom struct {
	X, Y        int
	Sprite      rune
	ColorSprite tcell.Color
}

func (s *Session) EndRoomInit() *EndRoom {
	e := EndRoom{}
	for {
		room := rand.Intn(9)
		if s.Rooms[room].IsStart {
			continue
		} else {
			nicePoint := false
			for !nicePoint {
				x := s.Rooms[room].CenterX + rand.Intn(3)
				y := s.Rooms[room].CenterY + rand.Intn(3)
				if s.Mapp[x][y] != '.' {
					continue
				} else {
					e.X = x
					e.Y = y
					s.Mapp[e.X][e.Y] = '%'
					nicePoint = true
				}
			}
			e.Sprite = '%'
			break
		}
	}
	e.ColorSprite = tcell.ColorGreen
	return &e
}

// Уровень;
type Level struct {
	CountEnemy int
}

// Комната;
type Room struct {
	IsStart, IsEnd                                  bool
	Width, Heigth                                   int
	Enemies                                         int
	Door                                            int
	StartPointFromX, StartPointFromY, EndPointFromY int
	CenterX, CenterY                                int
}

func NewRoom(startX, startY, width, heigth int) *Room {
	return &Room{
		StartPointFromX: startX,
		StartPointFromY: startY,
		Width:           width,
		Heigth:          heigth,
		CenterX:         startX + width/2,
		CenterY:         startY + heigth/2,
	}
}

// Рюкзак;
type Backpack struct {
	Items []interface{}
}

/*
Предмет:

	тип,
	подтип,
	здоровье (количество единиц повышения, для еды),
	максимальный уровень здоровья (количество единиц повышения, для свитков и эликсиров, вместе с этим повышается и сам уровень здоровья),
	ловкость (количество единиц повышения, для свитков и эликсиров),
	сила (количество единиц повышения, для свитков, эликсиров и оружия),
	стоимость (для сокровищ). */

type Item struct {
	ForRating     bool
	Coordinates   interface{}
	Type          interface{}
	SubType       interface{}
	HealingPoint  int
	PointForMAxHp int
	DexteryPoint  int
	StrengthPoint int
	Cost          int
}

// weapon

type Weapon struct {
	Name   string
	Damage int
}

func initWeapon() *Weapon {
	return &Weapon{
		Name:   "",
		Damage: 0,
	}
}
