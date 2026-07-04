package main

import (
	"fmt"
)

type Pos struct {
	l, c int
}

func contagemMinutos(grid [][]int, podres []Pos, quantFrescas int, minutos int) (int, int) {

	if quantFrescas == 0 {
		return 0, 0
	}

	direcoes := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(podres) > 0 && quantFrescas > 0 {

		minutos++

		for _, podreAtual := range podres {
			podres = podres[1:]

			for _, vizinhos := range direcoes {

				nl := podreAtual.l + vizinhos.l
				nc := podreAtual.c + vizinhos.c

				if nl < len(grid) && nl >= 0 && nc < len(grid[0]) && nc >= 0 && grid[nl][nc] == 1 {

					grid[nl][nc] = 2
					podres = append(podres, Pos{nl, nc})
					quantFrescas--
				}
			}
		}

	}

	return quantFrescas, minutos
}

func apodrecedor(grid [][]int) int {

	qtdfrescas := 0
	var filaPodres = []Pos{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == 1 {
				qtdfrescas++
			}

			if grid[i][j] == 2 {
				filaPodres = append(filaPodres, Pos{i, j})
			}
		}
	}

	minutos := 0
	qtdfrescas, minutos = contagemMinutos(grid, filaPodres, qtdfrescas, minutos)

	if qtdfrescas > 0 {
		return -1
	}

	return minutos
}

func main() {

	var l, c int
	fmt.Scan(&l, &c)

	mat := make([][]int, l)

	for i := 0; i < l; i++ {
		mat[i] = make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Scan(&mat[i][j])
		}
	}

	fmt.Println(apodrecedor(mat))

}
