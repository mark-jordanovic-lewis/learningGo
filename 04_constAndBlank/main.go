package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// this const has package scope, not visible externally
const p string = "otters are lovely"

// can do this
const s = 12
const (
	// Pi lovely Pi
	Pi = 3.141592
	// Language What lang we use?
	Language = "Go"
	chicken  = "bok bok"
)

// con't do this
// const t int

// P const has package scope, visible externally
const P string = "otters are super lovely"

// iota is an incremental concept
const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

// can be declared
const (
	D = iota // 0
	E        // 1
	F        // 2
	G        // 3
	H        // 4
)

func printIotas() {
	fmt.Println("Iota Block 1")
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println("Iota Block 2")
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	equal := A == D
	fmt.Printf("Block 1 A == Block 2 D ? %t\n", equal)
}

// bit shifting with Iota
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10) ie 2^10
	MB = 1 << (iota * 10) // 1 << (2 * 10) ie 2^20
	GB = 1 << (iota * 10) // 1 << (3 * 10) ie 2^30
	TB = 1 << (iota * 10) // 1 << (4 * 10) ie 2^40
)

func printBitshifts() {
	fmt.Println("\tDec\t\tBin")
	fmt.Printf("%d\t\t%b\n", KB, KB)
	fmt.Printf("%d\t\t%b\n", MB, MB)
	fmt.Printf("%d\t%b\n", GB, GB)
	fmt.Printf("%d\t%b\n", TB, TB)
}

func main() {
	// BLANK Variable
	// code with err checking
	res1, err := http.Get("http://www.twitter.com/")
	if err != nil {
		log.Fatal(err)
	}
	page1, err := ioutil.ReadAll(res1.Body)
	res1.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", page1[0])

	// code w.out err checking using blank
	res2, _ := http.Get("http://www.google.com/")
	page2, _ := ioutil.ReadAll(res2.Body)
	res2.Body.Close()
	fmt.Printf("%v\n", page2[0])

	fmt.Println()
	fmt.Println()

	// CONSTANTS

	const q = 42

	fmt.Println("p = ", p)
	fmt.Println("q = ", q)

	// Iotas
	fmt.Println()
	fmt.Println("Iotas")
	printIotas()
	fmt.Println()
	fmt.Println("Printing bit shifted Iotas")
	printBitshifts()
}
