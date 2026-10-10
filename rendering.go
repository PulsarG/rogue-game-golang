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

	В игре собирается и отображается в отдельном представлении статистика всех прохождений, отсортированная по количеству набранных сокровищ: количество сокровищ, достигнутый уровень, количество побежденных противников, количество съеденной еды, количество выпитых эликсиров, количество прочитанных свитков, количество нанесенных и пропущенных ударов, количество пройденных клеток.
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
	s = initNewGame()
	isRun = true
	for isRun {
		screen.Clear()
		switch s.CurrentState {
		case StateStartNewGame:
			GameLvl = 20
			s = initNewGame()
			s.CurrentState = StatePlay
		case StateNextLvl:
			s = initNewGame()
			s.CurrentState = StatePlay

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
		}
		screen.Show()

		event := screen.PollEvent()
		checkInput(&event, s.Player, &s.Mapp)
	}
}

func initNewGame() *Session {

	s := Session{}
	s.GenMap()
	s.Player = PlayerInit(&s.Rooms[s.StartRoomIdx])
	s.CurrentLvl = GameLvl
	s.EnemysListInit()
	//s.CurrentState = StateMainMenu
	s.EndRoom = s.EndRoomInit()
	return &s
}

func runMenu(screen tcell.Screen) {
	DrawString(screen, 3, 3, "1) New Game")
	DrawString(screen, 3, 4, "2) Continue")
	DrawString(screen, 3, 5, "3) Leaderboard")
	DrawString(screen, 3, 7, "4, q) Exit")
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
		15,
		fmt.Sprintf(": /n\n %c", s.Mapp[s.Player.X][s.Player.Y+1]),
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
				}
			}
		case 'a':
			if s.checkFightX(-1, player) {
				if CheckWall(-1, *mapp, player) {
					player.X -= 1
				}
			}
		case 's':
			if s.checkFightY(1, player) {
				if CheckRoof(1, *mapp, player) {
					player.Y += 1
				}
			}
		case 'd':
			if s.checkFightX(1, player) {
				if CheckWall(1, *mapp, player) {
					player.X += 1
				}
			}
		case 'i':
			switch s.CurrentState {
			case StatePlay:
				s.CurrentState = StateInventory
			case StateInventory:
				s.CurrentState = StatePlay
			}
		case '1':
			switch s.CurrentState {
			case StateGameMenu:
				s.CurrentState = StatePlay
			case StateMainMenu:
				s.CurrentState = StateStartNewGame
			}
		case '2':
			switch s.CurrentState {
			case StateMainMenu:
				var err error
				s, err = ReadSave[*Session]("session.json")
				if err != nil {
					fmt.Println(err)
					return
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
