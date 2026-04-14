package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	animais := make([]int, n)
	for i := range animais {
		fmt.Scan(&animais[i])
	}

	par := make(map[int]bool)
	qtdPares := 0

	for _, val := range animais {
		_, existe := par[-val]
		if existe {
			par[-val] = false
			qtdPares++
		} else {
			par[val] = true
		}
	}

	fmt.Println(qtdPares)
}
