package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

/*
	    При нахождении в непосредственной близости с комнатой со стороны коридора туман войны рассеивается только на области прямой видимости (применяется алгоритм Ray Casting и алгоритм Брезенхэма для определения видимой области).
	Пример реализации рендеринга уровней представлен в папке code-samples.

	    Передвижение при помощи клавиш WASD.
	    Применение оружия из рюкзака при помощи кнопки h.
	    Применение еды из рюкзака при помощи кнопки j.
	    Применение эликсира из рюкзака при помощи кнопки k.
	    Применение свитка из рюкзака при помощи e.
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
	GameLvl = 1
	s = initNewGame(10)
	isRun = true
	for isRun {
		screen.Clear()
		switch s.CurrentState {
		case StateStartNewGame:
			GameLvl = 20
			s = initNewGame(10)
			s.Statistic.RunNum = s.readNumRun()
			s.CurrentState = StatePlay
		case StateNextLvl:
			s = initNewGame(s.Player.CurrentHp)
			s.CurrentState = StatePlay
		case StateGameOver:
			DrawString(
				screen,
				10,
				18,
				fmt.Sprint("Игрок погиб! Жми 1"),
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
		}
		screen.Show()

		event := screen.PollEvent()
		checkInput(&event, s.Player, &s.Mapp)
	}
}

func initNewGame(hp int) *Session {

	s := Session{}
	s.GenMap()
	s.Player = PlayerInit(&s.Rooms[s.StartRoomIdx], hp)
	s.CurrentLvl = GameLvl
	s.EnemysListInit()
	//s.CurrentState = StateMainMenu
	s.EndRoom = s.EndRoomInit()
	s.Statistic = statisticInit()
	return &s
}

func statisticInit() *Statistic {
	return &Statistic{}
}

func runMenu(screen tcell.Screen) {
	DrawString(screen, 3, 3, "1) New Game")
	DrawString(screen, 3, 4, "2) Continue")
	DrawString(screen, 3, 5, "3) Leaderboard")
	DrawString(screen, 3, 7, "4, q) Exit")
}

func runStatistic(screen tcell.Screen) {
	st := s.readStatistic()
	DrawString(screen, 3, 1, "1) В главное меню")
	for i, v := range st {
		DrawString(screen, 3, i+2,
			fmt.Sprint("Сокровища: ", v.Gold, " Run ", v.RunNum, " Lvl ", v.Lvl, " Kill ", v.Hit, " miss ", v.Miss, " Cell ", v.Cell, " Food, Flask, Scroll ", v.Food, v.Flask, v.Scroll))
	}
}

func runIventory(screen tcell.Screen) {
	DrawString(screen, 3, 3, "IVENTORY")
}
func runGameMenu(screen tcell.Screen) {
	DrawString(screen, 3, 3, "1, m) Continue")
	DrawString(screen, 3, 4, "2) Save and Exit")
}

func runGame(screen tcell.Screen) {

	for i := range 100 {
		for j := range 40 {
			screen.SetContent(i, j, s.Mapp[i][j], nil, tcell.StyleDefault)
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
		fmt.Sprintf("Пройдено:  %d", s.Statistic.Cell),
	)
	DrawString(
		screen,
		102,
		17,
		fmt.Sprintf("Нанесено ударов:  %d", s.Statistic.Hit),
	)
	DrawString(
		screen,
		102,
		18,
		fmt.Sprintf("Получено ударов:  %d", s.Statistic.Miss),
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

func openMenu(screen tcell.Screen) {
	DrawString(screen, 10, 10, fmt.Sprintf("Inventory: %d", "!"))
}

func checkInput(ev *tcell.Event, player *Player, mapp *[100][40]rune) {
	switch event := (*ev).(type) {
	case *tcell.EventKey:
		switch event.Rune() {
		// удалить потом
		case 'q':
			isRun = false
		case 'w':
			if s.checkFightY(-1, player) {
				if CheckRoof(-1, *mapp, player) {
					player.Y -= 1
					s.Statistic.Cell += 1
				}
			}
		case 'a':
			if s.checkFightX(-1, player) {
				if CheckWall(-1, *mapp, player) {
					player.X -= 1
					s.Statistic.Cell += 1
				}
			}
		case 's':
			if s.checkFightY(1, player) {
				if CheckRoof(1, *mapp, player) {
					player.Y += 1
					s.Statistic.Cell += 1
				}
			}
		case 'd':
			if s.checkFightX(1, player) {
				if CheckWall(1, *mapp, player) {
					player.X += 1
					s.Statistic.Cell += 1
				}
			}
		case 'i':
			switch s.CurrentState {
			case StatePlay:
				s.CurrentState = StateInventory
			case StateInventory:
				s.CurrentState = StatePlay
			}
		case 'h': // weapon
			switch s.CurrentState {
			case StatePlay:
				s.PlayerStatusDoing = "Use weapon"
			}
		case 'j': // food
			switch s.CurrentState {
			case StatePlay:
				s.PlayerStatusDoing = "Use food"
			}
		case 'k': // flask
			switch s.CurrentState {
			case StatePlay:
				s.PlayerStatusDoing = "Use flask"
			}
		case 'e': // scroll
			switch s.CurrentState {
			case StatePlay:
				s.PlayerStatusDoing = "Use scroll"
			}
		case '1':
			switch s.CurrentState {
			case StateGameMenu:
				s.CurrentState = StatePlay
			case StateMainMenu:
				s.CurrentState = StateStartNewGame
			case StateGameOver:
				s.CurrentState = StateMainMenu
			case StateStatistic:
				s.CurrentState = StateMainMenu
			}
		case '2':
			switch s.CurrentState {
			case StateMainMenu:
				var err error
				s, err = ReadSave[*Session]("session.json")
				if err != nil {
					fmt.Println(err, "12123123")
					//return
				}
				s.CurrentState = StatePlay
			case StateGameMenu:
				err := WriteSave("session.json", s)
				if err != nil {
					fmt.Println(err)
					return
				}
				s.CurrentState = StateMainMenu
			}
		case '3':
			switch s.CurrentState {
			case StateMainMenu:
				s.CurrentState = StateStatistic
			case StateStatistic:
				s.CurrentState = StateMainMenu
			}
		case '4':
			switch s.CurrentState {
			case StateMainMenu:
				isRun = false
			}
		case 'm':
			switch s.CurrentState {
			case StatePlay:
				s.CurrentState = StateGameMenu
			case StateGameMenu:
				s.CurrentState = StatePlay
			}
		}
	}

}

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
	default:
		return false
	}
}
