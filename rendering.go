package main

import (
	"fmt"
	"log"

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
	IsRun, IsMenuOpen, IsInventoryOpen bool
}

func RunView(s *Session) {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()
	err = screen.Init()
	if err != nil {
		log.Fatal(err)
	}

	mapp := s.Mapp
	player := s.Player
	state := Stat{true, true, false}
	for state.IsRun {
		screen.Clear()
		if state.IsMenuOpen {
			runMenu(screen)
		} else if !state.IsInventoryOpen {
			runGame(screen, player, &mapp)
		} else {
			openMenu(screen)
		}
		screen.Show()

		event := screen.PollEvent()
		checkInput(&event, &state, player, &mapp)
	}
}

func runMenu(screen tcell.Screen) {
	DrawString(screen, 3, 3, "1) New Game")
	DrawString(screen, 3, 4, "2) Continue")
	DrawString(screen, 3, 5, "3) Leaderboard")
	DrawString(screen, 3, 7, "4, q) Exit")
}

func runGame(screen tcell.Screen, player *Player, mapp *[100][40]rune) {

	for i := range 100 {
		for j := range 40 {
			screen.SetContent(i, j, mapp[i][j], nil, tcell.StyleDefault)
		}
	}
	txt = 0
	DrawString(
		screen,
		3,
		42,
		fmt.Sprint("Message: ", txt),
	)
	Draw(player, screen)
}

func openMenu(screen tcell.Screen) {
	DrawString(screen, 10, 10, fmt.Sprintf("Inventory: %d", txt))
}

func checkInput(ev *tcell.Event, state *Stat, player *Player, mapp *[100][40]rune) {
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
		case '1':
			state.IsMenuOpen = false
		case '2':
		case '3':
		case '4':
			state.IsRun = false
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
