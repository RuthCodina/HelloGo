package functions

import (
	"fmt"
)

func Functions() {
	x := foo([]int{7, 8, 9, 10, 11, 12, 13}...)
	y := bar([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("--------------------")

	//	defering()
	p := person{
		first: "Romina",
		age:   60,
	}
	h := p.speak()
	fmt.Println(h)

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

func defering() {
	defer fmt.Println("defering")
	fmt.Println("caminar la vida")
	fmt.Println("Maravillosa")
}

type person struct {
	first string
	age   int
}

func (p person) speak() string {
	str := fmt.Sprintf("Hi, mi name is %v, and I have %v", p.first, p.age)
	return str
}
