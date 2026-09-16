package main

//Игровая сессия;
type Session struct {
	CurrentLvl               int
	Mapp                     [100][40]rune
	Player                   *Player
	Rooms                    [9]Room
	StartRoomIdx, EndRoomIdx int
	// возможно данные для статистики
}

//Уровень;
type Level struct {
	CountEnemy int
}

//Комната;
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

//Коридор;
type Сorridor struct {
	FirstDoor   interface{}
	SecondDoor  interface{}
	Coordinates interface{}
}

/* Персонаж:
   максимальный уровень здоровья,
   здоровье,
   ловкость,
   сила,
   текущее оружие; */
type Player struct {
	X, Y      int
	Sprite    rune
	MaxHp     int
	CurrentHp int
	Dextery   int
	Strength  int
	Weapon    interface{}
}

func PlayerInit(r *Room) *Player {
	return &Player{
		X:      r.CenterX,
		Y:      r.CenterY,
		Sprite: '@',
	}
}

func (p *Player) SetCoordinat(x, y int) {
	p.X = x
	p.Y = y
}

//Рюкзак;
type Backpack struct {
	Items []interface{}
}

/* Противник:
   тип,
   здоровье,
   ловкость,
   сила,
   враждебность; */
type Enemy struct {
	Coordinates interface{}
	Type        string
	Difficulty  int
	Hp          int
	Dextery     int
	Strength    int
	Agro        int
}

/* Предмет:
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
