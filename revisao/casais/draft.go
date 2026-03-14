package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Scan(&n)

	animais := make([]int, n)
	var quantDePares int

	// semPares := make([]int, n)

	macho := 0
	femea := 0

	for i := 0; i < n; i++ {

		fmt.Scan(&animais[i])

		if animais[i] > 0 {
			macho = animais[i]
		} else {
			femea = animais[i]
		}

		if macho != 0 && femea != 0 {
			quantDePares++
			femea = 0
		}

	}

	fmt.Printf("%d\n", quantDePares)

}
