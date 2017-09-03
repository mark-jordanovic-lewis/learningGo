package main

import (
	"fmt"
)

type special struct {
	name   string
	number int
}

func main() {
	// Arrays
	// declare new string array of length 58, all initialized to ""
	var stringArray [8]string
	fmt.Println("unpopulated stringArray:", stringArray, "has length:", len(stringArray))
	for i := 65; i < 65+len(stringArray); i++ {
		stringArray[i-65] = string(i)
	}
	fmt.Println("populated stringArray:", stringArray)
	for i, elem := range stringArray {
		fmt.Printf("element#%v: %T - %v - %b\n", i, elem, elem, []rune(elem))
	}
	// can put any type in an Array
	var specialArray [3]special
	specialArray[0] = special{"bob", 32}
	specialArray[1] = special{"alice", 28}
	specialArray[2] = special{"janine", 2}
	for i, elem := range specialArray {
		fmt.Printf("element#%v = %v\n", i, elem)
	}
	// Slices
	// make slice with 0 elements of capacity 5
	slice := make([]int, 0, 1)
	// capacity increase doubles each time a new array is needed
	for i := 0; i <= 8; i++ {
		slice = append(slice, i+1)
		fmt.Println("slice props: length =", len(slice), "capacity =", cap(slice), "n inserted:", slice[i])
	}
}
