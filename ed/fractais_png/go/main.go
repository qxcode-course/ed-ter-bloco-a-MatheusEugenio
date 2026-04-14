package main

import (
	"fmt"
)

func main() {
	pen := NewPen(300, 300)
	pen.Up()
	pen.SetPosition(50, 75) //x e y do plano
	pen.Down()

	lado := 200.0

	// distanciaDaQuina := (lado / 2) * 1.414

	for range 4 {
		pen.Walk(lado)
		pen.SetHeading(-90) //mudar angulo
	}

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
