package main

import (
	"math/rand"
)

/*Реализуй модуль генерации уровней в уровне domain.

  Каждый уровень должен быть логически разделен на 9 секций, в каждой из которых случайным образом генерируется комната с произвольным размером и положением.
  Комнаты произвольным образом соединены коридорами. Коридоры имеют свою геометрию, по ним тоже можно ходить, а значит, их координаты необходимо генерировать и хранить. При генерации необходимо проверять, что сгенерированный граф комнат — связный и не имеет ошибок.
  На каждом уровне одна комната помечена как стартовая, и еще одна — как конечная. В стартовой комнате начинается игровая сессия, а в конечной располагается блок, при прикосновении к которому игрок перемещается на следующий уровень.
  Пример реализации генерации уровней представлен в папке code-samples.
*/

var minWidth, maxWidth = 10, 20
var minHeigth, maxHeigth = 4, 7
var mapp = [100][40]rune{}

func (s *Session) GenMap() {

	setStartEndRoom(s)

	s.Rooms[0].StartPointFromX = 2 + rand.Intn(12) // для каждой комнаты, рандомить
	s.Rooms[0].StartPointFromY = 1 + rand.Intn(5)  // для каждой комнаты, рандомить
	s.Rooms[0].CreateRoom()

	s.Rooms[1].StartPointFromX = 35 + rand.Intn(12) // для каждой комнаты, рандомить
	s.Rooms[1].StartPointFromY = 1 + rand.Intn(5)   // для каждой комнаты, рандомить
	s.Rooms[1].CreateRoom()

	s.Rooms[2].StartPointFromX = 68 + rand.Intn(12) // для каждой комнаты, рандомить
	s.Rooms[2].StartPointFromY = 1 + rand.Intn(5)   // для каждой комнаты, рандомить
	s.Rooms[2].CreateRoom()

	s.Rooms[3].StartPointFromX = 2 + rand.Intn(12) // для каждой комнаты, рандомить
	s.Rooms[3].StartPointFromY = 13 + rand.Intn(5) // для каждой комнаты, рандомить
	s.Rooms[3].CreateRoom()

	s.Rooms[4].StartPointFromX = 35 + rand.Intn(12) // для каждой комнаты, рандомить
	s.Rooms[4].StartPointFromY = 13 + rand.Intn(5)  // для каждой комнаты, рандомить
	s.Rooms[4].CreateRoom()

	s.Rooms[5].StartPointFromX = 68 + rand.Intn(12) // для каждой комнаты, рандомить
	s.Rooms[5].StartPointFromY = 13 + rand.Intn(5)  // для каждой комнаты, рандомить
	s.Rooms[5].CreateRoom()

	s.Rooms[6].StartPointFromX = 2 + rand.Intn(12) // для каждой комнаты, рандомить
	s.Rooms[6].StartPointFromY = 26 + rand.Intn(5) // для каждой комнаты, рандомить
	s.Rooms[6].CreateRoom()

	s.Rooms[7].StartPointFromX = 35 + rand.Intn(12) // для каждой комнаты, рандомить
	s.Rooms[7].StartPointFromY = 26 + rand.Intn(5)  // для каждой комнаты, рандомить
	s.Rooms[7].CreateRoom()

	s.Rooms[8].StartPointFromX = 68 + rand.Intn(12) // для каждой комнаты, рандомить
	s.Rooms[8].StartPointFromY = 26 + rand.Intn(5)  // для каждой комнаты, рандомить
	s.Rooms[8].CreateRoom()

	s.Mapp = mapp
	s.connectRoom()
}

func setStartEndRoom(s *Session) {
	start := rand.Intn(9)
	s.StartRoomIdx = start
	s.Rooms[start].IsStart = true
	var end int
	for {
		end = rand.Intn(9)
		if start != end {
			s.EndRoomIdx = end
			s.Rooms[end].IsEnd = true
			break
		}
	}
}

func (r *Room) CreateRoom() {
	r.SetSize()
	r.GenRoom(&mapp)

}
func (r *Room) SetSize() {
	r.Width = minWidth + rand.Intn(maxWidth-minWidth+1)
	r.Heigth = minHeigth + rand.Intn(maxHeigth-minHeigth+1)
	r.EndPointFromY = r.StartPointFromY + r.Heigth
	r.CenterX = r.StartPointFromX + r.Width/2
	r.CenterY = r.StartPointFromY + r.Heigth/2
}

func (r *Room) GenRoom(
	mapp *[100][40]rune,
) {
	for i := r.StartPointFromX; i < r.StartPointFromX+r.Width; i++ {
		mapp[i][r.StartPointFromY] = '-'
		mapp[i][r.EndPointFromY] = '-'
	}
	for i := r.StartPointFromY; i < r.EndPointFromY+1; i++ {
		mapp[r.StartPointFromX][i] = '|'
		mapp[r.Width+r.StartPointFromX][i] = '|'
	}
	for i := r.StartPointFromX + 1; i < r.StartPointFromX+r.Width; i++ {
		for j := r.StartPointFromY + 1; j < r.EndPointFromY; j++ {
			mapp[i][j] = '.'
		}
	}

}

func (s *Session) connectRoom() {
	tun := GenTunnels()
	for _, vv := range *tun {
		v := vv[0]
		k := vv[1]
		if rand.Intn(2) == 0 {
			s.applyHorTunel(v, k)
			s.applyVerTunel(v, k)
		} else {
			s.applyVerTunel(v, k)
			s.applyHorTunel(v, k)
		}
	}
}

func (s *Session) applyHorTunel(a, b int) {
	minX, maxX := s.Rooms[a].CenterX, s.Rooms[b].CenterX
	y := s.Rooms[a].CenterY
	if s.Rooms[a].CenterX > s.Rooms[b].CenterX {
		maxX, minX = s.Rooms[a].CenterX, s.Rooms[b].CenterX
	}
	for x := minX; x <= maxX; x++ {
		if s.Mapp[x][y] == 0 {
			s.Mapp[x][y] = '#'
		} else if s.Mapp[x][y] == '|' {
			s.Mapp[x][y] = '+'
		}
	}
}

func (s *Session) applyVerTunel(a, b int) {
	minY, maxY := s.Rooms[a].CenterY, s.Rooms[b].CenterY
	x := s.Rooms[b].CenterX
	if s.Rooms[a].CenterY > s.Rooms[b].CenterY {
		maxY, minY = s.Rooms[a].CenterY, s.Rooms[b].CenterY
	}
	for y := minY; y <= maxY; y++ {
		if s.Mapp[x][y] == 0 {
			s.Mapp[x][y] = '#'
		} else if s.Mapp[x][y] == '-' {
			s.Mapp[x][y] = '+'
		}
	}
}
