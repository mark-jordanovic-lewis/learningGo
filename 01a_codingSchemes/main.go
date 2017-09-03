package main

import "fmt"

func main() {
	fmt.Printf("decimal: %d\n", 42)      // print decimal number
	fmt.Printf("binary: %b\n", 42)       // print formatted binary
	fmt.Printf("octal: %#o\n", 42)       // print formatted octal
	fmt.Printf("hexadecimal: %#x\n", 42) // print formatted hexadecimal
	fmt.Println()
	fmt.Printf("decimal: %d\n", 7035)      // print decimal number
	fmt.Printf("binary: %b\n", 7035)       // print formatted binary
	fmt.Printf("octal: %#o\n", 7035)       // print formatted octal
	fmt.Printf("hexadecimal: %#x\n", 7035) // print formatted hexadecimal

	fmt.Println("Looping through and formatting numbers")
	fmt.Println("decimal = binary = octal = hexadecimal = utf-8")
	for i := 0; i < 127; i++ {
		fmt.Printf("%d = %b = %#o = %#x = %q\n", i, i, i, i, i)
	}

}
