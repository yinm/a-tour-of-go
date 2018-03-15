package main

import "fmt"

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)

	v2 := 3.14
	fmt.Printf("v2 is ot type %T\n", v2)

	v3 := 0.867 + 0.5i
	fmt.Printf("v3 is of type %T\n", v3)
}