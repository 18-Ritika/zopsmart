package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Println("number 3")
			break
		}
		fmt.Println(i)
	}
}
