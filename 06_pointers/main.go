package main

// Go passes arguments by value not by reference so using pointers is important.

import "fmt"

func main() {
	a := 42
	fmt.Printf("a is a %T, has value: %v\n", a, a)
	fmt.Println("a is at addr = ", &a)
	fmt.Println()
	// declare a pointer to an int at the mem addr of a (this can be inferred)
	var b *int = &a
	fmt.Printf("b is a %T, has value: %v\n", b, b)
	fmt.Printf("b points to value: %v\n", *b)
	fmt.Println()
	// reassigning the value of a through b
	fmt.Println("reassigning a through *b")
	*b = 128
	fmt.Printf("*b = %v ahould be the same as a = %v\n", *b, a)
	fmt.Println()
	// testing value passing zeroing function
	x := 5
	fmt.Println("x =", x)
	fmt.Println("addr of var x = ", &x)
	fmt.Println("Trying to zero by passing to zeroer by value ")
	zeroNoPtr(x)
	fmt.Println("x is now =", x)
	fmt.Println()
	// testing ref passing zeroing function (still passing by value but the value is the ref addr)
	y := 5
	fmt.Println("y =", y)
	fmt.Println("addr of var y = ", &y)
	fmt.Println("Trying to zero by passing to zeroer by reference ")
	zeroWPtr(&y)
	fmt.Println("y is now =", y)

}

func zeroNoPtr(x int) {
	fmt.Println("addr of arg x = ", &x)
	x = 0
}

func zeroWPtr(x *int) {
	fmt.Println("arg x points to addr ", x)
	*x = 0
}
