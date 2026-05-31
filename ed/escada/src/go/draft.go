package main

import "fmt"

func contagemDePassos(n int) int {

	if n <= 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	return contagemDePassos(n-1) + contagemDePassos(n-3)
}

func main() {

	var n int
	fmt.Scan(&n)

	fmt.Println(contagemDePassos(n))
}
