package main

import (
	"fmt"
	"math"
)

func main() {

	var nCompetidores int

	fmt.Scan(&nCompetidores)

	var distancia_a, distancia_b, difDistancia int

	indiceMenorDistancia := -1

	menorDiferenca := 99999999999999

	for i := 0; i < nCompetidores; i++ {
		fmt.Scan(&distancia_a, &distancia_b)

		if distancia_a < 10 || distancia_b < 10 {
			continue
		}

		if (distancia_a > 100 || distancia_a < 1) || (distancia_b > 100 || distancia_b < 1) {
			continue
		}

		difDistancia = int(math.Abs(float64(distancia_a - distancia_b)))

		if difDistancia < menorDiferenca {
			menorDiferenca = difDistancia
			indiceMenorDistancia = i
		}
	}

	if indiceMenorDistancia == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(indiceMenorDistancia)
	}

}
