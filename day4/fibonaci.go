package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() {
	f1 := 0
	f2 := 1
	var f3 int
	return func() {
		fmt.Println(f1)
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
}

func main() {
	f := fibonacci()
	for i := 0; i <= 5; i++ {
		f()
	}
}
