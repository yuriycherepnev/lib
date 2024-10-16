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
	for i := 0; i < len(salary)-1; i++ {
		if salary[i] < salary[i+1] {
			return salary[:i] + salary[i+1:]
		}
	}
	return salary[:len(salary)-1]
}

/*
	При сравнении строк в Go происходит побайтовое сравнение (основанное на значениях байтов в UTF-8).
	Так как в нашем случае цифры в зарплате представлены строками (например, "5", "3", "6"),
	их сравнение работает как побайтовое сравнение ASCII кодов этих символов.
	Когда мы сравниваем символы (цифры) как строки, Go сравнивает их значения в таблице ASCII, и это дает корректный результат для цифр.
	'0' = 48, '1' = 49 и т.д.
	Это работает корректно только для цифр, так как они упорядочены в таблице ASCII последовательно,
	но в общем случае для других символов такое поведение может привести к непредсказуемым результатам.
*/
