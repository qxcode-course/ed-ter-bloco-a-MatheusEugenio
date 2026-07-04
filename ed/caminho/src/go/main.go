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
	return []Pos{{p.l, p.c + 1},
		{p.l, p.c - 1},
		{p.l + 1, p.c},
		{p.l - 1, p.c},
	}
}

// está na posição adequada
func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

// verifica se É o caractere referido
func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

// method principal
func search(grid [][]rune, startPos Pos, endPos Pos, visitados map[Pos]bool, caminho map[Pos]Pos, queue Queue[Pos]) bool {

	visitados[startPos] = true

	queue.Enqueue(startPos)

	for !queue.IsEmpty() {

		pos_front, _ := queue.Dequeue()

		if pos_front == endPos {
			return true
		}

		for _, vizinho := range pos_front.getNeig() {

			if inside(grid, vizinho) && !match(grid, vizinho, '#') && !visitados[vizinho] {

				visitados[vizinho] = true
				queue.Enqueue(vizinho)
				caminho[vizinho] = pos_front
			}
		}
	}

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

	p := fim
	for p != inicio {
		
		mat[p.l][p.c] = '.'
		p = caminho[p]
	}
	mat[p.l][p.c] = '.'

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
