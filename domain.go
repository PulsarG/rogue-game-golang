package main

import (
	"fmt"
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

// Игровая сессия;
type Session struct {
	CurrentLvl                          int
	Mapp                                [100][40]rune
	StatusMap                           [100][40]int
	RoomId                              [100][40]int
	Player                              *Player
	Rooms                               [9]Room
	StartRoomIdx, EndRoomIdx            int
	Enemys                              []*Enemy
	PlayerStatusDoing, EnemyStatusDoing string
	CurrentState                        GameState
	EndRoom                             *EndRoom
	Statistic                           *Statistic
	Backpack                            *Backpack
	IsGhostfight                        bool
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

type PointStatus int

const (
	Unvisible int = 0
	Visible   int = 1
	Remember  int = 2
)

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

func (s *Session) PlayerScanFov(status int) {
	startX := s.Player.X - RADIUS_SCAN
	if startX < 0 {
		startX = 0
	}

	endX := s.Player.X + RADIUS_SCAN
	if endX >= 100 {
		endX = 99
	}

	startY := s.Player.Y - RADIUS_SCAN
	if startY < 0 {
		startY = 0
	}

	endY := s.Player.Y + RADIUS_SCAN
	if endY >= 40 {
		endY = 39
	}

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			s.StatusMap[x][y] = status
		}
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

type Flask struct {
	Name     string
	NewMaxHp int
	StrAdd   int
	DextAdd  int
}

const (
	S                  string = "Силы"
	D                  string = "Ловкости"
	H                  string = "Здоровья"
	FL_SCR_NAME        string = "Повышения "
	WEAPON_AXE_NAME    string = "Топор "
	WEAPON_SPIROW_NAME string = "Копье "
	WEAPON_MACE_NAME   string = "Меч "
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
	if i < 0 || i >= len(s.Backpack.Flasks) {
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
	if i < 0 || i >= len(s.Backpack.Scrolls) {
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

func InitWeapon() *Weapon {
	return &Weapon{
		Name:   "",
		Damage: 0,
	}
}
func (s *Session) InitOneWeapon() *Weapon {
	num := rand.Intn(3)
	var f Weapon
	switch num {
	case 0:
		f.Damage = rand.Intn(3) + 1 + s.CurrentLvl/3
		f.Name = fmt.Sprint(WEAPON_MACE_NAME, f.Damage, " Силы")
	case 1:
		f.Damage = rand.Intn(3) + 1 + s.CurrentLvl/3
		f.Name = fmt.Sprint(WEAPON_SPIROW_NAME, f.Damage, " Силы")
	case 2:
		f.Damage = rand.Intn(3) + 1 + s.CurrentLvl/3
		f.Name = fmt.Sprint(WEAPON_AXE_NAME, f.Damage, " Силы")

	}
	return &f
}
func (s *Session) DropWeapon() {
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, d := range dirs {
		nx, ny := s.Player.X+d[0], s.Player.Y+d[1]
		if nx >= 0 && nx < 100 && ny < 40 && ny >= 0 && s.Mapp[nx][ny] == '.' {
			s.Mapp[nx][ny] = 'W'
			return
		}
	}
	//x := s.Player.X + 1
	//y := s.Player.Y + 1
	//s.Mapp[x][y] = 'W'
}
func (s *Session) UseWeapon(i int) {
	if len(s.Backpack.Weapons) == 0 {
		return
	}
	if i < 0 || i >= len(s.Backpack.Weapons) {
		return
	}

	s.Player.Weapon = &s.Backpack.Weapons[i]

	s.Backpack.Weapons[i] = s.Backpack.Weapons[len(s.Backpack.Weapons)-1]
	s.Backpack.Weapons = s.Backpack.Weapons[:len(s.Backpack.Weapons)-1]
}
