//необходимо найти самое маленькое число на самой высокой позиции и убрать его
/*
это можно сделать за один проход по числу
идем по цифрам числа.
Меньшее кладем в переменную и запоминаем его индекс
Если встречаем число меньше - обновляем даннные
по окончанию итерирования - удаляем число по индексу, склеиваем массив в число и возвращаем
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	var data *bufio.Reader
	data = bufio.NewReader(os.Stdin)
	var dataLength int = 1
	var lineCount int = 0

	for i := 1; true; i++ {
		line, err := data.ReadString('\n')
		line = strings.TrimSpace(line)

		if i > 1 {
			var salary int = 0
			if len(line) > 1 {
				salaryNumber, _ := strconv.Atoi(line)
				salary = cutSalary(salaryNumber)
			}
			lineCount++
			fmt.Println(salary)
		}

		if i == 1 {
			number, _ := strconv.Atoi(line)
			dataLength = number
		}

		if lineCount == dataLength {
			break
		}

		if err == io.EOF {
			break
		}
	}
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
