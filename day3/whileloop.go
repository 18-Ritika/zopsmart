package main

import "fmt"

func tilln(n int) int {
	i := 1
	varr := 0

	for i <= n {
		varr += i
		i += 1
	}
	return varr
}

func main() {

	fmt.Println(tilln(10))

}
