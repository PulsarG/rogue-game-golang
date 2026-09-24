package main

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

// Игровая сессия;
type Session struct {
	CurrentLvl               int
	Mapp                     [100][40]rune
	Player                   *Player
	Rooms                    [9]Room
	StartRoomIdx, EndRoomIdx int
	Enemys                   []*Enemy
	StatusDoing              string
	// возможно данные для статистики
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

/*
Персонаж:

	максимальный уровень здоровья,
	здоровье,
	ловкость,
	сила,
	текущее оружие;
*/
type Player struct {
	X, Y        int
	Sprite      rune
	ColorSprite tcell.Color
	MaxHp       int
	CurrentHp   int
	Dextery     int
	Strength    int
	Weapon      interface{}
	FightWith   *Enemy
}

func PlayerInit(r *Room) *Player {
	return &Player{
		X:           r.CenterX,
		Y:           r.CenterY,
		Sprite:      '@',
		ColorSprite: tcell.ColorDefault,
	}
}

func (p *Player) SetCoordinat(x, y int) {
	p.X = x
	p.Y = y
}

// Рюкзак;
type Backpack struct {
	Items []interface{}
}

/*
Противник:

	тип,
	здоровье,
	ловкость,
	сила,
	враждебность;
*/
type Enemy struct {
	X, Y        int
	Type        int
	Sprite      rune
	ColorSprite tcell.Color
	Difficulty  int
	Hp          int
	Dextery     int
	Strength    int
	Agro        int
}

func (s *Session) EnemysListInit() {
	countEnemys := 3 + rand.Intn(s.CurrentLvl)
	enemys := make([]*Enemy, 0, countEnemys)
	for _ = range countEnemys {
		enemys = append(enemys, s.enemyOneInit())
	}
	s.Enemys = enemys
}

func (s *Session) enemyOneInit() *Enemy {
	e := Enemy{}
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
					nicePoint = true
				}
			}
			e.Type = rand.Intn(4)
			e.SelectSprite()
			break
		}
	}
	return &e
}

func (e *Enemy) SelectSprite() {
	switch e.Type {
	case 0:
		e.Sprite = 'a'
		e.ColorSprite = tcell.ColorRed
	case 1:
		e.Sprite = 'z'
		e.ColorSprite = tcell.ColorGreen
	case 2:
		e.Sprite = 's'
		e.ColorSprite = tcell.ColorBlue
	case 3:
		e.Sprite = 'v'
		e.ColorSprite = tcell.ColorYellow
	default:
		e.Sprite = 'x'
		e.ColorSprite = tcell.ColorWhite
	}
}

func (e *Enemy) CheckFight(p *Player) bool {
	if e.X+e.Agro == p.X && e.Y+e.Agro == p.Y {
		return true
	}
	return false
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
