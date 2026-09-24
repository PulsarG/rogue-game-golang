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
var txt int

type Stat struct {
	IsRun, IsMainGameRun, IsMenuOpen, IsGameMenuOpen, IsInventoryOpen, IsStartNewGame bool
}

var s *Session

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
	//player := s.Player
	s = initNewGame()
	state := Stat{true, false, true, false, false, false}
	for state.IsRun {
		if state.IsStartNewGame {
			s = initNewGame()
			state.IsStartNewGame = false
		}
		screen.Clear()
		if state.IsMenuOpen {
			runMenu(screen)
		} else if state.IsInventoryOpen {
			//	mapp := s.Mapp
			//runGame(screen, s)
			runIventory(screen)
		} else if state.IsGameMenuOpen {
			runGameMenu(screen)
		} else if state.IsMainGameRun {
			runFight()
			runGame(screen, s)
		} else {
			//openMenu(screen)
		}
		screen.Show()

		event := screen.PollEvent()
		checkInput(screen, &event, &state, s.Player, &s.Mapp)
	}
}

func initNewGame() *Session {

	s := Session{}
	s.GenMap()
	s.Player = PlayerInit(&s.Rooms[s.StartRoomIdx])
	s.CurrentLvl = 5 + rand.Intn(5)
	s.EnemysListInit()
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
	DrawString(screen, 3, 3, "m) Continue")
	DrawString(screen, 3, 4, "e) Save and Exit")
}

func runGame(screen tcell.Screen, s *Session) {

	for i := range 100 {
		for j := range 40 {
			screen.SetContent(i, j, s.Mapp[i][j], nil, tcell.StyleDefault)
		}
	}
	DrawString(
		screen,
		102,
		1,
		fmt.Sprint("Message: ", s.StatusDoing),
	)
	txt = 0
	DrawString(
		screen,
		3,
		42,
		fmt.Sprint("Message: ", txt),
	)
	Draw(s.Player.X, s.Player.Y, s.Player.Sprite, screen, &s.Player.ColorSprite)
	for _, e := range s.Enemys {
		Draw(e.X, e.Y, e.Sprite, screen, &e.ColorSprite)
	}
}

func runFight() {
	for i, e := range s.Enemys {
		if e.CheckFight(s.Player) {
			s.StatusDoing = fmt.Sprintf("Killing enemys: %c", e.Sprite)
			s.Enemys[i] = s.Enemys[len(s.Enemys)-1]
			s.Enemys = s.Enemys[:len(s.Enemys)-1]
			//s.Enemys = append(s.Enemys[:i], s.Enemys[i+1:]...)
		}
	}
}

func openMenu(screen tcell.Screen) {
	DrawString(screen, 10, 10, fmt.Sprintf("Inventory: %d", txt))
}

func checkInput(screen tcell.Screen, ev *tcell.Event, state *Stat, player *Player, mapp *[100][40]rune) {
	playerMoved := false
	switch event := (*ev).(type) {
	case *tcell.EventKey:
		switch event.Rune() {
		case 'q':
			state.IsRun = false
		case 'w':
			if CheckRoof(-1, *mapp, player) {
				player.Y -= 1
				playerMoved = true
			}
		case 'a':
			if CheckWall(-1, *mapp, player) {
				player.X -= 1
				playerMoved = true
			}
		case 's':
			if CheckRoof(1, *mapp, player) {
				player.Y += 1
				playerMoved = true
			}
		case 'd':
			if CheckWall(1, *mapp, player) {
				player.X += 1
				playerMoved = true
			}
		case 'i':
			if !state.IsInventoryOpen {
				state.IsInventoryOpen = true
			} else {
				state.IsInventoryOpen = false
			}
		case 'e':
			if state.IsGameMenuOpen {
				err := WriteSave("session.json", s)
				if err != nil {
					fmt.Println(err)
					return
				}
				state.IsGameMenuOpen = false
				state.IsMenuOpen = true
				state.IsMainGameRun = false
			}
		case '1':
			state.IsMenuOpen = false
			state.IsStartNewGame = true
			state.IsMainGameRun = true
		case '2':
			if state.IsMenuOpen {
				var err error
				s, err = ReadSave[*Session]("session.json")
				if err != nil {
					fmt.Println(err)
					return
				}
				state.IsMenuOpen = false
				state.IsMainGameRun = true
			}
		case '4':
			state.IsRun = false
		case 'm':
			// новая карта
			//			screen.Clear()
			//			s.GenMap()
			if !state.IsGameMenuOpen {
				state.IsGameMenuOpen = true
			} else {
				state.IsGameMenuOpen = false
			}
		}
	}

	if playerMoved {
		txt++
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
	default:
		return false
	}
}
