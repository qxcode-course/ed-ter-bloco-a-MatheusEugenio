package main

import (
	"fmt"
	"strings"
)

func main() {

	var n, quantDeixaramFila int
	fmt.Scan(&n)

	idPessoasDafila := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&idPessoasDafila[i])
	}

	fmt.Scan(&quantDeixaramFila)

	idPessoasDeixaramFila := make(map[int]int) // K: [int]; V: int

	for i := 0; i < quantDeixaramFila; i++ {
		var id int
		fmt.Scan(&id)
		idPessoasDeixaramFila[id] = 1
	}

	pessoasQueFicaramNaFila := make([]int, 0)

	for _, id := range idPessoasDafila {
		_, existe := idPessoasDeixaramFila[id]

		if !existe {
			pessoasQueFicaramNaFila = append(pessoasQueFicaramNaFila, id)
		}
	}

	fmt.Print(strings.Trim(fmt.Sprint(pessoasQueFicaramNaFila), "[]"), " \n")

}
