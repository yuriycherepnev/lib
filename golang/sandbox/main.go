//с числом работаем как со строкой!
//в типе инт не хватает длины цифры для обработки всего числа

package main

import (
	"fmt"
	"strconv"
)

func main() {
	salary := 12345678
	fmt.Println(salary)
	cutSalary := cutSalary(salary)
	fmt.Println(cutSalary)
}

func cutSalary(salary int) int {
	textNumber := strconv.Itoa(salary)
	//sliceSalary := make([]string, 0, utf8.RuneCountInString(textNumber))
	//minNumber := 9
	//minIndex := 0
	//currentIndex := 0

	for i, v := range textNumber {
		fmt.Println(i)
		fmt.Println(v)
	}

	//for salary > 0 {
	//	currentNumber := salary % 10
	//	sliceSalary = append(sliceSalary, currentNumber)
	//	salary = salary / 10
	//
	//	if currentNumber <= minNumber {
	//		minNumber = currentNumber
	//		minIndex = currentIndex
	//	}
	//	if salary != 0 {
	//		currentIndex++
	//	}
	//}
	//
	//sliceSalary = append(sliceSalary[:minIndex], sliceSalary[minIndex+1:]...)

	return 10
}
