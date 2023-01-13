package main

import "fmt"

func main() {

	var pow = []int{1, 2, 3, 4, 5}
	for i, value := range pow {
		fmt.Println(i, value)
	}
}
