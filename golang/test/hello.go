package main

import "fmt"

func main() {
	salary := 563345
	intSlice, minNumber, minNumberIndex := IntToSlice(salary, make([]int, 0))

	fmt.Println(intSlice)
	fmt.Println(minNumber)
	fmt.Println(minNumberIndex)
}

func IntToSlice(salary int, intSlice []int) ([]int, int, int) {
	minNumber := 9
	reverseIndex := 1
	currentIndex := 1

	for salary > 0 {
		currentNumber := salary % 10
		intSlice = append([]int{currentNumber}, intSlice...)
		salary = salary / 10

		if currentNumber <= minNumber {
			minNumber = currentNumber
			reverseIndex = currentIndex
		}
		if salary != 0 {
			currentIndex++
		}
	}

	return intSlice, minNumber, currentIndex - reverseIndex
}
