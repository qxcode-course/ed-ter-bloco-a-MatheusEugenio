package main

import "fmt"

func fatorial(n int) int {
	// TODO: Implemente recursão
	// Base: n <= 1 -> 1
	// Recursivo: n * fatorial(n-1)
	return 0 // Stub
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fatorial(n))
}

