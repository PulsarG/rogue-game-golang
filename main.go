package main

import "math/rand"

func main() {
	s := Session{}
	s.GenMap()
	s.Player = PlayerInit(&s.Rooms[s.StartRoomIdx])
	s.CurrentLvl = 5 + rand.Intn(5)
	s.EnemysListInit()
	RunView(&s)
}
