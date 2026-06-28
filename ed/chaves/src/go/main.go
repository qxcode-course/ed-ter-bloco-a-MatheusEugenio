package main

import "fmt"

func main() {

	times := []rune{
		'A', 'B', 'C', 'D',
		'E', 'F', 'G', 'H',
		'I', 'J', 'K', 'L',
		'M', 'N', 'O', 'P',
	}

	fila := NewQueue[rune]()

	i := 0
	for i < len(times) {
		fila.Enqueue(times[i])
		i++
	}

	var campeao rune

	for fila.items.Len() >= 1 {

		if fila.items.Len() == 1 {
			campeao = fila.Dequeue()
			break
		}

		time1, time2 := fila.Dequeue(), fila.Dequeue()
		var golst1, golst2 int

		fmt.Scan(&golst1, &golst2)

		if golst1 > golst2 {
			fila.Enqueue(time1)
		} else {
			fila.Enqueue(time2)
		}

	}

	fmt.Printf("%c\n", campeao)
}
