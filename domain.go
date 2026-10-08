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
	Name        string
	Sprite      rune
	ColorSprite tcell.Color
	Difficulty  int
	Hp          int
	Dextery     int
	Strength    int
	Agro        int
	IsTakeSmash bool
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
		e.Hp = 3
		e.Dextery = 1
		e.Strength = 1
	case 1:
		e.Sprite = 'z'
		e.ColorSprite = tcell.ColorGreen
		e.Agro = 2
		e.Hp = 3
		e.Dextery = 1
		e.Name = "Зоиби"
		e.Strength = 5
	case 2:
		e.Sprite = 's'
		e.ColorSprite = tcell.ColorBlue
		e.Agro = 2
		e.Hp = 3
		e.Dextery = 7
		e.Name = "Змея"
		e.Strength = 1
	case 3:
		e.Sprite = 'v'
		e.ColorSprite = tcell.ColorYellow
		e.Agro = 2
		e.Hp = 3
		e.Dextery = 1
		e.Strength = 1
	default:
		e.Sprite = 'x'
		e.ColorSprite = tcell.ColorWhite
		e.Agro = 2
		e.Hp = 3
		e.Dextery = 1
		e.Strength = 1
	}
}

func (e *Enemy) CheckFight(p *Player) bool {
	res := false
	// check
	if e.checkAgro(p) {
		//	s.StatusDoing = fmt.Sprintf("In agro %c", e.Sprite)
		e.moveToPlayer(p)
	}
	//smash
	if e.IsTakeSmash {
		s.answerHit(e)
		e.IsTakeSmash = false
	}
	//smashPLayer()
	//dea
	if e.Hp <= 0 {
		s.Mapp[e.X][e.Y] = '.'
		res = true
	}
	// remove
	//drop
	return res
}

func (e *Enemy) checkAgro(p *Player) bool {
	dx := abs(e.X - p.X)
	dy := abs(e.Y - p.Y)
	distance := max(dx, dy)
	if distance <= e.Agro {
		return true
	} else {
		return false
	}
}

func (e *Enemy) moveToPlayer(p *Player) {
	vectorX := p.X - e.X
	vectorY := p.Y - e.Y
	if rand.Intn(2) == 0 {
		if vectorX != 0 {
			if vectorX > 0 {
				if e.checkRuneX(1, p) {
					e.X += 1
					s.Mapp[e.X][e.Y] = e.Sprite
					s.Mapp[e.X-1][e.Y] = '.'
				}
			} else {
				if e.checkRuneX(-1, p) {
					e.X -= 1
					s.Mapp[e.X][e.Y] = e.Sprite
					s.Mapp[e.X+1][e.Y] = '.'
				}
			}
		} else if vectorY != 0 {
			if vectorY > 0 {
				if e.checkRuneY(1, p) {
					e.Y += 1
					s.Mapp[e.X][e.Y] = e.Sprite
					s.Mapp[e.X][e.Y-1] = '.'
				}
			} else {
				if e.checkRuneY(-1, p) {
					e.Y -= 1
					s.Mapp[e.X][e.Y] = e.Sprite
					s.Mapp[e.X][e.Y+1] = '.'
				}
			}
		}
	} else {
		if vectorY != 0 {
			if vectorY > 0 {
				if e.checkRuneY(1, p) {
					e.Y += 1
					s.Mapp[e.X][e.Y] = e.Sprite
					s.Mapp[e.X][e.Y-1] = '.'
				}
			} else {
				if e.checkRuneY(1, p) {
					e.Y -= 1
					s.Mapp[e.X][e.Y] = e.Sprite
					s.Mapp[e.X][e.Y+1] = '.'
				}
			}
		} else if vectorX != 0 {
			if vectorX > 0 {
				if e.checkRuneX(1, p) {
					e.X += 1
					s.Mapp[e.X][e.Y] = e.Sprite
					s.Mapp[e.X-1][e.Y] = '.'
				}
			} else {
				if e.checkRuneX(1, p) {
					e.X -= 1
					s.Mapp[e.X][e.Y] = e.Sprite
					s.Mapp[e.X+1][e.Y] = '.'
				}
			}
		}
	}
}

func (e *Enemy) checkRuneX(v int, p *Player) bool {
	if e.X+v == p.X && e.Y == p.Y {
		return false
	} else if s.Mapp[e.X+v][e.Y] != '+' && s.Mapp[e.X+v][e.Y] != '#' && s.Mapp[e.X+v][e.Y] != '.' {
		return false
	} else {
		return true
	}
}
func (e *Enemy) checkRuneY(v int, p *Player) bool {
	if e.X == p.X && e.Y+v == p.Y {
		return false
	} else if s.Mapp[e.X][e.Y+v] != '+' && s.Mapp[e.X][e.Y+v] != '#' && s.Mapp[e.X][e.Y+v] != '.' {
		return false
	} else {
		return true
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
