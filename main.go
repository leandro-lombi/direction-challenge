package main

import (
	"fmt"

	directionReduce "github.com/direction-challenge/directionReduce"
)

func main() {

	// list: ["NORTH","SOUTH","SOUTH","EAST","WEST","NORTH","WEST"]
	// result: ["WEST"]

	// list: ["NORTH","SOUTH","SOUTH","EAST","WEST","NORTH"]
	// result: []

	// list: ["NORTH","WEST","SOUTH","EAST"]
	// result: ["NORTH","WEST","SOUTH","EAST"]

	list := []string{"NORTH", "WEST", "SOUTH", "EAST"}
	result := directionReduce.DirReduc(list)

	fmt.Println(result)
}
