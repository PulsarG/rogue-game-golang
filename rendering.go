package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"

	"github.com/gdamore/tcell/v2"
)

/*
	    При нахождении в непосредственной близости с комнатой со стороны коридора туман войны рассеивается только на области прямой видимости (применяется алгоритм Ray Casting и алгоритм Брезенхэма для определения видимой области).
	Пример реализации рендеринга уровней представлен в папке code-samples.

	Любое использование чего-либо из рюкзака должно приводить к печати списка предметов этого типа на экран с вопросом игроку, что нужно выбрать (1–9).
	При выборе оружия также должна иметься возможность убрать оружие из рук, не выбрасывая из инвентаря (соответственно, для оружия выбор будет 0–9).

Статистика

*/

var s *Session
var isRun bool
var GameLvl int

func RunView() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()
	err = screen.Init()
	if err != nil {
		log.Fatal(err)
	}
	GameLvl = FIRST_START_GAME_LVL
	s = initNewGame(NEW_GAME_HP, INIT_NEW_BACKPACK)
	isRun = true
	for isRun {
		screen.Clear()
		switch s.CurrentState {
		case StateStartNewGame:
			GameLvl = FIRST_START_GAME_LVL // need 1 new game lvl
			s = initNewGame(NEW_GAME_HP, INIT_NEW_BACKPACK)
			s.CurrentState = StatePlay
		case StateNextLvl:
			s = initNewGame(s.Player.CurrentHp, INIT_SAVE_BACKPACK)
			s.CurrentState = StatePlay
		case StateGameOver:
			DrawString(
				screen,
				10,
				18,
				"Игрок погиб! Жми 1",
			)
		}

		switch s.CurrentState {
		case StateMainMenu:
			runMenu(screen)
		case StateInventory:
			runIventory(screen)
		case StateGameMenu:
			runGameMenu(screen)
		case StatePlay:
			runFight()
			runGame(screen)
		case StateStatistic:
			runStatistic(screen)
		case StateFlaskList:
			runFlaskList(screen)
		case StateScrollList:
			runScrollList(screen)
		case StateWeaponList:
			runWeaponList(screen)
		}
		screen.Show()

		event := screen.PollEvent()
		checkInput(&event, s.Player, &s.Mapp)
	}
}

func initNewGame(hp, isNewBackPack int) *Session {

	ns := Session{}
	ns.GenMap()
	ns.Player = PlayerInit(&ns.Rooms[ns.StartRoomIdx], hp)
	ns.CurrentLvl = GameLvl
	ns.EnemysListInit()
	//s.CurrentState = StateMainMenu
	ns.EndRoom = ns.EndRoomInit()
	ns.GoldInit()
	ns.Statistic = statisticInit()
	if isNewBackPack == 0 {
		ns.Backpack = InitBackpack()
	} else {
		ns.Backpack = s.Backpack
		ns.Player.Weapon = s.Player.Weapon
	}
	ns.FlasksInit()
	ns.ScrollsInit()
	ns.WeaponsInit()
	ns.PlayerScanFov()
	return &ns
}

func statisticInit() *Statistic {
	return &Statistic{
		RunNum: s.readNumRun(),
	}
}

func runMenu(screen tcell.Screen) {
	DrawString(screen, 3, 3, "1) New Game")
	DrawString(screen, 3, 4, "2) Continue")
	DrawString(screen, 3, 5, "3) Leaderboard")
	DrawString(screen, 3, 7, "4, q) Exit")
}

func runStatistic(screen tcell.Screen) {
	st := s.readStatistic()
	sort.Slice(st, func(i, j int) bool {
		return st[i].Gold > st[j].Gold
	})
	DrawString(screen, 3, 1, "1) В главное меню")
	DrawString(screen, 3, 2, "2) Очистить")
	for i, v := range st {
		DrawString(screen, 3, i+4,
			fmt.Sprint("Сокровища: ", v.Gold, " Run ", v.RunNum, " Lvl ", v.Lvl, " Kill ", v.Hit, " miss ", v.Miss, " Cell ", v.Cell, " Food, Flask, Scroll ", v.Food, v.Flask, v.Scroll))
	}
}

