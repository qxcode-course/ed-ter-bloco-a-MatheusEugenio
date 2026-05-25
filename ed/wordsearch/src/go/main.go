package main

import (
	"bufio"
	"fmt"
	"os"
	"structs"
)

type Pos struct{
	l, c int
}

func exist_letra(grid[][]byte, word string, l, c int) bool{
	

}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string, visited map[Pos]bool) bool {

	if len(visited) == word{
		return true
	}

	num_linhas := len(grid) - 1
	num_colums := len(grid[0])

	if num_linhas < 0 || num_colums < 0 || num_linhas >= len(grid) || num_colums >= len(grid[0]){
		return false
	}
	
	for i := 0; i < num_linhas; i++{
		for j := 0; j < num_colums; j++ {
			
			pos := Pos(i,j)
			_, exist := visited[pos]

			if exist && {

			}
			
			if grid[num_linhas][num_colums] == visited[]{
				
			}
		}
	}

	if grid[num_linhas][num_colums] != "esperado"{

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

	visited := make(map[rune]bool)
	if exist(grid, word, visited) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
