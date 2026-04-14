package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	animais := make([]int, n)
	for i := range animais {
		fmt.Scan(&animais[i])
	}

	par := make(map[int]int)
	qtdPares := 0

	for _, val := range animais {
		if par[-val] > 0 {
			par[-val]--
			qtdPares++
		} else {
			par[val]++
		}
	}

	fmt.Println(qtdPares)
}
