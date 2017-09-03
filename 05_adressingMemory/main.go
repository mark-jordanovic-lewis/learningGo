package main

import (
	"fmt"
)

const metersToYards float64 = 1.090361

func main() {
	a := 42
	fmt.Println("a = ", a)
	fmt.Println("addr of a = ", &a)
	fmt.Printf("addr of a = %d\n", &a)
	fmt.Println()
	// get inout from user and look at memory addr
	var meters float64                       // assigning a type to meters
	fmt.Println("How far away is the shop?") // prompt user
	fmt.Scan(&meters)                        // get user input, and assign to memory location of meters, casting to type.
	yards := meters * metersToYards
	fmt.Println(meters, " m, or ", yards, "yards")
}
