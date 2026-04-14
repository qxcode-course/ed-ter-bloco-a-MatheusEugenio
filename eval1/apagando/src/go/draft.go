package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	pessoas := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&pessoas[i])
	}

	var qtdDeixaramFila int
	fmt.Scan(&qtdDeixaramFila)

	posicoesDeixadas := make(map[int]bool)

	for i := 0; i < qtdDeixaramFila; i++ {
		var saiu int
		fmt.Scan(&saiu)
		posicoesDeixadas[saiu] = true
	}

	novoArrayPessoas := make([]int, 0)
	for _, pessoa := range pessoas {
		if !posicoesDeixadas[pessoa] {
			novoArrayPessoas = append(novoArrayPessoas, pessoa)
		}
	}

	for _, val := range novoArrayPessoas {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

}
