package main

import (
	"math/rand"
)

func initTunnels() *[9][]int {
	var TunnelAll [9][]int

	TunnelAll[0] = []int{1, 3}
	TunnelAll[1] = []int{0, 2, 4}
	TunnelAll[2] = []int{1, 5}
	TunnelAll[3] = []int{0, 4, 6}
	TunnelAll[4] = []int{1, 3, 5, 7}
	TunnelAll[5] = []int{2, 4, 8}
	TunnelAll[6] = []int{3, 7}
	TunnelAll[7] = []int{4, 6, 8}
	TunnelAll[8] = []int{5, 7}

	return &TunnelAll
}

func GenTunnels() *[9][2]int {
	//tun := make(map[int]int, 9)
	var res [9][2]int
	tAll := initTunnels()
	for i, v := range tAll {
		l := len(v)
		var x int
		if l != 0 {
			x = v[rand.Intn(l)]
			deleteDouplicate(&x, &i, tAll)
			res[i][0] = x
			res[i][1] = i
		}
	}
	return &res
}

func deleteDouplicate(x *int, i *int, tAll *[9][]int) {

	iRemove := -1
	for j, w := range tAll[*x] {
		if w == *i {
			iRemove = j
			break
		}
	}
	tAll[*x] = append(tAll[*x][:iRemove], tAll[*x][iRemove+1:]...)
}
