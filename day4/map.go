package main

import "fmt"

func main() {
	/*
		mpp := map[int]int{0: 1, 1: 2, 2: 3}

		for key, value := range mpp {
			fmt.Println(key, mpp[key], value)

		}
	*/
	//var mpp2 map[int] string
	mpp := make(map[int]int)
	mpp[0] = 1
	mpp[1] = 2
	fmt.Println(mpp)

	delete(mpp, 0)
	fmt.Println(mpp)

	key, value := mpp[1]
	fmt.Println(key, value)

}
