package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	number := 10

	primes, err := getPrimes(number)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números primos menores ou iguais a", number, ":", primes)
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func getPrimes(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("O número é negativo")
	}

	primes := make([]int, 0)

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes, nil
}
