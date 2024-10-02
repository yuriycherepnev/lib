//проблема reverse number!

package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	salary := 990778876
	sliceSalary, minNumberIndex := getSliceSalary(salary)
	newSlice := append(sliceSalary[:minNumberIndex], sliceSalary[minNumberIndex+1:]...)

	fmt.Println(newSlice)
}

func getSliceSalary(salary int) ([]int, int) {
	textNumber := strconv.Itoa(123)
	numberCount := utf8.RuneCountInString(textNumber)
	sliceSalary := make([]int, 0, numberCount)
	minNumber := 9
	currentIndex := 0
	reverseIndex := 0

	for salary > 0 {
		currentNumber := salary % 10
		sliceSalary = append(sliceSalary, currentNumber)
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
