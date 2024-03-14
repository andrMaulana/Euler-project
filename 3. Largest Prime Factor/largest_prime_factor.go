package main

import (
	"fmt"
	"math"
	"time"
)

func printPrime(n int) {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++{
		if n%i == 0 {
			fmt.Println(i)
			n /= i
			i--
		}
	}
	if n > 0 {
		fmt.Println(n)
	}
}



func main () {
	before := time.Now()
	n := 13195
	printPrime(n)
	after := time.Now()
	fmt.Println("Waktu eksekusi ", after.Nanosecond()-before.Nanosecond(), "nano second")
}
