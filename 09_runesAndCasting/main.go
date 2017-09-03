package main

import "fmt"

func main() {
	a := 'a'
	fmt.Printf(
		"a is of type %T, has value: %v, cast to rune has value: %v, cast to run has value: %v\n",
		a, a, rune(a), string(a))
	fmt.Println("Hello cast to bytes is:", []byte("Hello"))
	fmt.Println("Hello as a list of runes is: [", 'H', 'e', 'l', 'l', 'o', "]")
	fmt.Println("[ 72 101 108 108 111 ] cast to string is", runesToString([]rune{72, 101, 108, 108, 111}))

	for i := 5000; i <= 5005; i++ {
		fmt.Printf("%v is the string %v with byte value %v\n", i, string(i), []byte(string(i)))
	}
}

// this function uses a named return (not amazing) and a slice argument
func runesToString(runes []rune) (outString string) {
	// don't need index so _
	for _, v := range runes {
		outString += string(v)
	}
	return
}
