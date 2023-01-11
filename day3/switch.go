package main

import "fmt"

func main() {
	office := "sector1"
	switch office {

	case "sector1":
		fmt.Println("Zopsmart old")

	case "sector2":
		fmt.Println("Zopsmart new")

	case "noida":
		fmt.Println("Zopsmart Noida")

		//default:
		//fmt.Println("Zopsmart")

	}
}
