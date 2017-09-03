package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(`
The number, 197, is called a circular prime because all rotations of the digits:
   197, 971, and 719,
are themselves prime.

There are thirteen such primes below 100:
2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
`)
	knownPrimes := primesBelowOneMillion()
	cPrimes := circularPrimes(&knownPrimes)
	sort.Ints(cPrimes)
	fmt.Println("Found these circular primes")
	fmt.Println(cPrimes)
	fmt.Println()
	fmt.Println("There are", len(cPrimes), "below one million.")
}

func primesBelowOneMillion() []int {
	primes := []int{2, 3, 5, 7, 11, 15}
	for i := 15; i <= 1000000; i++ {
		if isPrime(&i, &primes) {
			primes = append(primes, i)
		}
	}
	// fmt.Println("list of primes", primes)
	return primes
}

func isPrime(i *int, knownPrimes *[]int) bool {
	indivisible := true
	for _, v := range *knownPrimes {
		if *i%v == 0 {
			indivisible = false
			break
		}
	}
	return indivisible
}

func contains(slice *[]int, i *int) bool {
	for _, v := range *slice {
		if v == *i {
			return true
		}
	}
	return false
}

func circularPrimes(knownPrimes *[]int) []int {
	cPrimes := []int{2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, 97}
	var containsEven bool
	for _, prime := range *knownPrimes {
		if contains(&cPrimes, &prime) {
			continue
		}
		rotations := []int{}
		current := prime
	AddToCPrimes:
		for {
			current, containsEven = rotateInt(&current)
			switch {
			case containsEven:
				break AddToCPrimes
			case !contains(knownPrimes, &current):
				// fmt.Println("SHOULD BREAK OUT OF LOOP,", current, " NOT PRIME")
				break AddToCPrimes
			case (current == prime):
				rotations = append(rotations, prime)
				cPrimes = append(cPrimes, rotations...)
				//fmt.Println("SHOULD BREAK OUT OF LOOP,", rotations, " ADDED.")
				break AddToCPrimes
			}
			rotations = append(rotations, current)
		}
	}
	return cPrimes
}

func rotateInt(i *int) (int, bool) {
	if *i < 10 {
		return *i, false
	}
	slice := strings.Split(strconv.Itoa(*i), "")
	rotatedString := slice[len(slice)-1]
	rest := slice[0 : len(slice)-1]
	for _, char := range rest {
		rotatedString += char
	}
	j, _ := strconv.Atoi(rotatedString)

	return j, hasEven(&slice)
}

func hasEven(ints *[]string) bool {
	for _, v := range *ints {
		x, _ := strconv.Atoi(v)
		if x%2 == 0 {
			return true
		}
	}
	return false
}
