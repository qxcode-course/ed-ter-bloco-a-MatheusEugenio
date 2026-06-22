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

func pintar(image [][]int, i, j int, color int, visitados map[Pos]bool, colorAtual int) bool {

	if i < 0 || j < 0 || i >= len(image) || j >= len(image[0]) {
		return false
	}

	pos := Pos{i, j}
	if visitados[pos] || image[i][j] != colorAtual {
		return false
	}

	visitados[pos] = true
	image[i][j] = color

	if pintar(image, i, j+1, color, visitados, colorAtual) ||
		pintar(image, i, j-1, color, visitados, colorAtual) ||
		pintar(image, i-1, j, color, visitados, colorAtual) ||
		pintar(image, i+1, j, color, visitados, colorAtual) {
		return true
	}

	return false
}

// Não modifique a assinatura da função floodFill
func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	pos := Pos{sr, sc}
	corAtual := image[pos.l][pos.c]
	visitados := make(map[Pos]bool)

	pintar(image, pos.l, pos.c, color, visitados, corAtual)
	return image
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	image := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		rowStr := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc; j++ {
			row[j], _ = strconv.Atoi(rowStr[j])
		}
		image[i] = row
	}

	// Lê sr, sc e color
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	sr, _ := strconv.Atoi(params[0])
	sc, _ := strconv.Atoi(params[1])
	color, _ := strconv.Atoi(params[2])

	result := floodFill(image, sr, sc, color)

	// Imprime a matriz resultante
	for _, row := range result {
		for j, val := range row {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}
