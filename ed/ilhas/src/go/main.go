package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func marcaIlhas(grid [][]byte, l, c int, visitados map[Pos]bool) bool {

	if l < 0 || c < 0 || l >= len(grid) || c >= len(grid[0]) {
		return false
	}

	posAtual := Pos{l, c}
	if visitados[posAtual] || grid[l][c] != '1' {
		return false
	}

	visitados[posAtual] = true

	if marcaIlhas(grid, l, c+1, visitados) ||
		marcaIlhas(grid, l, c-1, visitados) ||
		marcaIlhas(grid, l+1, c, visitados) ||
		marcaIlhas(grid, l-1, c, visitados) {
		return true
	}

	return false
}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {

	visitados := make(map[Pos]bool)
	ilhas := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			pos := Pos{i, j}
			_, exist := visitados[pos]

			if !exist && grid[i][j] == '1' {
				ilhas++
				marcaIlhas(grid, i, j, visitados)
			}

		}
	}

	return ilhas
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
