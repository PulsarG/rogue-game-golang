package main

func main() {
	s := Session{}
	s.GenMap()
	s.Player = PlayerInit(&s.Rooms[s.StartRoomIdx])
	RunView(&s)
}
