package main

import (
	"bufio"
	"fmt"
	"os"
)

func buscaDeVizinhos(board [][]byte, l, c int) {

	if l < 0 || l >= len(board) || c < 0 || c > len(board[0]) || board[l][c] != 'O' {
		return
	}

	board[l][c] = 'X'

	buscaDeVizinhos(board, l, c+1)
	buscaDeVizinhos(board, l, c-1)
	buscaDeVizinhos(board, l-1, c)
	buscaDeVizinhos(board, l+1, c)
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {

	numL := len(board)
	numC := len(board[0])

	for i := 0; i < numL; i++ {
		for j := 0; j < numC; j++ {
			if board[i][j] == 'O' {
				buscaDeVizinhos(board, i, j)
				return
			}
		}
	}

}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
