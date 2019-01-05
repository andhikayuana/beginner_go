package main

import "fmt"

func main() {
	a, b := swap("Jarjit", "Singh")
	fmt.Println(a, b)
}

func swap(x, y string) (string, string) {
	return y, x
}
