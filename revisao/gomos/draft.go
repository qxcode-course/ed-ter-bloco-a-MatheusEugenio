package main

import "fmt"

type Par struct {
	x, y int
}

func main() {

	var direcao string //* 'L' = ESQUERDA ; 'R' = DIREITA ; 'U' = PRA CIMA ; 'D' = PRA BAIXO
	var quantidadeDeGomos int
	fmt.Scan(&quantidadeDeGomos, &direcao)

	gomos := make([]Par, quantidadeDeGomos)

	for i := 0; i < quantidadeDeGomos; i++ {
		fmt.Scan(&gomos[i].x, &gomos[i].y)
	}

	//* A cauda se mexe ate chegar na cabeça
	for i := quantidadeDeGomos - 1; i > 0; i-- {
		gomos[i] = gomos[i-1]
	}

	//* movimento da cabeça
	switch direcao {
	case "U":
		gomos[0].y--
	case "D":
		gomos[0].y++
	case "R":
		gomos[0].x++
	case "L":
		gomos[0].x--
	}

	for i := 0; i < quantidadeDeGomos; i++ {
		fmt.Printf("%d %d\n", gomos[i].x, gomos[i].y)
	}

}
