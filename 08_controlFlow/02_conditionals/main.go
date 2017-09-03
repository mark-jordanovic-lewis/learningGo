package main

import "fmt"

type SwitchType struct {
	str string
	i   int
	f   float64
}

func switchOnType(x interface{}) {
	switch x.(type) { // type is a keyword .(type) introspects variable type
	case int:
		fmt.Println("passed an int:", x)
	case float64:
		fmt.Println("passed a float64:", x)
	case string:
		fmt.Println("passed a string:", x)
	case SwitchType:
		fmt.Println("passed our default SwitchType:", x)
	default:
		fmt.Println("passed an unknown:", x)
	}
}

func main() {
	thisString := "string_a"
	fmt.Println("std switch statement")
	switch thisString {
	case "string_a":
		fmt.Println("string_a:", "matched")
	case "string_A":
		fmt.Println("string_A:", "no match")
	case "STRING_A":
		fmt.Println("STRING_A:", "no match")
	}
	fmt.Println()
	fmt.Println("switch statement with fall through")
	switch thisString {
	case "string_a":
		fmt.Println("case string_a:", "correct case and...")
		fallthrough
	case "string_A":
		fmt.Println("case string_A:", "fallen through to and..")
		fallthrough
	case "STRING_A":
		fmt.Println("STRING_A:", "fallthrough ends here.")
	case "sTrInG_a":
		fmt.Println(thisString, "not seen.")
	}
	fmt.Println()
	fmt.Println("switch on multiple")
	switch thisString {
	case "sTrInG_a", "STRING_A":
		fmt.Println("STRING_A:", "not matched")
	case "string_A", "string_a":
		fmt.Println("string_a", "matched on string_A || string_a")
	}
	fmt.Println("switch with no condition")
	switch {
	case false:
		fmt.Println("false so never seen")
		fmt.Println()
	case true:
		fmt.Println("true so is seen")
	}
	fmt.Println()
	fmt.Println("switch with no condition and default with closure var")
	switch defaultString := "hi there i am default"; {
	case !true:
		fmt.Println("never seen")
	case false:
		fmt.Println("unseen")
	default:
		fmt.Println(defaultString)
	}
	fmt.Println()
	fmt.Println("switch based on type")
	switchOnType(42)
	switchOnType(3.14)
	switchOnType("beep")
	switchOnType(true)
	switchOnType(SwitchType{"hey", 42, 3.14})
}
