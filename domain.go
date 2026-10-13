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
	Statistic                           *Statistic
	Backpack                            *Backpack
}

type Statistic struct {
	RunNum,
	Gold,
	Lvl,
	Enemy,
	Food,
	Flask,
	Scroll,
	Hit,
	Miss,
	Cell int
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
	StateStatistic
	StateFlaskList
	StateScrollList
	StateWeaponList
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
	Gold    int
	Food    int
	Flasks  []Flask
	Scrolls []Scroll
	Weapons []Weapon
}

func InitBackpack() *Backpack {
	return &Backpack{}
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

type Flask struct {
	Name     string
	NewMaxHp int
	StrAdd   int
	DextAdd  int
}

const (
	S           string = "Силы"
	D           string = "Ловкости"
	H           string = "Здоровья"
	FL_SCR_NAME string = "Повышения "
)

func (s *Session) InitFlask() *Flask {
	num := rand.Intn(3)
	var f Flask
	switch num {
	case 0:
		f.NewMaxHp = rand.Intn(3) + 1
		f.Name = FL_SCR_NAME + H
	case 1:
		f.StrAdd = rand.Intn(3) + 1
		f.Name = FL_SCR_NAME + S
	case 2:
		f.DextAdd = rand.Intn(3) + 1
		f.Name = FL_SCR_NAME + D

	}
	return &f
}
func (s *Session) UseFlask(i int) {
	if len(s.Backpack.Flasks) == 0 {
		return
	}
	s.Player.MaxHp += s.Backpack.Flasks[i].NewMaxHp
	if s.Backpack.Flasks[i].NewMaxHp > 0 {
		s.Player.CurrentHp = s.Player.MaxHp
	}
	s.Player.Strength += s.Backpack.Flasks[i].StrAdd
	s.Player.Dextery += s.Backpack.Flasks[i].DextAdd

	s.Backpack.Flasks[i] = s.Backpack.Flasks[len(s.Backpack.Flasks)-1]
	s.Backpack.Flasks = s.Backpack.Flasks[:len(s.Backpack.Flasks)-1]
}

type Scroll struct {
	Name     string
	NewMaxHp int
	StrAdd   int
	DextAdd  int
}

func (s *Session) InitScroll() *Scroll {
	num := rand.Intn(3)
	var f Scroll
	switch num {
	case 0:
		f.NewMaxHp = rand.Intn(3) + 1
		f.Name = FL_SCR_NAME + H
	case 1:
		f.StrAdd = rand.Intn(3) + 1
		f.Name = FL_SCR_NAME + S
	case 2:
		f.DextAdd = rand.Intn(3) + 1
		f.Name = FL_SCR_NAME + D

	}
	return &f
}
func (s *Session) UseScroll(i int) {
	if len(s.Backpack.Scrolls) == 0 {
		return
	}
	s.Player.MaxHp += s.Backpack.Scrolls[i].NewMaxHp
	if s.Backpack.Scrolls[i].NewMaxHp > 0 {
		s.Player.CurrentHp = s.Player.MaxHp
	}
	s.Player.Strength += s.Backpack.Scrolls[i].StrAdd
	s.Player.Dextery += s.Backpack.Scrolls[i].DextAdd

	s.Backpack.Scrolls[i] = s.Backpack.Scrolls[len(s.Backpack.Scrolls)-1]
	s.Backpack.Scrolls = s.Backpack.Scrolls[:len(s.Backpack.Scrolls)-1]
}

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
