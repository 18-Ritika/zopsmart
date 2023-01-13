package main

import "fmt"

func main() {

	slic := []int{1, 2, 3, 4, 5, 6}

	for i, j := 0, len(slic)-1; i < j; i, j = i+1, j-1 {
		slic[i], slic[j] = slic[j], slic[i]
	}
	fmt.Println(slic)
}
