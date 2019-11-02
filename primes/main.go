package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	// integers go in, primes come out
	integers := make(chan int, 1000000)
	primes := make(chan int, 1000000)

	// I don't know how many cores I have
	go calculatePrimes(integers, primes)
	go calculatePrimes(integers, primes)
	go calculatePrimes(integers, primes)
	go calculatePrimes(integers, primes)

	// Shovel integers into the input buffer
	go func() {
		for i := 3; ; i += 2 {
			integers <- i
		}
	}()

	// Print the output buffer every second or so
	go func() {
		i := 0
		for {
			select {
			case prime := <-primes:
				fmt.Println(i, prime)
				i++
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(1 * time.Minute)
}

func calculatePrimes(integers chan int, primes chan int) {
	for {
		i := <-integers
		if isPrime(i) {
			primes <- i
			// fmt.Println(i)
		}
	}
}

func isPrime(arg int) bool {
	// checks odd factors up to sqrt(arg)
	if arg%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Ceil(math.Sqrt(float64(arg)))); i += 2 {
		if arg%i == 0 {
			return false
		}
	}
	return true
}
