package main

import (
	"fmt"
	"math"
)

func print_primes(n int) {
	var not_prime = make([]bool, n)

	sq := int(math.Sqrt(float64(n)))

	// Mark all multiples as not primes.
	for i := 2; i <= sq; i++ {
		// If i is already marked as not prime then all its multiples have been
		// marked already too.
		if not_prime[i] {
			continue
		}

		// Start at i squared.
		j := i * i

		for j < n {
			not_prime[j] = true
			j += i
		}
	}

	// Print all the numbers that we have not ruled out.
	for i := 2; i < n; i++ {
		if !not_prime[i] {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Print("\n")
}

func main() {
	print_primes(1000)
}
