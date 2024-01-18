package practiceExercises

import (
	"fmt"
	"math/rand"
)

func PracticeExercises() {
	// fmt.Printf("return random func: %v \n", ramdom())
	// _, _ = ramdomTwo()
	// loops()
	//ramdomThree()
	//loopsTwo()
	//loopsThree()
	loopsFour()
}

// Create a random number between 0 a 250
func ramdom() int {
	a := rand.Intn(250)
	switch {
	case a <= 100:
		fmt.Printf("the value of a is between 0 and 100 switch : %v\n", a)
	case a > 100 && a <= 200:
		fmt.Printf("the value of a is between 101 and 200 switch: %v\n", a)
	default:
		fmt.Printf("the value of a is between 201 and 250 switch: %v\n", a)
	}

	if a >= 0 && a <= 100 {
		fmt.Printf("the value of a is between 0 and 100 : %v\n", a)
	}
	if a > 100 && a <= 200 {
		fmt.Printf("the value of a is between 101 and 200: %v\n", a)
	}
	// this is because the rand.Intn() doesn't include the right extreme
	if a >= 200 && a < 250 {
		fmt.Printf("the value of a is between 201 and 250: %v\n", a)
	}
	return a
}

func ramdomTwo() (int, int) {
	a := rand.Intn(10)
	b := rand.Intn(10)

	switch {
	case a < 4 && b < 4:
		fmt.Printf("the variables values are both less than 4 switch: %v,%v:\n", a, b)
	case a > 6 && b > 6:
		fmt.Printf("the variables values are both bigger than 6 switch: %v,%v:\n", a, b)
	case a >= 4 && a <= 6:
		fmt.Printf("the variable a value is bigger than 4 and less than 6 switch: %v,%v:\n", a, b)
	case b != 5:
		fmt.Printf("the variable b is not 5 switch: %v\n", b)
	default:
		fmt.Printf("none found, values are: %v,%v \n", a, b)
	}

	if a < 4 && b < 4 {
		fmt.Printf("the variables values are both less than 4: %v,%v:\n", a, b)
	} else if a > 6 && b > 6 {
		fmt.Printf("the variables values are both bigger than 6: %v,%v:\n", a, b)
	} else if a >= 4 && a <= 6 {
		fmt.Printf("the variable a value is bigger than 4 and less than 6: %v,%v:\n", a, b)
	} else if b != 5 {
		fmt.Printf("the variable b is not 5: %v\n", b)
	} else {
		fmt.Printf("none found, values are: %v,%v \n", a, b)
	}

	return a, b
}

func loops() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func ramdomThree() {
	for i := 0; i < 43; i++ {
		x := rand.Intn(5)
		switch {
		case x < 3:
			fmt.Printf("the variable es less than 3, itiration %v: %v\n", i, x)
		default:
			fmt.Printf("the variable is greater or iqual than three, itiration %v: %v\n", i, x)
		}
	}
}

func loopsTwo() {
	i := 0
	for i < 20 {
		if i == 13 {
			break
		} else if i == 8 {
			continue
		} else {
			fmt.Println(i)
			i++
		}
	}
}

func loopsThree() {
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Printf("this is the outer loop %v, and the inner loop is %v \n", i, j)
		}
	}
}

func loopsFour() {
	xi := []int{42, 43, 44, 45, 46, 47}

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for i, v := range xi {
		fmt.Printf("this is the value, %v in the index %v of xi array\n", v, i)
	}

	fmt.Println("----------")

	for k, v := range m {
		fmt.Printf("this is the key, %v an the value %v of m map\n", k, v)
	}

	fmt.Println("----------")

	if v, ok := m["Q"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("the Q doesn't exits in the m map")
	}

	fmt.Println("----------")

	count := 0
	for i := 0; i <= 50; i++ {
		if x := rand.Intn(5); x == 3 {
			count++
			fmt.Printf("x is three #%v\n", count)
		}
	}

	fmt.Println("----------")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
