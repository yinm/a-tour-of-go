// 2nd
package main

import "fmt"

func fibonacci() func() int {
	a1, a2 := 0, 1
	return func() int {
		ret := a1
		a1, a2 = a2, a1 + a2
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}