//проблема reverse number!
//скорее всего самый удобный способ - привести число к строке и через итерирование строки заполнять slice

package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

func main() {
	salary := 9311377
	cutSalary := cutSalary(salary)
	fmt.Println(cutSalary)
}

func cutSalary(salary int) int {
	cutSalary := 0
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

	fmt.Println(sliceSalary)
	for index, value := range sliceSalary {
		if cutSalary == 0 {
			cutSalary = value
		} else {
			factor := math.Pow(10, float64(index))
			cutSalary += value * int(factor)
		}
	}

	return cutSalary
}
