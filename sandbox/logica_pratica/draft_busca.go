package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}

	var count int
	// TODO: Conte ocorrências de k no array
	fmt.Println(count)
}

