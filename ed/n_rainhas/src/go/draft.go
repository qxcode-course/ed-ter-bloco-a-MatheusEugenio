package main

import "fmt"

type Pos struct {
	i, j int
}

func posicionar(tabuleiro [][]int, rainhas int) {

	posicoesDiagonaisPrim := make([]Pos, 0)
	posicoesDiagonaisSec := make([]Pos, 0)

	n := len(tabuleiro[0])
	for i := 0; i < n; i++ {
		posicoesDiagonaisPrim = append(posicoesDiagonaisPrim, Pos{i, i})
		posicoesDiagonaisSec = append(posicoesDiagonaisSec, Pos{i, n - 1 - i})
	}

}

func main() {

	var rainhas int
	fmt.Scan(&rainhas)

	tabuleiro := make([][]int, rainhas)

	posicionar(tabuleiro, rainhas)

}
