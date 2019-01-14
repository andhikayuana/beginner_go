package main

import "fmt"

func finbonanci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return finbonanci(i-1) + finbonanci(i-2)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", finbonanci(i))
	}
}
