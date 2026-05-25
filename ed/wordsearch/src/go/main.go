package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func dfs(grid [][]byte, word string, l, c int, visited map[Pos]bool, index int) bool {

	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]) {
		return false
	}

	posAtual := Pos{l, c}

	if visited[posAtual] {
		return false
	}

	if grid[l][c] != word[index] {
		return false
	}

	if index == len(word)-1 {
		return true
	}

	visited[posAtual] = true

	if dfs(grid, word, l, c+1, visited, index+1) ||
		dfs(grid, word, l, c-1, visited, index+1) ||
		dfs(grid, word, l-1, c, visited, index+1) ||
		dfs(grid, word, l+1, c, visited, index+1) {
		return true
	}

	delete(visited, posAtual)

	return false
}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {

	num_linhas := len(grid)
	num_colums := len(grid[0])

	for i := 0; i < num_linhas; i++ {
		for j := 0; j < num_colums; j++ {

			visited := make(map[Pos]bool)

			if dfs(grid, word, i, j, visited, 0) {
				return true
			}

		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
