package main

import (
	"fmt"
)

func desenhaArvore(caneta *Pen, comprimento float64) {

	if comprimento < 4.5 {
		return
	}

	caneta.Walk(comprimento)

	caneta.Left(35)
	desenhaArvore(caneta, comprimento*0.8)
	caneta.Right(35)

	caneta.Right(35)
	desenhaArvore(caneta, comprimento*0.8)
	caneta.Left(35)

	caneta.Up()
	caneta.Walk(-comprimento)
	caneta.Down()
}

func desenhaArvoreComFruta(caneta *Pen, comprimento float64) {

	if comprimento < 4.5 {
		return
	}

	caneta.Walk(comprimento)

	caneta.Left(35)
	desenhaArvore(caneta, comprimento*0.8)

	caneta.SetRGB(255, 0, 0)
	caneta.FillCircle(1.5)

	caneta.Right(35)

	caneta.SetRGB(0, 0, 0)

	caneta.Right(35)
	desenhaArvore(caneta, comprimento*0.8)

	caneta.SetRGB(0, 0, 0)

	caneta.FillCircle(1.5)
	caneta.Left(35)

	caneta.Up()
	caneta.Walk(-comprimento)
	caneta.Down()

}

func main() {
	pen := NewPen(300, 300)
	pen.Up()
	pen.Down()

	pen.SetLineWidth(0.7)
	pen.SetHeading(90)

	var comprimento float64
	comprimento = 49
	desenhaArvore(pen, comprimento)

	// desenhaArvoreComFruta(pen, comprimento)

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
