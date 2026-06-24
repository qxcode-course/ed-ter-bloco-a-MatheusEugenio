package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return nil
}

// está na posição adequada
func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

// verifica se é o caractere referido
func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

// method principal
func search(grid [][]rune, startPos Pos, endPos Pos, visitados map[Pos]bool, caminho map[Pos]Pos, fila Queue[Pos]) bool {

	if !inside(grid, startPos) {
		return false
	}

	if match(grid, startPos, '#') || visitados[startPos] {
		return false
	}

	if match(grid, startPos, 'F') && startPos == endPos {
		return true
	}

	visitados[startPos] = true
	
	fila.Enqueue(startPos)
	caminho[startPos] =  
	// registrar o nó atual como anterior do vizinho no mapa caminho
	grid[startPos.l][startPos.c] = '.'

	if search(grid, startPos, endPos, visitados, caminho) ||
		search(grid, startPos, endPos, visitados, caminho) ||
		search(grid, startPos, endPos, visitados, caminho) ||
		search(grid, startPos, endPos, visitados, caminho) {
		return true
	}

	visitados[startPos] = false
	grid[startPos.l][startPos.c] = ' '

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	queue := *NewQueue[Pos]()
	visitados := make(map[Pos]bool)
	caminho := make(map[Pos]Pos)

	search(mat, inicio, fim, visitados, caminho, queue)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