func runIventory(screen tcell.Screen) {

	DrawString(screen, 3, 1, "IVENTORY")
	DrawString(screen, 3, 3, fmt.Sprintf("Сокровища: %d", s.Statistic.Gold))
	DrawString(screen, 3, 4, fmt.Sprintf("Еда: %d", s.Backpack.Food))
	if len(s.Backpack.Flasks) > 0 {
		for i, v := range s.Backpack.Flasks {
			DrawString(screen, 3, i+6, v.Name)
		}
	}
	if len(s.Backpack.Scrolls) > 0 {
		for i, v := range s.Backpack.Scrolls {
			DrawString(screen, 3, i+15, v.Name)
		}
	}
	if len(s.Backpack.Weapons) > 0 {
		for i, v := range s.Backpack.Weapons {
			DrawString(screen, 3, i+30, v.Name)
		}
	}
}
func runFlaskList(screen tcell.Screen) {

	DrawString(screen, 3, 1, "Эликсиры")
	DrawString(screen, 3, 2, "Выбор пономеру")
	if len(s.Backpack.Flasks) > 0 {
		for i, v := range s.Backpack.Flasks {
			DrawString(screen, 3, i+4, fmt.Sprint(i+1, ") ", v.Name))
		}
	}
}
func runScrollList(screen tcell.Screen) {

	DrawString(screen, 3, 1, "Свитки")
	DrawString(screen, 3, 2, "Выбор пономеру")
	if len(s.Backpack.Scrolls) > 0 {
		for i, v := range s.Backpack.Scrolls {
			DrawString(screen, 3, i+4, fmt.Sprint(i+1, ") ", v.Name))
		}
	}
}
func runWeaponList(screen tcell.Screen) {

	DrawString(screen, 3, 1, "Оружие")
	DrawString(screen, 3, 2, "Выбор по номеру")
	DrawString(screen, 3, 3, "0) Сбросить оружие")
	if len(s.Backpack.Weapons) > 0 {
		for i, v := range s.Backpack.Weapons {
			DrawString(screen, 3, i+5, fmt.Sprint(i+1, ") ", v.Name))
		}
	}
}
func runGameMenu(screen tcell.Screen) {
	DrawString(screen, 3, 3, "1, m) Continue")
	DrawString(screen, 3, 4, "2) Save and Exit")
}

func runGame(screen tcell.Screen) {

	for i := range 100 {
		for j := range 40 {
			if s.StatusMap[i][j] == 1 {
				screen.SetContent(i, j, s.Mapp[i][j], nil, tcell.StyleDefault)
			}
		}
	}
	DrawString(
		screen,
		102,
		1,
		fmt.Sprintf("%s", s.PlayerStatusDoing),
	)
	DrawString(
		screen,
		102,
		2,
		fmt.Sprintf("%s", s.EnemyStatusDoing),
	)
	DrawString(
		screen,
		102,
		10,
		fmt.Sprint("Lvl: ", s.CurrentLvl),
	)
	DrawString(
		screen,
		102,
		11,
		fmt.Sprint("HP: ", s.Player.CurrentHp),
	)
	DrawString(
		screen,
		102,
		12,
		fmt.Sprint("Max HP: ", s.Player.MaxHp),
	)
	DrawString(
		screen,
		102,
		13,
		fmt.Sprint("Str: ", s.Player.Strength),
	)
	DrawString(
		screen,
		102,
		14,
		fmt.Sprint("Dex: ", s.Player.Dextery),
	)
	DrawString(
		screen,
		102,
		16,
		"H - выбор оружия",
	)
	DrawString(
		screen,
		102,
		17,
		"J - использовать еду",
	)
	DrawString(
		screen,
		102,
		18,
		"K - использовать эликсир",
	)
	DrawString(
		screen,
		102,
		19,
		"E - использовать свиток",
	)
	DrawString(
		screen,
		102,
		21,
		fmt.Sprint("Оружие: ", s.Player.Weapon.Name),
	)
	Draw(s.Player.X, s.Player.Y, s.Player.Sprite, screen, &s.Player.ColorSprite)
	for _, e := range s.Enemys {
		Draw(e.X, e.Y, e.Sprite, screen, &e.ColorSprite)
	}
	//	Draw(s.EndRoom.X, s.EndRoom.Y, s.EndRoom.Sprite, screen, &s.EndRoom.ColorSprite)
}

func runFight() {
	for i, e := range s.Enemys {
		if e.CheckFight(s.Player) {
			s.PlayerStatusDoing = fmt.Sprintf("Игрок убил %s", e.Name)
			s.Enemys[i] = s.Enemys[len(s.Enemys)-1]
			s.Enemys = s.Enemys[:len(s.Enemys)-1]
			s.Statistic.Enemy += 1
			//s.Enemys = append(s.Enemys[:i], s.Enemys[i+1:]...)
		}
	}
}

/*func openMenu(screen tcell.Screen) {
	DrawString(screen, 10, 10, fmt.Sprintf("Inventory: %d", "!"))
}*/

func (s *Session) checkFightX(v int, p *Player) bool {
	for i, e := range s.Enemys {
		if p.X+v == e.X && p.Y == e.Y {
			s.Player.FightWith = e
			s.giveHit(s.Enemys[i])
			e.IsTakeSmash = true
			return false
		}
	}
	return true
}
func (s *Session) checkFightY(v int, p *Player) bool {
	for i, e := range s.Enemys {
		if p.X == e.X && p.Y+v == e.Y {
			s.Player.FightWith = e
			s.giveHit(s.Enemys[i])
			return false
		}
	}
	return true
}

