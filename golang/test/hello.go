//проблема reverse number!
//скорее всего самый удобный способ - привести число к строке и через итерирование строки заполнять slice

package main

import (
	"strconv"
	"unicode/utf8"
)

func main() {
	salary := 990778876
	textNumber := strconv.Itoa(salary)
	numberCount := utf8.RuneCountInString(textNumber)
	sliceSalary := make([]int, 0, numberCount)

	for _, currentTextNumber := range textNumber {
		currentNumber, _ := strconv.Atoi(string(currentTextNumber))
		sliceSalary = append(sliceSalary, currentNumber)
	}
}

func getSliceSalary(salary int) ([]int, int) {
	textNumber := strconv.Itoa(salary)
	sliceSalary := make([]int, 0, utf8.RuneCountInString(textNumber))
	minNumber := 9
	minIndex := 0

	for _, currentTextNumber := range textNumber {
		currentNumber, _ := strconv.Atoi(string(currentTextNumber))
		sliceSalary = append(sliceSalary, currentNumber)

		if currentNumber < minNumber {
			minNumber = currentNumber

			minIndexя
		}
	}

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
