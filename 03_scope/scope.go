package main

import "fmt"

// variables declared out side of any closure have package scope
// scope is cross files of the same package!!

// declaration of r, the string
var r string

// Bob is a visible package variable (it's callable) as it's capitalized
var Bob int

// declaration and assignment of s the int, type can be inferred from RHS.
var s int = 52

// declaration and

func main() {
	// assignment of previously declared variable, r has package scope!
	r = "beebop"
	fmt.Printf("the s variable is in scope: %d\n", s)

	// variables declared in a func have function scope, but ...
	// here is a closure, no loop, nothing, just a closure declared with braces.
	{
		// declaring, initializing and assigning using shorthand declaration
		a := 1
		b := "Zaphod"
		c := 3.14
		d := true

		fmt.Printf("a is %d and is a %T\n", a, a)
		fmt.Printf("b is %s and is a %T\n", b, b)
		fmt.Printf("c is %f and is a %T\n", c, c)
		fmt.Printf("d is %t and is a %T\n", d, d)
		fmt.Printf("percent v prints out all type values: %v, %v, %v, %v\n", a, b, c, d)
	}
	// uncommenting the line below causes non-compilation as the vars are not in scope!!
	// fmt.Printf("percent v prints out all type values: %v, %v, %v, %v\n", a, b, c, d)

	// declaring to zero value of type (all types have a zero value)
	var e int // or int8, int16, int32, int64
	var f string
	var g float64 // or float32
	var h bool
	// just a quick example of how closure allows var INTO it
	{
		fmt.Printf("e is %d and is a %T\n", e, e)
		fmt.Printf("f is %s and is a %T\n", f, f)
		fmt.Printf("g is %f and is a %T\n", g, g)
	}
	fmt.Printf("h is %t and is a %T\n", h, h)
	fmt.Printf("percent v prints out all type values: %v, %v, %v, %v\n", e, f, g, h)

	// anonymous functions are declared as so: (all scope rules count)
	myAnon := func() int {
		e = 12
		return e
	}
	fmt.Printf("e is %d and is a %T, it has not been modified.\n", e, e)
	// function call in the closure modifying variable in function scope
	{
		myAnon()
	}
	// modifies the variable.
	fmt.Printf("e is %d and is a %T, it has been modified.\n", e, e)
	fmt.Print("counting... ")
	increment := wrapper()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", increment())
	}
	fmt.Println()
	// can initialize variables on an if
	fmt.Println("Initialize closure variable on an if statement")
	if hi := "Hello"; true {
		fmt.Println(hi, "from in the if")
	}
	// fmt.Println(hi, "will throw a compile error")
}

// functions can return functions
// declared below but still in package scope so available inside all code in pacakage
// note, we do not have access to x, just the return value of the anonymous function.
func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
