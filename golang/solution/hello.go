package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(cutSalary(9123))
}

func cutSalary(salary int) int {
	salarysArray := make([]int, 0, salary)
	splitSalary := splitInt(salary)
	for index, _ := range splitSalary {
		salarysArray = append(salarysArray, deleteNumber(splitSalary, index))
	}
	sort.Ints(salarysArray)
	return salarysArray[len(salarysArray)-1]
}

func deleteNumber(splitSalary []int, numberToRemove int) int {
	copy(splitSalary[numberToRemove:], splitSalary[numberToRemove+1:])
	splitSalary = splitSalary[:len(splitSalary)-1]

	return sliceToInt(splitSalary)
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}

	var reverse []int

	for i := len(slc) - 1; i >= 0; i-- {
		reverse = append(reverse, slc[i])
	}

	return reverse
}
