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
	player := NewSprite('@', 5, 5)
	mapp := GenMap()
	isRun := true
	isMenuOpen := false
	for isRun {
		screen.Clear()
		if !isMenuOpen {
			runGame(screen, player, &mapp)
		} else {
			openMenu(screen)
		}
		screen.Show()

		event := screen.PollEvent()
		checkInput(screen, &event, &isRun, &isMenuOpen, player, &mapp)
	}
}

func runGame(screen tcell.Screen, player *Sprite, mapp *[100][40]rune) {

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
	player.Draw(screen)
}

func openMenu(screen tcell.Screen) {
	DrawString(screen, 10, 10, fmt.Sprintf("Inventory: %d", txt))
}

func checkInput(screen tcell.Screen, ev *tcell.Event, isRun *bool, isMenuOpen *bool, player *Sprite, mapp *[100][40]rune) {
	playerMoved := false
	switch event := (*ev).(type) {
	case *tcell.EventKey:
		switch event.Rune() {
		case 'q':
			*isRun = false
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
			if !*isMenuOpen {
				*isMenuOpen = true
			} else {
				*isMenuOpen = false
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

func CheckRoof(vector int, mapp [100][40]rune, player *Sprite) bool {
	if mapp[player.X][player.Y+vector] != '.' && mapp[player.X][player.Y+vector] != '#' {
		return false
	} else {
		return true
	}
}

func CheckWall(vector int, mapp [100][40]rune, player *Sprite) bool {
	if mapp[player.X+vector][player.Y] != '.' && mapp[player.X+vector][player.Y] != '#' {
		return false
	} else {
		return true
	}
}
