package main

import (
	"fmt"
)

func main() {
	grades := make(map[string]float32)

	grades["Mike"] = 44
	grades["Emily"] = 20
	grades["Einstein"] = 97

	fmt.Println(grades)

	mikeScore := grades["Mike"]

	fmt.Println(mikeScore)
	delete(grades, "Einstein")

	for k, v := range grades {
		fmt.Printf("%s has scored %v\n", k, v)
	}
}
