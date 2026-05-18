package main

import "fmt"

func ehPrimoAux(n int, divisor int) bool {

	if divisor == 1 {
		return true
	}

	if n%divisor == 0 {
		return false
	}
	return ehPrimoAux(n, divisor-1)
}

func ehPrimo(n int) bool {

	if n <= 1 {
		return false
	}

	return ehPrimoAux(n, n-1)
}

func encontrarPrimo(n int, atual int) int {

	if ehPrimo(atual) {
		if n == 1 {
			return atual
		}
		return encontrarPrimo(n-1, atual+1)
	}

	return encontrarPrimo(n, atual+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(encontrarPrimo(n, 2))
}
