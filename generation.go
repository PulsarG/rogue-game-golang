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

func GenMap() [100][40]rune {
	mapp := [100][40]rune{}

	startPointFromX := 3 // для каждой комнаты, рандомить
	startPointFromY := 3 // для каждой комнаты, рандомить
	widthRoom, heigthRoom, endPointFromY := setSize(&startPointFromY)
	genRoom(&startPointFromX, &startPointFromY, &endPointFromY, &widthRoom, &heigthRoom, &mapp)

	startPointFromX = 33 // для каждой комнаты, рандомить
	startPointFromY = 3  // для каждой комнаты, рандомить
	widthRoom, heigthRoom, endPointFromY = setSize(&startPointFromY)
	genRoom(&startPointFromX, &startPointFromY, &endPointFromY, &widthRoom, &heigthRoom, &mapp)

	startPointFromX = 66 // для каждой комнаты, рандомить
	startPointFromY = 3  // для каждой комнаты, рандомить
	widthRoom, heigthRoom, endPointFromY = setSize(&startPointFromY)
	genRoom(&startPointFromX, &startPointFromY, &endPointFromY, &widthRoom, &heigthRoom, &mapp)
	startPointFromX = 3  // для каждой комнаты, рандомить
	startPointFromY = 13 // для каждой комнаты, рандомить
	widthRoom, heigthRoom, endPointFromY = setSize(&startPointFromY)
	genRoom(&startPointFromX, &startPointFromY, &endPointFromY, &widthRoom, &heigthRoom, &mapp)
	startPointFromX = 33 // для каждой комнаты, рандомить
	startPointFromY = 13 // для каждой комнаты, рандомить
	widthRoom, heigthRoom, endPointFromY = setSize(&startPointFromY)
	genRoom(&startPointFromX, &startPointFromY, &endPointFromY, &widthRoom, &heigthRoom, &mapp)
	startPointFromX = 66 // для каждой комнаты, рандомить
	startPointFromY = 13 // для каждой комнаты, рандомить
	widthRoom, heigthRoom, endPointFromY = setSize(&startPointFromY)
	genRoom(&startPointFromX, &startPointFromY, &endPointFromY, &widthRoom, &heigthRoom, &mapp)
	startPointFromX = 3  // для каждой комнаты, рандомить
	startPointFromY = 23 // для каждой комнаты, рандомить
	widthRoom, heigthRoom, endPointFromY = setSize(&startPointFromY)
	genRoom(&startPointFromX, &startPointFromY, &endPointFromY, &widthRoom, &heigthRoom, &mapp)
	startPointFromX = 33 // для каждой комнаты, рандомить
	startPointFromY = 23 // для каждой комнаты, рандомить
	widthRoom, heigthRoom, endPointFromY = setSize(&startPointFromY)
	genRoom(&startPointFromX, &startPointFromY, &endPointFromY, &widthRoom, &heigthRoom, &mapp)
	startPointFromX = 66 // для каждой комнаты, рандомить
	startPointFromY = 23 // для каждой комнаты, рандомить
	widthRoom, heigthRoom, endPointFromY = setSize(&startPointFromY)
	genRoom(&startPointFromX, &startPointFromY, &endPointFromY, &widthRoom, &heigthRoom, &mapp)
	return mapp
}

func setSize(startPointFromY *int) (int, int, int) {
	widthRoom := minWidth + rand.Intn(maxWidth-minWidth+1)
	heigthRoom := minHeigth + rand.Intn(maxHeigth-minHeigth+1)
	endPointFromY := *startPointFromY + heigthRoom
	return widthRoom, heigthRoom, endPointFromY
}

func genRoom(
	startPointFromX *int,
	startPointFromY *int,
	endPointFromY *int,
	widthRoom *int,
	heigthRoom *int,
	mapp *[100][40]rune,
) {
	for i := *startPointFromX; i < *startPointFromX+*widthRoom; i++ {
		mapp[i][*startPointFromY] = '-'
		mapp[i][*endPointFromY] = '-'
	}
	for i := *startPointFromY; i < *endPointFromY+1; i++ {
		mapp[*startPointFromX][i] = '|'
		mapp[*widthRoom+*startPointFromX][i] = '|'
	}
	for i := *startPointFromX + 1; i < *startPointFromX+*widthRoom; i++ {
		for j := *startPointFromY + 1; j < *endPointFromY; j++ {
			mapp[i][j] = '.'
		}
	}

}
