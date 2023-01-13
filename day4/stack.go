package main

import "fmt"

type shape struct {
	x int
	y int
}

func main() {
	/*
			i, j := 2, 3

			p := &i
			fmt.Println(*p)
			q := &j
			fmt.Println(*q)

			*p = 4
			fmt.Println(i)
			fmt.Println(*p)
			fmt.Println(p)
		square := shape{}
		fmt.Println(square.x, square.y)

		square.x = 2
		square.y = 2
		fmt.Println(square.x)
	*/

	var (
		square    = shape{2, 2}
		recrangle = shape{4, 6}

		p = &shape{3, 7}
	)

	fmt.Println(square.x, square.y)
	fmt.Println(recrangle.x, recrangle.y)
	fmt.Println(p)
	fmt.Println(square.x, square.y)

}
