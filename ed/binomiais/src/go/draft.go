package main

import "fmt"

func calcCoeficientes(n, k int) int {

	if k == 0 {
		return 1
	}

	if k == n {
		return 1
	}

	return calcCoeficientes(n-1, k-1) + calcCoeficientes(n-1, k)
}
func main() {

	var n, k int
	fmt.Scan(&n, &k)

	fmt.Println(calcCoeficientes(n, k))
}
