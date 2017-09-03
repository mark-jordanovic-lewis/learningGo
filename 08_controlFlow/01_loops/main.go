package main

import "fmt"

func main() {
	fmt.Println("Std for loop")
	for i := 0; i < 4; i++ {
		fmt.Println("i is ", i)
	}
	fmt.Println()
	fmt.Println("nested for loop")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println("i is", i, ", j is", j)
		}
	}
	fmt.Println()
	fmt.Println("while like for loop")
	someBooleanExpression := true
	i := 0
	for someBooleanExpression {
		fmt.Println("i is ", i)
		if i == 3 {
			someBooleanExpression = false
		}
		i++
	}
	fmt.Println()
	fmt.Println("infinite while like for loop with break condition")
	i = 0
	for {
		fmt.Println("i is ", i)
		if i == 3 {
			break
		}
		i++
	}
	fmt.Println()
	fmt.Println("Print only evens with loop continue")
	i = 0
	for {
		if i == 6 {
			break
		}
		i++
		if i%2 == 1 {
			continue
		}
		fmt.Println("i is ", i)
	}
}
