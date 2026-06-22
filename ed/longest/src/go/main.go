package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	l, c int
}

func buscaCaminho(matrix [][]int, i, j int) int {

	melhorCaminho := 1

	direcoes := [][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	for _, dir := range direcoes {

		vizI := i + dir[0]
		vizJ := j + dir[1]

		if vizI < 0 || vizJ < 0 || vizI >= len(matrix) || vizJ >= len(matrix[0]) {
			continue
		}

		vizinho := matrix[vizI][vizJ]
		atual := matrix[i][j]

		if vizinho > atual {

			caminhoVizinho := buscaCaminho(matrix, vizI, vizJ)

			if caminhoVizinho+1 > melhorCaminho {
				melhorCaminho = caminhoVizinho + 1
			}
		}
	}

	return melhorCaminho
}

func longestIncreasingPath(matrix [][]int) int {

	num_l := len(matrix)
	num_c := len(matrix[0])

	melhor := 0
	for i := 0; i < num_l; i++ {
		for j := 0; j < num_c; j++ {

			resultado := buscaCaminho(matrix, i, j)

			melhor = max(melhor, resultado)
		}
	}

	return melhor
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
