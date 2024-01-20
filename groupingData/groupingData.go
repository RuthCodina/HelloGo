package groupingData

import "fmt"

func GroupingData() {
	//arraies()
	//slice()
	//arraiesTwo()
	arraiesThree()
}

func arraies() {
	//literal array
	arr := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "BittersweetChocolate (GF)", "Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)", "Browned ButterCookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
		"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "MadagascarVanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy ChocolatePeanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy SunbutterCinnamon (GF, V)",
		"Orange Cream (GF) ", "Peanut Butter Cookie Dough", "RaspberrySorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different",
		"Strawberry LemonadeSorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "WolverineTracks (GF)"}

	fmt.Printf("array length: %v %T", len(arr), arr)
}

func slice() {
	s := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "BittersweetChocolate (GF)", "Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)", "Browned ButterCookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
		"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "MadagascarVanilla (GF)", "Matcha (GF)",
		"Midnight Chocolate Sorbet (GF, V)", "Non-Dairy ChocolatePeanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)",
		"Non-Dairy SunbutterCinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "RaspberrySorbet (GF, V)", "Salty Caramel (GF)",
		"Slate Slate Different", "Strawberry LemonadeSorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "WolverineTracks (GF)"}

	for i, v := range s {
		fmt.Printf("index %v to the value %v\n", i, v)
	}

	s = append(s, "added value")
}

func arraiesTwo() {
	a := [5]int{5, 10, 15, 20, 25}
	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range a {
		fmt.Printf("the value %v, in the index %v, with type %T\n", i, v, v)
	}
	for i, v := range b {
		fmt.Printf("the value %v, in the index %v, with type %T in slice\n", i, v, v)
	}
	b1 := b[0:5]
	fmt.Println("the b1 is: ", b1)

	b2 := b[5:]
	fmt.Println("the b2 is: ", b2)

	b3 := b[2:7]
	fmt.Println("the b3 is: ", b3)

	b4 := b[1:6]
	fmt.Println("the b4 is: ", b4)
}

func arraiesThree() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	fmt.Println("the x slice is: ", x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println("the x slice is: ", x)

	fmt.Println("----------")

	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x1 := x[0:3]
	fmt.Println("the x1 slice is: ", x1)
	x2 := x[6:]
	fmt.Println("the x2 slice is: ", x2)
	z := []int{}
	z = append(z, x1...)
	z = append(z, x2...)
	fmt.Println("the z slice is: ", z)

	fmt.Println("----------")

	a := make([]string, 0)
	a = append(a, " Alabama", " Alaska", " Arizona", " Arkansas", " California", " Colorado", " Connecticut", "Delaware",
		" Florida", " Georgia", " Hawaii", " Idaho", " Illinois", " Indiana", " Iowa", " Kansas", "Kentucky", " Louisiana", " Maine",
		" Maryland", " Massachusetts", " Michigan", " Minnesota", "Mississippi", " Missouri", " Montana", " Nebraska", " Nevada", " New Hampshire", " New Jersey",
		" New Mexico", " New York", " North Carolina", " North Dakota", " Ohio", " Oklahoma", " Oregon",
		" Pennsylvania", " Rhode Island", " South Carolina", " South Dakota", " Tennessee", " Texas", "Utah", " Vermont",
		" Virginia", " Washington", " West Virginia", " Wisconsin", " Wyoming")

	for i, v := range a {
		fmt.Printf("this is the s slice, with cap %v, len %v, and the value in the index %v, is %v\n", cap(a), len(a), i, v)
	}

	fmt.Println("----------")

	b := make([][]string, 0)

	b = append(b, []string{"James", "Bond", "Shaken, not stirred"})
	b = append(b, []string{"Miss", "Moneypenny", "I'm 008."})

	for _, v := range b {
		fmt.Printf("this is the value %v\n", v)
	}

}
