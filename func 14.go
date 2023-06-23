package main

import (
	"errors"
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 5, 6}

	intersection, err := findIntersection(slice1, slice2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Interseção:", intersection)
	}
}

func findIntersection(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Um dos slices está vazio")
	}

	intersection := make([]int, 0)

	numbersMap := make(map[int]bool)
	for _, num := range slice1 {
		numbersMap[num] = true
	}

	for _, num := range slice2 {
		if numbersMap[num] {
			intersection = append(intersection, num)
		}
	}

	return intersection, nil
}
