package main

import "fmt"

func main() {
	salary := 990778866
	sliceSalary, minNumberIndex := IntToSlice(salary, make([]int, 0))
	newSlice := append(sliceSalary[:minNumberIndex], sliceSalary[minNumberIndex+1:]...)

	fmt.Println(newSlice)
}

func IntToSlice(salary int, sliceSalary []int) ([]int, int) {
	minNumber := 9
	reverseIndex := 1
	currentIndex := 1

	for salary > 0 {
		currentNumber := salary % 10
		sliceSalary = append([]int{currentNumber}, sliceSalary...)
		salary = salary / 10

		if currentNumber <= minNumber {
			minNumber = currentNumber
			reverseIndex = currentIndex
		}
		if salary != 0 {
			currentIndex++
		}
	}

	return sliceSalary, currentIndex - reverseIndex
}
