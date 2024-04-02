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

	defering()
	p := person{
		first: "Romina",
		age:   60,
	}
	h := p.speak()
	fmt.Println(h)

	a := func() {
		fmt.Printf("anonymous Func, that show person info:  %v\n", h)
	}
	a()

	b := ask()

	fmt.Printf("this is the ask func response %v\n", b())
	fmt.Println("--------------------")

	printTripleAge(p.tripleAge)
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

func ask() func() int {
	return func() int {
		return 42
	}
}

func (p person) tripleAge() int {
	return p.age * 3
}

func printTripleAge(cb func() int) {
	fmt.Printf("This is the person age tripled: %v", cb())
}
