package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

// acessar vizinho
func getNeig(p Pos) []Pos {
	return []Pos{{p.l, p.c - 1}, {p.l - 1, p.c}, {p.l, p.c + 1}, {p.l + 1, p.c}}
}

// se a POS está dentro da matriz ou não
func inside(grid [][]rune, p Pos) bool {
	return !(p.l < 0 || p.l >= len(grid) || p.c < 0 || p.c >= len(grid[0]))
}

// se corresponde ao caractere passado como parâmetro
func match(grid [][]rune, p Pos, value rune) bool {
	return inside(grid, p) && grid[p.l][p.c] == value
}

// Função recursiva que tenta encontrar o caminho do início ao fim
func search(grid [][]rune, startPos, endPos Pos, visitados map[Pos]bool) bool {

	// verfica se são POS válidas
	if !inside(grid, startPos) {
		return false
	}

	if grid[startPos.l][startPos.c] != ' ' || visitados[startPos] {
		return false
	}

	// se a POS atual/de começo é igual a de destino
	if startPos == endPos {
		grid[endPos.l][endPos.c] = '.'
		visitados[endPos] = true
		return true
	}

	visitados[startPos] = true
	grid[startPos.l][startPos.c] = '.'

	vizinhos := getNeig(startPos)
	for _, vizinhoActual := range vizinhos {

		if search(grid, vizinhoActual, endPos, visitados) {
			return true
		}
	}

	grid[startPos.l][startPos.c] = ' '

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()

	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)

	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	visitados := make(map[Pos]bool)
	search(grid, startPos, endPos, visitados)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
