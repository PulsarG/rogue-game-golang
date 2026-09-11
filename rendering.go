package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

//Отображение
//
/*   Рендеринг среды — стены, пол, проем в стене, коридоры между комнатами.
    Рендеринг акторов — персонаж, противники, подбираемые предметы.
    Рендеринг интерфейса — отображение игрового интерфейса (панель статуса, инвентаря, простое меню).
    Туман войны — зависимость рендеринга сцены от состояния игры:
        Неизведанные комнаты и коридоры не отображаются.
        Просмотренные комнаты, но в которых не находится игрок, отображаются только как стены.
        В комнате, в которой находится игрок, отображаются стены, пол, акторы и предметы.
        При нахождении в непосредственной близости с комнатой со стороны коридора туман войны рассеивается только на области прямой видимости (применяется алгоритм Ray Casting и алгоритм Брезенхэма для определения видимой области).
    Пример реализации рендеринга уровней представлен в папке code-samples.

Управление

    Управление персонажем:
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

func DrawString(screen tcell.Screen, x, y int, msg string) {
	for i, char := range msg {
		screen.SetContent(x+i, y, char, nil, tcell.StyleDefault)
	}
}

func CheckRoof(vector int, mapp [100][40]rune, player *Sprite) bool {
	if mapp[player.X][player.Y+vector] != '.' {
		return false
	} else {
		return true
	}
}

func CheckWall(vector int, mapp [100][40]rune, player *Sprite) bool {
	if mapp[player.X+vector][player.Y] != '.' {
		return false
	} else {
		return true
	}
}

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

	for isRun {
		screen.Clear()
		for i := 0; i < 100; i++ {
			for j := 0; j < 40; j++ {
				screen.SetContent(i, j, mapp[i][j], nil, tcell.StyleDefault)
			}
		}
		txt := 0
		DrawString(
			screen,
			3,
			29,
			fmt.Sprint("Message: ", txt),
		)
		player.Draw(screen)
		screen.Show()

		playerMoved := false
		event := screen.PollEvent()

		switch event := event.(type) {
		case *tcell.EventKey:
			switch event.Rune() {
			case 'q':
				isRun = false
			case 'w':
				if CheckRoof(-1, mapp, player) {
					player.Y -= 1
					playerMoved = true
				} else {
					break
				}
			case 'a':
				if CheckWall(-1, mapp, player) {
					player.X -= 1
					playerMoved = true
				} else {
					break
				}
			case 's':
				if CheckRoof(1, mapp, player) {
					player.Y += 1
					playerMoved = true
				} else {
					break
				}
			case 'd':
				if CheckWall(1, mapp, player) {
					player.X += 1
					playerMoved = true
				} else {
					break
				}
			}
		}
		if playerMoved {
			txt++
		}
	}
}
