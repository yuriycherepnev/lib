//с числом работаем как со строкой!
//в типе инт не хватает длины цифры для обработки всего числа

package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	salary := "2003004005678"
	fmt.Println(salary)
	cutSalary := cutSalary(salary)
	fmt.Println(cutSalary)
}

func cutSalary(salary string) string {
	cutSalary := ""
	sliceSalary := make([]string, 0, utf8.RuneCountInString(salary))
	minNumber := 9
	minIndex := 0

	for index, value := range salary {
		stringNumber := string(value)
		currentNumber, _ := strconv.Atoi(stringNumber)
		sliceSalary = append(sliceSalary, stringNumber)

		if currentNumber < minNumber {
			minNumber = currentNumber
			minIndex = index
		}
	}

	sliceSalary = append(sliceSalary[:minIndex], sliceSalary[minIndex+1:]...)

	for _, value := range sliceSalary {
		cutSalary += value
	}

	return cutSalary
}
