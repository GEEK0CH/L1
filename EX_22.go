package main

import (
	"fmt"
	"math/big"
)

func bigCalc(a *big.Float, b *big.Float) {
	fmt.Printf("First numb = %f\nSecond numb = %f\n", a, b)

	plus := new(big.Float) //oператор +
	fmt.Println("Operation + :", plus.Add(a, b))

	sub := new(big.Float) //oператор -
	fmt.Println("Operation - :", sub.Sub(a, b))

	div := new(big.Float) //oператор /
	fmt.Println("Operation / :", div.Quo(a, b))
}

func main() {
	a := big.NewFloat(10855760.0504)
	b := big.NewFloat(20555760.8503)
	
	bigCalc(a, b)
}
