package main

import (
	"fmt"

	directionReduce "github.com/direction-challenge/directionReduce"
)

func main() {
	list := []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
	result, err := directionReduce.DirReduc(list)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	fmt.Println(result)
}
