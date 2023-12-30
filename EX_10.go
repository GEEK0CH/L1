package main

import (
	"fmt"
)

func combGroups1(mass []float64) map[int][]float64 {
	m := make(map[int][]float64) //создание мапы с ключами
	m[-20] = []float64{}
	m[10] = []float64{}
	m[20] = []float64{}
	m[30] = []float64{}

	for i := 0; i <= len(mass)-1; i++ { //цикл для перебора массива и записи
		switch {
		case -30 < mass[i] && -20 > mass[i]:
			m[-20] = append(m[-20], mass[i])
		case 10 < mass[i] && 20 > mass[i]:
			m[10] = append(m[10], mass[i])
		case 20 < mass[i] && 30 > mass[i]:
			m[20] = append(m[20], mass[i])
		case 30 < mass[i] && 40 > mass[i]:
			m[30] = append(m[30], mass[i])
		}
	}
	return m
}

func combGroups2(mass []float64) map[int][]float64 {
	m := make(map[int][]float64)        //создание мапы
	for i := 0; i <= len(mass)-1; i++ { //цикл для перебора mass и создания соответствующих ключей
		m[int(mass[i]/10)*10] = append(m[int(mass[i]/10)*10], mass[i])
	}
	return m
}

func main() {
	var mass = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(combGroups1(mass))
	fmt.Println(combGroups2(mass))
}
