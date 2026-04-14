package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	valoresPessoas := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&valoresPessoas[i])
	}

	var qtdDeixaramFila int
	fmt.Scan(&qtdDeixaramFila)

	posicoesDeixadas := make([]int, qtdDeixaramFila)

	for i := 0; i < qtdDeixaramFila; i++ {
		fmt.Scan(&posicoesDeixadas[i])
	}

	novoArrayPessoas := make([]int, 0)
	for i := 0; i < len(valoresPessoas); i++ {
		if valoresPessoas[i] == posicoesDeixadas[i] {
			continue
		} else {
			novoArrayPessoas = append(novoArrayPessoas, valoresPessoas[i])
		}
	}

	for i, val := range novoArrayPessoas {
		if i == len(novoArrayPessoas)-1 {
			fmt.Printf("%d", val)
			fmt.Println()
		} else {
			fmt.Printf("%d ", val)
		}
	}

}
