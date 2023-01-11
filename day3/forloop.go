package main

func total(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}
func main() {
	summ := total(4)
	println(summ)
}