func (s *Session) giveHit(e *Enemy) {
	if rand.Intn(7)+s.Player.Dextery < e.Dextery {
		s.PlayerStatusDoing = "Игрок промахнулся"
		//return
	} else {
		damage := rand.Intn(s.Player.Strength+s.Player.Weapon.Damage) + 1
		e.Hp -= damage
		s.Statistic.Hit += 1
		s.PlayerStatusDoing = fmt.Sprintf("Игрок нанес %d урона", damage)
	}
	e.IsTakeSmash = true
}
func (s *Session) answerHit(e *Enemy) {
	if e.Hp <= 0 {
		if rand.Intn(2) == 0 {
			e.Drop(&s.Mapp)
		}
		return
	}
	ch := rand.Intn(7)
	if ch+e.Dextery < s.Player.Dextery {
		s.EnemyStatusDoing = "Враг промахнулся"
		//return
	} else {
		damage := rand.Intn(e.Strength) + 1
		s.Player.CurrentHp -= damage
		if s.Player.CurrentHp <= 0 {
			s.writeStatistic()
			s.writeNumRun()
			s.CurrentState = StateGameOver
			return
		}
		s.Statistic.Miss += 1
		s.EnemyStatusDoing = fmt.Sprintf("Враг нанес %d урона", damage)
	}
}

func DrawString(screen tcell.Screen, x, y int, msg string) {
	for i, char := range msg {
		screen.SetContent(x+i, y, char, nil, tcell.StyleDefault)
	}
}

func CheckRoof(vector int, mapp [100][40]rune, player *Player) bool {
	switch mapp[player.X][player.Y+vector] {
	case '.':
		return true
	case '#':
		return true
	case '+':
		return true
	case '%':
		GameLvl = s.CurrentLvl + 1
		s.CurrentState = StateNextLvl
		return true
	case 'G':
		s.Statistic.Gold += 1
		s.Mapp[player.X][player.Y+vector] = '.'
		return true
	case 'F':
		if len(s.Backpack.Flasks) >= 9 {
			s.PlayerStatusDoing = "Нельзя больше взять"
			return false
		}
		s.Backpack.Flasks = append(s.Backpack.Flasks, *s.InitFlask())
		s.Mapp[player.X][player.Y+vector] = '.'
		return true
	case 'E':
		if len(s.Backpack.Scrolls) >= 9 {
			s.PlayerStatusDoing = "Нельзя больше взять"
			return false
		}
		s.Backpack.Scrolls = append(s.Backpack.Scrolls, *s.InitScroll())
		s.Mapp[player.X][player.Y+vector] = '.'
		return true
	case 'W':
		if len(s.Backpack.Weapons) >= 9 {
			s.PlayerStatusDoing = "Нельзя больше взять"
			return false
		}
		s.Backpack.Weapons = append(s.Backpack.Weapons, *s.InitOneWeapon())
		s.Mapp[player.X][player.Y+vector] = '.'
		return true
	default:
		return false
	}
}

func CheckWall(vector int, mapp [100][40]rune, player *Player) bool {
	switch mapp[player.X+vector][player.Y] {
	case '.':
		return true
	case '#':
		return true
	case '+':
		return true
	case '%':
		GameLvl = s.CurrentLvl + 1
		s.CurrentState = StateNextLvl
		return true
	case 'G':
		s.Statistic.Gold += 1
		s.Mapp[player.X+vector][player.Y] = '.'
		return true
	case 'F':
		if len(s.Backpack.Flasks) >= 9 {
			s.PlayerStatusDoing = "Нельзя больше взять"
			return false
		}
		s.Backpack.Flasks = append(s.Backpack.Flasks, *s.InitFlask())
		s.Mapp[player.X+vector][player.Y] = '.'
		return true
	case 'E':
		if len(s.Backpack.Scrolls) >= 9 {
			s.PlayerStatusDoing = "Нельзя больше взять"
			return false
		}
		s.Backpack.Scrolls = append(s.Backpack.Scrolls, *s.InitScroll())
		s.Mapp[player.X+vector][player.Y] = '.'
		return true
	case 'W':
		if len(s.Backpack.Weapons) >= 9 {
			s.PlayerStatusDoing = "Нельзя больше взять"
			return false
		}
		s.Backpack.Weapons = append(s.Backpack.Weapons, *s.InitOneWeapon())
		s.Mapp[player.X+vector][player.Y] = '.'
		return true
	default:
		return false
	}
}
