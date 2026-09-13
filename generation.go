package main

import "math/rand"

/*Реализуй модуль генерации уровней в уровне domain.

  Каждый уровень должен быть логически разделен на 9 секций, в каждой из которых случайным образом генерируется комната с произвольным размером и положением.
  Комнаты произвольным образом соединены коридорами. Коридоры имеют свою геометрию, по ним тоже можно ходить, а значит, их координаты необходимо генерировать и хранить. При генерации необходимо проверять, что сгенерированный граф комнат — связный и не имеет ошибок.
  На каждом уровне одна комната помечена как стартовая, и еще одна — как конечная. В стартовой комнате начинается игровая сессия, а в конечной располагается блок, при прикосновении к которому игрок перемещается на следующий уровень.
  Пример реализации генерации уровней представлен в папке code-samples.
*/

var minWidth, maxWidth = 10, 20
var minHeigth, maxHeigth = 4, 7
var mapp = [100][40]rune{}

type Room struct {
	StartPointFromX, StartPointFromY, EndPointFromY, Width, Heigth, CenterX, CenterY int
}

func GenMap() [100][40]rune {

	rooms := [9]Room{}
	//r1 := Room{}
	rooms[0].StartPointFromX = 2 + rand.Intn(12) // для каждой комнаты, рандомить
	rooms[0].StartPointFromY = 1 + rand.Intn(5)  // для каждой комнаты, рандомить
	rooms[0].CreateRoom()

	//r2 := Room{}
	rooms[1].StartPointFromX = 35 + rand.Intn(12) // для каждой комнаты, рандомить
	rooms[1].StartPointFromY = 1 + rand.Intn(5)   // для каждой комнаты, рандомить
	rooms[1].CreateRoom()

	r3 := Room{}
	r3.StartPointFromX = 68 + rand.Intn(12) // для каждой комнаты, рандомить
	r3.StartPointFromY = 1 + rand.Intn(5)   // для каждой комнаты, рандомить
	r3.CreateRoom()

	r4 := Room{}
	r4.StartPointFromX = 2 + rand.Intn(12) // для каждой комнаты, рандомить
	r4.StartPointFromY = 13 + rand.Intn(5) // для каждой комнаты, рандомить
	r4.CreateRoom()

	r5 := Room{}
	r5.StartPointFromX = 35 + rand.Intn(12) // для каждой комнаты, рандомить
	r5.StartPointFromY = 13 + rand.Intn(5)  // для каждой комнаты, рандомить
	r5.CreateRoom()

	r6 := Room{}
	r6.StartPointFromX = 68 + rand.Intn(12) // для каждой комнаты, рандомить
	r6.StartPointFromY = 13 + rand.Intn(5)  // для каждой комнаты, рандомить
	r6.CreateRoom()

	r7 := Room{}
	r7.StartPointFromX = 2 + rand.Intn(12) // для каждой комнаты, рандомить
	r7.StartPointFromY = 26 + rand.Intn(5) // для каждой комнаты, рандомить
	r7.CreateRoom()

	r8 := Room{}
	r8.StartPointFromX = 35 + rand.Intn(12) // для каждой комнаты, рандомить
	r8.StartPointFromY = 26 + rand.Intn(5)  // для каждой комнаты, рандомить
	r8.CreateRoom()

	r9 := Room{}
	r9.StartPointFromX = 68 + rand.Intn(12) // для каждой комнаты, рандомить
	r9.StartPointFromY = 26 + rand.Intn(5)  // для каждой комнаты, рандомить
	r9.CreateRoom()

	connectRoom(&mapp, &rooms)

	return mapp
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

func connectRoom(mapp *[100][40]rune, r *[9]Room) {
	if rand.Intn(2) == 0 {
		applyHorTunel(mapp, r[0].CenterX, r[1].CenterX, r[0].CenterY)
		applyVerTunel(mapp, r[0].CenterY, r[1].CenterY, r[1].CenterX)
	} else {
		applyVerTunel(mapp, r[0].CenterY, r[1].CenterY, r[0].CenterX)
		applyHorTunel(mapp, r[0].CenterX, r[1].CenterX, r[1].CenterY)
	}
}

func applyHorTunel(mapp *[100][40]rune, x1, x2, y int) {
	minX, maxX := x1, x2
	if x1 > x2 {
		minX, maxX = x2, x1
	}
	for x := minX; x <= maxX; x++ {
		if mapp[x][y] == 0 || mapp[x][y] == '|' {
			mapp[x][y] = '#'
		}
	}
}

func applyVerTunel(mapp *[100][40]rune, y1, y2, x int) {
	minY, maxY := y1, y2
	if y1 > y2 {
		minY, maxY = y2, y1
	}
	for y := minY; y <= maxY; y++ {
		mapp[x][y] = '.'
	}
}
