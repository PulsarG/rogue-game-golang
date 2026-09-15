package main

func main() {
	s := Session{}
	s.Mapp, s.Rooms = GenMap()
	s.Player = PlayerInit(&s.Rooms[0])
	RunView(&s)
}
