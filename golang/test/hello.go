package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	textNumber := strconv.Itoa(123)
	numberLength := utf8.RuneCountInString(textNumber)
	fmt.Println(numberLength)
}

func main() {
	salary := 990778866
	sliceSalary, minNumberIndex := IntToSlice(salary, make([]int, 0))
	newSlice := append(sliceSalary[:minNumberIndex], sliceSalary[minNumberIndex+1:]...)

	fmt.Println(newSlice)
}

func IntToSlice(number int, numberCount int) ([]int, int) {
	numberSlice = make([]int, 0, numberCount)

	for number > 0 {
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

func splitInt(n int) []int {
	slc := []int{}

	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	return slc
}
