package main

import "fmt"

func main() {
	myslice := make([]int, 2)
	myslice[0] = 1
	myslice[1] = 2
	myslice = append(myslice, 2)
	fmt.Println(myslice)

	arr1 := []int{1, 2, 3, 4}
	var slicee []int = arr1[0:2]
	fmt.Println(slicee)

	slicee[0] = 9
	fmt.Println(arr1)

	var sli []int
	fmt.Println(sli)
	sli = append(sli, 2)
	fmt.Println(sli)

}
