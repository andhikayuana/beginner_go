package main

import "fmt"

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Println("type of x :", i)
	case int:
		fmt.Println("x is int")
	case float64:
		fmt.Println("x is float64")
	case func(int) float64:
		fmt.Println("x is func(int)")
	case bool, string:
		fmt.Println("x is bool or string")
	default:
		fmt.Println("dntknow type!")
	}
}
