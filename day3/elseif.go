package main

import "fmt"

func greaterthan(x, y, z int) int {

	if x > y && x > z {
		return x
	} else if y > x && y > z {
		return y
	} else {
		return z
	}

}

func main() {
	fmt.Println(greaterthan(2, 3, 4))
}
