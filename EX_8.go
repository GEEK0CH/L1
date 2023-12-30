package main

import "fmt"

func setBit(numb int64, p, oOr1 int) int64 {
	if oOr1 == 1 { // обрабока в случае бита 1
		return numb | (1 << p)
	}
	return numb &^ (1 << p) // в противном случае обработать
}

func main() {
	var numb int64 = 15              // 15=1111
	var numb2 int64 = 14             // 14=1110
	fmt.Println(setBit(numb, 0, 0))  // 14 = 1110
	fmt.Println(setBit(numb, 1, 0))  // 13 = 1101
	fmt.Println(setBit(numb2, 0, 1)) // 15 = 1111
	fmt.Println(setBit(numb2, 2, 0)) // 10 = 1010
}
