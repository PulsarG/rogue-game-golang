package main

import "math/rand"

/*Реализуй модуль генерации уровней в уровне domain.

  Каждый уровень должен быть логически разделен на 9 секций, в каждой из которых случайным образом генерируется комната с произвольным размером и положением.
  Комнаты произвольным образом соединены коридорами. Коридоры имеют свою геометрию, по ним тоже можно ходить, а значит, их координаты необходимо генерировать и хранить. При генерации необходимо проверять, что сгенерированный граф комнат — связный и не имеет ошибок.
  На каждом уровне одна комната помечена как стартовая, и еще одна — как конечная. В стартовой комнате начинается игровая сессия, а в конечной располагается блок, при прикосновении к которому игрок перемещается на следующий уровень.
  Пример реализации генерации уровней представлен в папке code-samples.
*/

func GenMap() [100][40]rune {
	mapp := [100][40]rune{}
	minDlina, maxDlina := 10, 20
	minVysota, maxVysota := 4, 7
	otstupDlinySleva := 3
	dlinaRoom := minDlina + rand.Intn(maxDlina-minDlina+1)
	startVerha := 3
	vysotaRoom := minVysota + rand.Intn(maxVysota-minVysota+1)
	startNiza := startVerha + vysotaRoom

	for i := otstupDlinySleva; i < otstupDlinySleva+dlinaRoom; i++ {
		mapp[i][startVerha] = '-'
		mapp[i][startNiza] = '-'
	}
	for i := startVerha; i < startNiza+1; i++ {
		mapp[otstupDlinySleva][i] = '|'
		mapp[dlinaRoom+otstupDlinySleva][i] = '|'
	}
	for i := otstupDlinySleva + 1; i < otstupDlinySleva+dlinaRoom; i++ {
		for j := startVerha + 1; j < startNiza; j++ {
			mapp[i][j] = '.'
		}
	}
	return mapp
}
