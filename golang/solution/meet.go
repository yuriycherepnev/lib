package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func meet() {
	foo, bar := 42, 314
	bar, bazz := 31, 7

	fmt.Println(foo, bar, bazz)

	for i, v := range pow {
		fmt.Println(i, v)
	}
	some1()

	foo1 := 42
	if foo2, aloha := addIntAndString(60, 1, "aloha", "kobanos"); foo2 == 61 {
		fmt.Println(foo1)
		fmt.Println(foo2)
		fmt.Println(aloha)
	}
}

var foo int = 34

func some1() {
	// because foo here is scoped to some func
	foo := 42 // <-- legal
	foo = 314 // <-- legal

	fmt.Println(foo, foo)
}

func add1(x int, y int) int {
	return x + y
}

func addIntAndString(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}
