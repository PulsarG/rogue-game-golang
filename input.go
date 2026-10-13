package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

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
