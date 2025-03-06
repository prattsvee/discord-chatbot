package main

import (
	"fmt"
	"time"
)

func main() {
	// find prime numbers between 1 and 100
	//print start time

	fmt.Println("time start", time.Now().Format("2006-01-02 15:04:05"))

	//print end time

	for i := 1; i <= 10000000; i++ {
		if isPrime(i) {
			println(i)
		}
	}

	fmt.Println("time end", time.Now().Format("2006-01-02 15:04:05"))
}
func isPrime(i int) bool {
	if i < 2 {
		return false
	}
	for j := 2; j < i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}
