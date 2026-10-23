package main

import (
	"fmt"
	"math/rand"
)

func runFight() {
	alive := s.Enemys[:0]
	for _, e := range s.Enemys {
		if e.CheckFight(s.Player) {
			s.PlayerStatusDoing = fmt.Sprintf("Игрок убил %s", e.Name)
			s.Statistic.Enemy++
			continue
		}
		alive = append(alive, e)
	}
	s.Enemys = alive
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
	if e.Sprite == 'v' && !e.IsFirstHit {
		s.PlayerStatusDoing = "Игрок промахнулся"
		e.IsFirstHit = true
		return
	}
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
	if e.Sprite == 'O' && e.IsSecondHitOgr {
		e.IsSecondHitOgr = false
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
		if e.Sprite == 'v' {
			s.Player.MaxHp -= 1
		}
		e.IsSecondHitOgr = true
		s.Statistic.Miss += 1
		s.EnemyStatusDoing = fmt.Sprintf("Враг нанес %d урона", damage)
	}
}
