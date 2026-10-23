package main

func CheckRoof(vector int, mapp *[100][40]rune, player *Player) bool {
	switch mapp[player.X][player.Y+vector] {
	case '.':
		return true
	case '#':
		return true
	case '+':
		return true
	case '%':
		GameLvl = s.CurrentLvl + 1
		if GameLvl == 22 {
			s.writeStatistic()
			s.writeNumRun()
			s.CurrentState = StateGameOver
			return true
		}
		s.CurrentState = StateNextLvl
		return true
	case 'G':
		s.Statistic.Gold += 1
		s.Mapp[player.X][player.Y+vector] = '.'
		return true
	case 'F':
		if len(s.Backpack.Flasks) >= 9 {
			s.PlayerStatusDoing = TXT_NO_MORE_PICKUP
			return false
		}
		s.Backpack.Flasks = append(s.Backpack.Flasks, *s.InitFlask())
		s.Mapp[player.X][player.Y+vector] = '.'
		return true
	case 'E':
		if len(s.Backpack.Scrolls) >= 9 {
			s.PlayerStatusDoing = TXT_NO_MORE_PICKUP
			return false
		}
		s.Backpack.Scrolls = append(s.Backpack.Scrolls, *s.InitScroll())
		s.Mapp[player.X][player.Y+vector] = '.'
		return true
	case 'W':
		if len(s.Backpack.Weapons) >= 9 {
			s.PlayerStatusDoing = TXT_NO_MORE_PICKUP
			return false
		}
		s.Mapp[player.X][player.Y+vector] = '.'
		s.Backpack.Weapons = append(s.Backpack.Weapons, *s.InitOneWeapon())
		return true
	case 'J':
		if s.Backpack.Food > 9 {
			s.PlayerStatusDoing = TXT_NO_MORE_PICKUP
			return false
		}
		s.Mapp[player.X][player.Y+vector] = '.'
		s.Backpack.Food++
		return true
	default:
		return false
	}
}

func CheckWall(vector int, mapp *[100][40]rune, player *Player) bool {
	switch mapp[player.X+vector][player.Y] {
	case '.':
		return true
	case '#':
		return true
	case '+':
		return true
	case '%':
		GameLvl = s.CurrentLvl + 1
		if GameLvl == 22 {
			s.writeStatistic()
			s.writeNumRun()
			s.CurrentState = StateGameOver
			return true
		}
		s.CurrentState = StateNextLvl
		return true
	case 'G':
		s.Statistic.Gold += 1
		s.Mapp[player.X+vector][player.Y] = '.'
		return true
	case 'F':
		if len(s.Backpack.Flasks) >= 9 {
			s.PlayerStatusDoing = TXT_NO_MORE_PICKUP
			return false
		}
		s.Backpack.Flasks = append(s.Backpack.Flasks, *s.InitFlask())
		s.Mapp[player.X+vector][player.Y] = '.'
		return true
	case 'E':
		if len(s.Backpack.Scrolls) >= 9 {
			s.PlayerStatusDoing = TXT_NO_MORE_PICKUP
			return false
		}
		s.Backpack.Scrolls = append(s.Backpack.Scrolls, *s.InitScroll())
		s.Mapp[player.X+vector][player.Y] = '.'
		return true
	case 'W':
		if len(s.Backpack.Weapons) >= 9 {
			s.PlayerStatusDoing = TXT_NO_MORE_PICKUP
			return false
		}
		s.Backpack.Weapons = append(s.Backpack.Weapons, *s.InitOneWeapon())
		s.Mapp[player.X+vector][player.Y] = '.'
		return true
	case 'J':
		if s.Backpack.Food > 9 {
			s.PlayerStatusDoing = TXT_NO_MORE_PICKUP
			return false
		} else {
			s.Backpack.Food++
			s.Mapp[player.X+vector][player.Y] = '.'
			return true
		}
	default:
		return false
	}
}
