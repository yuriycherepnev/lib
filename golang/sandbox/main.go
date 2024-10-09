package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	salary := 30101023
	fmt.Println(salary)
	cutSalary := cutSalary(salary)
	fmt.Println(cutSalary)
}

func cutSalary(salary int) int {
	cutSalary := 0
	multiplier := 10
	textNumber := strconv.Itoa(salary)
	sliceSalary := make([]int, 0, utf8.RuneCountInString(textNumber))
	minNumber := 9
	minIndex := 0
	currentIndex := 0

	for salary > 0 {
		currentNumber := salary % 10
		sliceSalary = append(sliceSalary, currentNumber)
		salary = salary / 10

		if currentNumber <= minNumber {
			minNumber = currentNumber
			minIndex = currentIndex
		}
		if salary != 0 {
			currentIndex++
		}
	}

	sliceSalary = append(sliceSalary[:minIndex], sliceSalary[minIndex+1:]...)

	for index, value := range sliceSalary {
		if index == 0 {
			multiplier = 1
		} else {
			multiplier = multiplier * 10
		}
		cutSalary = value*multiplier + cutSalary
	}

	return cutSalary
}
