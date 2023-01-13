package main

import "fmt"

func main() {
	var arr [3]string
	arr[0] = "Ritika"
	arr[1] = "jaiswal"

	fmt.Println(arr)

	arr2 := [2]int{1, 2}
	fmt.Println(arr2)

	for i, value := range arr {
		fmt.Println(i, value)
	}

	arr[2] = "Zopsmart"

	fmt.Println(arr)

}
