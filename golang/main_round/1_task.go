//необходимо найти самое маленькое число на самой высокой позиции и убрать его
/*
Надо взять все цифры в числе, найти из них минимальное
и получить его индекс

Удалить цифру по индексу и вернуть число

Алгоритм:
перевести число в массив
найти мин цифру в массиве
найти ее индекс
удалть эту цифру из массива
перевести массив в число
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	return salary
}
