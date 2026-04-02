package main

import "fmt"

type ParQR struct {
	quociente, resto int
}

func imprimir_invertido(pares []ParQR) { // func recursiva

	if len(pares) == 0 {
		return
	}

	pos := len(pares) - 1

	quo := pares[pos].quociente
	rest := pares[pos].resto
	fmt.Printf("%d %d\n", quo, rest)

	imprimir_invertido(pares[:len(pares)-1])
	// atualiza o tamanho do slice (cria um sub-slice), diminuindo um elemento da posição "len(pares)-1", até chegar no caso base
}

func formaParQR(num int) []ParQR {

	paresSlice := make([]ParQR, 0)
	resto := 0

	for {
		if num == 0 && resto == 1 {
			break
		}

		num, resto = num/2, num%2

		paresSlice = append(paresSlice, ParQR{num, resto})
	}

	return paresSlice
}

func main() {

	var num int
	fmt.Scan(&num)

	paresRQ := formaParQR(num) //map[quociente]resto

	imprimir_invertido(paresRQ)

}
