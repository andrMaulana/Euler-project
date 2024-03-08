package main

import (
	"fmt"
	"time"
)

func main() {
	before := time.Now()
	sumOf3 := 0
	sumOf5 := 0

	for i := 0; i < 1_000; i++ {
		if i%5 == 0 {
			fmt.Println("Data perelement adalah :", "Buzz")
			sumOf5 += i
		} else if i%3 == 0 {
			fmt.Println("data perelement adalah :", "Fizz")
			sumOf3 += i
		} else {
			fmt.Println("Data perelement adalah :", i)
		}
	}

	fmt.Println("Jumlah data merupakan kelipatan 3 adalah :", sumOf3)
	fmt.Println("Jumlah data merupakan kelipatan 5 adalah :", sumOf5)
	fmt.Println("Jumlah total kelipatan 3 dan 5 :", sumOf5+sumOf3)

	after := time.Now()
	fmt.Println("Waktu eksekusi ", after.Nanosecond()-before.Nanosecond(), "Nano second")
}
