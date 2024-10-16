package main

import (
	"fmt"
)

func main() {
	salary := "2003004005678"
	fmt.Println(salary)
	cutSalary := cutSalary(salary)
	fmt.Println(cutSalary)
}

func cutSalary(salary string) string {
	for i := 0; i < len(salary)-1; i++ {
		if salary[i] < salary[i+1] {
			return salary[:i] + salary[i+1:]
		}
	}
	return salary[:len(salary)-1]
}
