package main

import "fmt"

// Recursion works as you would expect.

func main() {
	// call func with two same args
	twoArgsSameType("Beep", "Boop")
	// string returned from stringPrinter
	concatedString := stringPrinter("Blerp", "Blorp")
	fmt.Println(concatedString)
	// named return function call
	concatedString = namedReturnS("Broop", "Braap")
	fmt.Println(concatedString)
	// multiple returns
	concatedString1, concatedString2 := multipleReturns("First", "Last")
	fmt.Println(concatedString1, concatedString2)
	// variadic parameter call
	fmt.Println("mean of 1,2,3,4,5,6,7,8 is", mean(1, 2, 3, 4, 5, 6, 7, 8, 9))
	// make a slice and unwrap into a variatic parameter
	slice := []float64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println("mean of 10,11,12,13,14,15,16,17,18,19 is", mean(slice...))
	// anonymous function, can be passed. Obeys scope rules
	anon := func(i int) string {
		return fmt.Sprint("I am anon#", i)
	}
	fmt.Println(anon(0))
	// using anon func as callback to function
	callbackPrint([]int{432, 654, 13, 654, 6, 11, 101, 32, 47, 18, 200}, anon)
	// filtering out low level anons and printing
	fmt.Println("only experienced anons allowed!")
	callbackPrint(
		filterIntSlice(
			[]int{432, 654, 13, 654, 6, 11, 101, 32, 47, 18, 200},
			func(x int) bool {
				return x > 100
			}),
		anon)
	// calling simple function with very simple defer
	helloWorld()
	// anonymous self executing function
	func() {
		fmt.Println("Main func ends.")
	}()
}

func twoArgsSameType(str1, str2 string) {
	fmt.Println(str1, str2)
}

func stringPrinter(str1, str2 string) string {
	// Sprint prints TO a string variable (like concatination)
	return fmt.Sprint(str1, " ", str2)
}

func namedReturnS(str1, str2 string) (s string) {
	// s is assigned as the return value so is set as the actual return when the function ends
	s = fmt.Sprint(str1, " ", str2)
	x := 1
	x++
	return
}

func multipleReturns(str1, str2 string) (string, string) {
	return fmt.Sprint(str1, " ", str2), fmt.Sprint(str2, " ", str1)
}

// arg type prefixed with ... implies unlimited number af args
func mean(fs ...float64) float64 {
	total := 0.0
	// the underscore is ignoring the index of the slice (a slice is kindof an array-ish)
	for _, v := range fs {
		total += v
	}
	return total / float64(len(fs))
}

// callbacks are just passing functions as arguments, this callback takes and int and returns a string
func callbackPrint(ints []int, callback func(int) string) {
	for i, v := range ints {
		fmt.Println(i, callback(v))
	}
}

func filterIntSlice(xs []int, callback func(x int) bool) (filteredxs []int) {
	for _, v := range xs {
		if callback(v) {
			filteredxs = append(filteredxs, v)
		}
	}
	return
}

func helloWorld() {
	defer world()
	hello()
}

func hello() {
	fmt.Print("Hello, ")
}

func world() {
	fmt.Println("world!")
}
