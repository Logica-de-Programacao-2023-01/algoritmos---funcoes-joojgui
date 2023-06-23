package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var num int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&num)

	sum, err := sumDigits(num)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Soma dos dígitos:", sum)
	}
}

func sumDigits(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("Número negativo não é válido")
	}

	sum := 0

	for num != 0 {
		digit := num % 10
		sum += int(math.Abs(float64(digit)))
		num /= 10
	}

	return sum, nil
}
