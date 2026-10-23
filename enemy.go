package main

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

type Enemy struct {
	X, Y           int
	Type           int
	Name           string
	Sprite         rune
	ColorSprite    tcell.Color
	Difficulty     int
	Hp             int
	Dextery        int
	Strength       int
	Agro           int
	IsTakeSmash    bool
	IsFirstHit     bool
	IsSecondHitOgr bool
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
			e.Type = rand.Intn(5)
			e.SelectSprite(s.CurrentLvl)
			break
		}
	}
	//	s.Mapp[e.X][e.Y] = e.Sprite
	return &e
}

func (e *Enemy) SelectSprite(l int) {
	switch e.Type {
	case 0:
		e.Sprite = 'g'
		e.ColorSprite = tcell.ColorWheat
		e.Agro = 1
		e.Hp = 1 + (l+1)*2
		e.Dextery = 3 + (l+1)*2
		e.Strength = 1 + (l+1)*2
	case 1:
		e.Sprite = 'z'
		e.ColorSprite = tcell.ColorGreen
		e.Agro = 2
		e.Hp = 3 + (l+1)*2
		e.Dextery = 1 + (l+1)*2
		e.Name = "Зоиби"
		e.Strength = 2 + (l+1)*2
	case 2:
		e.Sprite = 's'
		e.ColorSprite = tcell.ColorWhite
		e.Agro = 3
		e.Hp = 2 + (l+1)*2
		e.Dextery = 4 + (l+1)*2
		e.Name = "Змей-маг"
		e.Strength = 1 + (l+1)*2
	case 3:
		e.Sprite = 'v'
		e.ColorSprite = tcell.ColorRed
		e.Name = "Вампир"
		e.Agro = 3
		e.Hp = 3 + (l+1)*2
		e.Dextery = 3 + (l+1)*2
		e.Strength = 2 + (l+1)*2
	default:
		e.Sprite = 'O'
		e.ColorSprite = tcell.ColorYellow
		e.Name = "Огр"
		e.Agro = 2
		e.Hp = 4 + (l+1)*2
		e.Dextery = 1 + (l+1)*2
		e.Strength = 4 + (l+1)*2
	}
}

func (e *Enemy) CheckFight(p *Player) bool {
	res := false
	// check
	if e.checkAgro(p) {
		//	s.StatusDoing = fmt.Sprintf("In agro %c", e.Sprite)
		if e.Sprite == 'g' {
			s.IsGhostfight = true
		}
		e.moveToPlayer(p, &s.Mapp)
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

func (e *Enemy) moveToPlayer(p *Player, mapp *[100][40]rune) {
	vectorX := p.X - e.X
	vectorY := p.Y - e.Y
	step := 1
	if e.Sprite == 'O' {
		step = 2
	}
	if rand.Intn(2) == 0 {
		if vectorX != 0 {
			if vectorX > 0 {
				if e.checkRuneX(1, p) {
					e.X += step
					mapp[e.X][e.Y] = e.Sprite
					mapp[e.X-step][e.Y] = '.'
				}
			} else {
				if e.checkRuneX(-1, p) {
					e.X -= step
					mapp[e.X][e.Y] = e.Sprite
					mapp[e.X+step][e.Y] = '.'
				}
			}
		} else if vectorY != 0 {
			if vectorY > 0 {
				if e.checkRuneY(1, p) {
					e.Y += step
					mapp[e.X][e.Y] = e.Sprite
					mapp[e.X][e.Y-step] = '.'
				}
			} else {
				if e.checkRuneY(-1, p) {
					e.Y -= step
					mapp[e.X][e.Y] = e.Sprite
					mapp[e.X][e.Y+step] = '.'
				}
			}
		}
	} else {
		if vectorY != 0 {
			if vectorY > 0 {
				if e.checkRuneY(1, p) {
					e.Y += step
					mapp[e.X][e.Y] = e.Sprite
					mapp[e.X][e.Y-step] = '.'
				}
			} else {
				if e.checkRuneY(1, p) {
					e.Y -= step
					mapp[e.X][e.Y] = e.Sprite
					mapp[e.X][e.Y+step] = '.'
				}
			}
		} else if vectorX != 0 {
			if vectorX > 0 {
				if e.checkRuneX(1, p) {
					e.X += step
					mapp[e.X][e.Y] = e.Sprite
					mapp[e.X-step][e.Y] = '.'
				}
			} else {
				if e.checkRuneX(1, p) {
					e.X -= step
					mapp[e.X][e.Y] = e.Sprite
					mapp[e.X+step][e.Y] = '.'
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

func (e *Enemy) Drop(m *[100][40]rune) {
	x := e.X + 1
	y := e.Y + 1
	r := rand.Intn(4)
	switch r {
	case 0:
		m[x][y] = 'W'
	case 1:
		m[x][y] = 'E'
	case 2:
		m[x][y] = 'F'
	case 3:
		m[x][y] = 'J'
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
func (s *Session) GhostIdle() {
	for _, e := range s.Enemys {
		if e.Sprite == 'g' {
			for try := 0; try < 10; try++ {
				x := deltaVector()
				y := deltaVector()
				if e.X+x >= 100 || e.Y+y >= 40 || e.X+x < 0 || e.Y+y < 0 {
					continue
				}
				if s.Mapp[e.X+x][e.Y+y] == '.' {
					s.Mapp[e.X][e.Y] = '.'
					e.X += x
					e.Y += y
					break
				} else {
					continue
				}
			}
		}
	}
}

func deltaVector() int {
	x := rand.Intn(5)
	y := rand.Intn(2)
	if y == 0 {
		x *= -1
	}
	return x
}
