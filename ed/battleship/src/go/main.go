package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func preencher(board [][]byte, i int, j int, visitados map[Pos]bool) bool {

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}

	pos := Pos{i, j}
	if visitados[pos] || board[i][j] == '.' {
		return false
	}

	visitados[pos] = true

	if preencher(board, i, j+1, visitados) ||
		preencher(board, i, j-1, visitados) ||
		preencher(board, i+1, j, visitados) ||
		preencher(board, i-1, j, visitados) {
		return true
	}

	return false
}

// Função que será chamada no LeetCode
func countBattleships(board [][]byte) int {

	num_l := len(board)
	num_c := len(board[0])
	visitados := make(map[Pos]bool)
	navios := 0

	for i := 0; i < num_l; i++ {
		for j := 0; j < num_c; j++ {

			pos := Pos{i, j}
			_, exist := visitados[pos]

			if !exist && board[i][j] == 'X' {
				navios++
				preencher(board, i, j, visitados)
			}
		}
	}

	return navios
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	board := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}

	result := countBattleships(board)
	fmt.Println(result)
}
