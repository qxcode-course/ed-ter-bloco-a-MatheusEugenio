package main

import "fmt"

func main() {

	var h, p, f, d int

	fmt.Scan(&h, &p, &f, &d)

	if d == -1 { //horario ----> "diminuir posições no caminho"

		for {
			if f == p {
				fmt.Print("N\n")
				return
			} else if f == h {
				fmt.Print("S\n")
				return
			} else if f == -1 {
				f = 15
			} else {
				f--
			}
		}

	} else if d == 1 { //anti horario ----> "aumentar posições no caminho"

		for {

			if f == p {
				fmt.Print("N\n")
				return
			} else if f == h {
				fmt.Print("S\n")
				return
			} else if f == 16 {
				f = 0
			} else {
				f++
			}
		}
	}
}
