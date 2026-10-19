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
					s.PlayerScanFov()
				}
			}
		case 'a':
			if s.checkFightX(-1, player) {
				if CheckWall(-1, *mapp, player) {
					player.X -= 1
					s.Statistic.Cell += 1
					s.PlayerScanFov()
				}
			}
		case 's':
			if s.checkFightY(1, player) {
				if CheckRoof(1, *mapp, player) {
					player.Y += 1
					s.Statistic.Cell += 1
					s.PlayerScanFov()
				}
			}
		case 'd':
			if s.checkFightX(1, player) {
				if CheckWall(1, *mapp, player) {
					player.X += 1
					s.Statistic.Cell += 1
					s.PlayerScanFov()
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
				s.CurrentState = StateWeaponList
			case StateWeaponList:
				s.CurrentState = StatePlay
			}
		case 'j': // food
			switch s.CurrentState {
			case StatePlay:
				s.PlayerStatusDoing = "Use food"
			}
		case 'k': // flask
			switch s.CurrentState {
			case StatePlay:
				s.CurrentState = StateFlaskList
				s.PlayerStatusDoing = "Какой использовать?"
			case StateFlaskList:
				s.CurrentState = StatePlay
			}
		case 'e': // scroll
			switch s.CurrentState {
			case StatePlay:
				s.CurrentState = StateScrollList
				s.PlayerStatusDoing = "Какой использовать?"
			case StateScrollList:
				s.CurrentState = StatePlay
			}
		case '0':
			switch s.CurrentState {
			case StateWeaponList:
				s.Player.Weapon.Name = ""
				s.Player.Weapon.Damage = 0
				s.DropWeapon()
				s.PlayerStatusDoing = "Оружие сброшено"
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
			case StateFlaskList:
				s.UseFlask(0)
				s.PlayerStatusDoing = "Использован эликсир"
			case StateScrollList:
				s.UseScroll(0)
				s.PlayerStatusDoing = "Использован свиток"
			case StateWeaponList:
				s.Player.Weapon = &s.Backpack.Weapons[0]
				s.UseWeapon(0)
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
			case StateStatistic:
				s.ClearStatistic()
			case StateFlaskList:
				s.UseFlask(1)
				s.PlayerStatusDoing = "Использован эликсир"
			case StateScrollList:
				s.UseScroll(1)
			case StateWeaponList:
				s.Player.Weapon = &s.Backpack.Weapons[1]
				s.UseWeapon(1)
			}
		case '3':
			switch s.CurrentState {
			case StateMainMenu:
				s.CurrentState = StateStatistic
			case StateStatistic:
				s.CurrentState = StateMainMenu
			case StateFlaskList:
				s.UseFlask(2)
				s.PlayerStatusDoing = "Использован эликсир"
			case StateScrollList:
				s.UseScroll(2)
			case StateWeaponList:
				s.Player.Weapon = &s.Backpack.Weapons[2]
				s.UseWeapon(2)
			}
		case '4':
			switch s.CurrentState {
			case StateMainMenu:
				isRun = false
			case StateFlaskList:
				s.UseFlask(3)
				s.PlayerStatusDoing = "Использован эликсир"
			case StateScrollList:
				s.UseScroll(3)
			case StateWeaponList:
				s.Player.Weapon = &s.Backpack.Weapons[3]
				s.UseWeapon(3)
			}
		case '5':
			switch s.CurrentState {
			case StateFlaskList:
				s.UseFlask(4)
				s.PlayerStatusDoing = "Использован эликсир"
			case StateScrollList:
				s.UseScroll(4)
			case StateWeaponList:
				s.Player.Weapon = &s.Backpack.Weapons[4]
				s.UseWeapon(4)
			}
		case '6':
			switch s.CurrentState {
			case StateFlaskList:
				s.UseFlask(5)
				s.PlayerStatusDoing = "Использован эликсир"
			case StateScrollList:
				s.UseScroll(5)
			case StateWeaponList:
				s.Player.Weapon = &s.Backpack.Weapons[5]
				s.UseWeapon(5)
			}
		case '7':
			switch s.CurrentState {
			case StateFlaskList:
				s.UseFlask(6)
				s.PlayerStatusDoing = "Использован эликсир"
			case StateScrollList:
				s.UseScroll(6)
			case StateWeaponList:
				s.Player.Weapon = &s.Backpack.Weapons[6]
				s.UseWeapon(6)
			}
		case '8':
			switch s.CurrentState {
			case StateFlaskList:
				s.UseFlask(7)
				s.PlayerStatusDoing = "Использован эликсир"
			case StateScrollList:
				s.UseScroll(7)
			case StateWeaponList:
				s.Player.Weapon = &s.Backpack.Weapons[7]
				s.UseWeapon(7)
			}
		case '9':
			switch s.CurrentState {
			case StateFlaskList:
				s.UseFlask(8)
				s.PlayerStatusDoing = "Использован эликсир"
			case StateScrollList:
				s.UseScroll(8)
			case StateWeaponList:
				s.Player.Weapon = &s.Backpack.Weapons[8]
				s.UseWeapon(8)
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
