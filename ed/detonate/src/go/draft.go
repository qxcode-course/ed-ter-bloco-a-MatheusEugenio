package main

import (
	"fmt"
)

func procuraBombas(bombaAtual int, adj [][]int, visitados map[int]bool) int {

	visitados[bombaAtual] = true
	cont := 1

	for _, bombadj := range adj[bombaAtual] {

		if !visitados[bombadj] {
			cont += procuraBombas(bombadj, adj, visitados)
		}
	}

	return cont
}

func distancia(x1, y1, x2, y2 int) float64 {

	x := float64(x1) - float64(x2)
	y := float64(y1) - float64(y2)
	return x*x + y*y
}

func raio(r int) float64 {
	return float64(r) * float64(r)
}

func estouros(bombs [][]int) int {

	n := len(bombs)
	adj := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			if i == j {
				continue
			}

			x1, y1, r1 := bombs[i][0], bombs[i][1], bombs[i][2]
			x2, y2 := bombs[j][0], bombs[j][1]

			if distancia(x1, y1, x2, y2) <= raio(r1) {
				adj[i] = append(adj[i], j)
			}
		}
	}

	explosoes := 0

	for i := 0; i < n; i++ {

		visitados := make(map[int]bool)
		resu := procuraBombas(i, adj, visitados)

		if resu > explosoes {
			explosoes = resu
		}
	}

	return explosoes
}

func main() {

	var n, m int
	fmt.Scan(&n, &m)

	bombs := make([][]int, n)

	for i := 0; i < n; i++ {
		var x, y, r int
		fmt.Scan(&x, &y, &r)

		bombs[i] = []int{x, y, r}
	}

	bang := estouros(bombs)
	fmt.Println(bang)
}
