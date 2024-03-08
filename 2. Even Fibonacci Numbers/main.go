package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	before := time.Now()
	firstNumber := 0
	secondNumber := 1
	sumEven := 0
	sumOdd := 0

	const max = 4_000_000
	for secondNumber < max {
		buff := firstNumber + secondNumber
		if secondNumber%2 == 0 {
			fmt.Println("Bilanganya genap dari FIB :", secondNumber)
			sumEven += secondNumber
		} else {
			fmt.Println("BIlanganya ganjil dari FIB :", secondNumber)
			sumOdd += secondNumber
		}

		firstNumber = secondNumber
		secondNumber = buff
	}

	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Jumlah dari bilangan ganjil : [%d]\n", sumOdd)
	fmt.Printf("Jumlah dari bilangan genap : [%d]\n", sumEven)
	after := time.Now()
	fmt.Println("Waktu eksekusi :", (after.Nanosecond() - before.Nanosecond()), "Nano second")
}
