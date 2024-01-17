package practiceExercises

import (
	"fmt"
	"math/rand"
)

func PracticeExercises() {
	fmt.Printf("return random func: %v \n", ramdom())
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
