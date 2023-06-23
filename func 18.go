package main

import (
	"errors"
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	sum, err := applyFunction(numbers, square)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Soma dos resultados:", sum)
	}
}

func square(n int) int {
	return n * n
}

func applyFunction(numbers []int, fn func(int) int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("O slice está vazio")
	}

	sum := 0

	for _, num := range numbers {
		result := fn(num)
		sum += result
	}

	return sum, nil
}
