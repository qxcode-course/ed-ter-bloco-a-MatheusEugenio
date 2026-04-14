package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}

	var min, max int = nums[0], nums[0]
	var soma int
	// TODO: Encontre min/max e some o resto
	// Loop para min/max
	// Outro loop ou soma total - min - max
	fmt.Println("COMPLETE AQUI")
}

