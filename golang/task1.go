package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main1() {
	var data *bufio.Reader
	data = bufio.NewReader(os.Stdin)
	var dataLength int = 1
	var lineCount int = 0
	var salary int = 0

	for i := 1; true; i++ {
		line, err := data.ReadString('\n')
		line = strings.TrimSpace(line)

		if i > 1 {
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
	salarysArray := make([]int, 0, salary)
	for index, _ := range strconv.Itoa(salary) {
		salarysArray = append(salarysArray, deleteNumber(salary, index+1))
	}
	sort.Ints(salarysArray)
	return salarysArray[len(salarysArray)-1]
}

func deleteNumber(salary, numberToRemove int) int {
	var result, divisor int
	divisor = 1

	for i := 1; salary != 0; i++ {
		digit := salary % 10

		if i != numberToRemove {
			result += digit * divisor
			divisor *= 10
		}

		salary /= 10
	}

	return result
}
