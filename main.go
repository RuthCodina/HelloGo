package main

import (
	"fmt"

	"github.com/RuthCodina/HelloGo/functions"
	"github.com/RuthCodina/HelloGo/groupingData"
	"github.com/RuthCodina/HelloGo/practiceExercises"
)

func main() {
	fmt.Println("This is where initialization for my program occurs")
	practiceExercises.PracticeExercises()
	groupingData.GroupingData()
	functions.Functions()
	functions.Interfaces()
}
