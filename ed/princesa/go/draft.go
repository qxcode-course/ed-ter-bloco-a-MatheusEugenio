package main

import (
	"fmt"
)

type Jogador struct {
	estaComEspada bool
	id            int
}

func main() {

	var n, e int
	fmt.Scan(&n, &e)

	jogadores := make([]Jogador, n)

	for i := 0; i < len(jogadores); i++ {
		jogadores[i].id = i + 1
	}

	// vai do 0 ao n-1, por isso o e-1
	indiceEspada := e - 1

	jogadores[indiceEspada].estaComEspada = true

	for {

		fmt.Printf("[ ")
		for j := 0; j < len(jogadores); j++ {

			if jogadores[j].estaComEspada == false {
				fmt.Printf("%d ", jogadores[j].id)
			} else {
				fmt.Printf("%d> ", jogadores[j].id)
			}

		}
		fmt.Printf("]\n")

		if len(jogadores) == 1 {
			break
		}

		indiceDoMorto := (indiceEspada + 1) % len(jogadores)

		jogadores = append(jogadores[:indiceDoMorto], jogadores[indiceDoMorto+1:]...)

		if indiceDoMorto < indiceEspada {
			indiceEspada--
		}

		// Agora a espada vai para o próximo vivo
		indiceEspada = (indiceEspada + 1) % len(jogadores)

		//reinicia o status da espada de todos os jogadores, pois ao remover bagunça a ordem dos jogadores
		for i := 0; i < len(jogadores); i++ {
			jogadores[i].estaComEspada = false
		}

		//jogador com indice da espada recebe a espada
		jogadores[indiceEspada].estaComEspada = true

	}
}
