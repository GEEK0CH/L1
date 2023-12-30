package main

import "fmt"

func changePlaces(fir, sec int) (int, int) { //функция со стандартвой возможностью golang
	fir, sec = sec, fir
	return fir, sec
}

func changePlacesWithCalc(fir, sec int) (int, int) { //функция с использованием вычислений
	sec = fir + sec
	fir = sec - fir
	sec = sec - fir
	return fir, sec
}

func main() {
	a := 23
	b := 11
	fmt.Println(changePlaces(a, b))
	fmt.Println(changePlacesWithCalc(a, b))
}
