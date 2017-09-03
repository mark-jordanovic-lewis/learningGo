package main

import "fmt"

func main() {
	// simple remainder function (modulo)
	x := 1337 % 13
	fmt.Println("remainder 1337/13 =", x)

	for i := 0; i <= 20; i++ {
		if i%2 == 1 {
			fmt.Println(i, "is Odd.")
		} else {
			fmt.Println(i, "is Even.")
		}
	}
}
