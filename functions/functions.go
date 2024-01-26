package functions

import "fmt"

func Functions() {
	x := foo([]int{7, 8, 9, 10, 11, 12, 13}...)
	y := bar([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
	fmt.Println(y)
}

func foo(array ...int) int {
	count := 0
	for _, v := range array {
		count += v
	}
	return count
}

func bar(array []int) int {
	count := 0
	for _, v := range array {
		count += v
	}
	return count
}
