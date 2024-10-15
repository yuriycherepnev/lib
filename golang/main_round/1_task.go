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
			salary := "0"
			if len(line) > 1 {
				salary = cutSalary(line)
			}
			fmt.Println(salary)
			lineCount++
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

func cutSalary(salary string) string {
	cutSalary := ""
	sliceSalary := make([]string, 0, utf8.RuneCountInString(salary))
	minNumber := 9
	minIndex := 0

	fmt.Printf("len: %d, cap: %d\n", len(sliceSalary), cap(sliceSalary))

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
