package main

import (
	"fmt"
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
		e.Agro = 2
	case 1:
		e.Sprite = 'z'
		e.ColorSprite = tcell.ColorGreen
		e.Agro = 2
	case 2:
		e.Sprite = 's'
		e.ColorSprite = tcell.ColorBlue
		e.Agro = 2
	case 3:
		e.Sprite = 'v'
		e.ColorSprite = tcell.ColorYellow
		e.Agro = 2
	default:
		e.Sprite = 'x'
		e.ColorSprite = tcell.ColorWhite
		e.Agro = 2
	}
}

func (e *Enemy) CheckFight(p *Player) bool {
	// check
	if e.checkAgro(p) {
		s.StatusDoing = fmt.Sprintf("In agro %c", e.Sprite)
		e.moveToPlayer(p)
	}
	// move
	//moveToPlayer()
	//smash
	//smashPLayer()
	//dea
	// remove
	//drop
	if e.X == p.X && e.Y == p.Y {
		s.StatusDoing = fmt.Sprintf("!!! %d", 1)
		return true
	} else {
		return false
	}
}

func (e *Enemy) checkAgro(p *Player) bool {
	dx := abs(e.X - p.X)
	dy := abs(e.Y - p.Y)
	distance := max(dx, dy)
	//if dy > dx {
	//	distance = dy
	//	}

	if distance <= e.Agro {
		return true
	} else {
		return false
	}
}

func (e *Enemy) moveToPlayer(p *Player) {
	vectorX := p.X - e.X
	vectorY := p.Y - e.Y
	if vectorX != 0 {
		if vectorX > 0 {
			e.X += 1
		} else {
			e.X -= 1
		}
	}
	if vectorY != 0 {
		if vectorY > 0 {
			e.Y += 1
		} else {
			e.Y -= 1
		}
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
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
