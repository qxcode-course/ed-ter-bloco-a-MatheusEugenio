package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	lines  *List[*List[rune]] // texto
	line   *Node[*List[rune]]
	cursor *Node[rune]
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.cursor = e.line.Value.Insert(e.cursor, r)
	e.cursor = e.cursor.Next()
}

func (e *Editor) KeyLeft() {
	if e.cursor != e.line.Value.Front() { // Se o cursor não está no início da linha
		e.cursor = e.cursor.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.line != e.lines.Front() { // Se não está na primeira linha
		e.line = e.line.Prev()        // Move para a linha anterior
		e.cursor = e.line.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() {

	newLine := NewList[rune]()

	if e.cursor != e.line.Value.End() {

		indexCursor := e.line.Value.IndexOf(e.cursor)
		qtdMovida := e.line.Value.size - indexCursor

		primeiroAQuebrar := e.cursor
		ultimoAQuebrar := e.line.Value.End().prev
		ultimoQueFica := e.cursor.prev

		newLine.root.next = primeiroAQuebrar
		primeiroAQuebrar.prev = newLine.root

		newLine.root.prev = ultimoAQuebrar
		ultimoAQuebrar.next = newLine.root

		ultimoQueFica.next = e.line.Value.root
		e.line.Value.root.prev = ultimoQueFica

		newLine.size = qtdMovida
		e.line.Value.size -= qtdMovida
	}

	e.lines.Insert(e.line.Next(), newLine)
	e.line = e.line.Next()
	e.cursor = e.line.Value.Front()
}

func (e *Editor) KeyRight() {
	if e.cursor != e.line.Value.End() {
		e.cursor = e.cursor.Next()
	}

	if e.cursor == e.line.Value.End() {

		if e.line.Next() != e.line.root {
			e.line = e.line.Next()
			e.cursor = e.line.Value.Front()
		}
		return
	}
}

func (e *Editor) KeyUp() {

	if e.line == e.lines.Front() {
		return
	}

	linhaAnterior := e.line.prev

	colunaAtual := e.line.Value.IndexOf(e.cursor)

	novoNoCursor := linhaAnterior.Value.Front()

	if colunaAtual >= linhaAnterior.Value.size {
		novoNoCursor = linhaAnterior.Value.root
	} else {
		for i := 0; i < colunaAtual; i++ {
			novoNoCursor = novoNoCursor.Next()
		}
	}

	e.line = linhaAnterior
	e.cursor = novoNoCursor
}

func (e *Editor) KeyDown() {

	if e.line == e.lines.Back() {
		return
	}

	linhaDeBaixo := e.line.Next()

	colunaAtualCursor := e.line.Value.IndexOf(e.cursor)

	novoNoCursor := linhaDeBaixo.Value.Front()

	if colunaAtualCursor >= linhaDeBaixo.Value.size {
		novoNoCursor = linhaDeBaixo.Value.root
	} else {
		for i := 0; i < colunaAtualCursor; i++ {
			novoNoCursor = novoNoCursor.Next()
		}
	}

	e.line = linhaDeBaixo
	e.cursor = novoNoCursor
}

func (e *Editor) KeyBackspace() {
	noDeTras := e.cursor.prev

	if noDeTras == e.line.Value.root {

		if e.line.Prev() != nil {

			linhaAnterior := e.line.Prev()
			linhaAtual := e.line

			novoNoCursor := linhaAnterior.Value.End()

			if linhaAtual.Value.size > 0 {

				ultimaPosLinhaAnterior := linhaAnterior.Value.End().prev
				primeiroDaLinhaAtual := linhaAtual.Value.Front()
				ultimoDaLinhaAtual := linhaAtual.Value.Back()

				ultimaPosLinhaAnterior.next = primeiroDaLinhaAtual
				primeiroDaLinhaAtual.prev = ultimaPosLinhaAnterior

				// Costura o final da linha atual no root da linha anterior
				ultimoDaLinhaAtual.next = linhaAnterior.Value.root
				linhaAnterior.Value.root.prev = ultimoDaLinhaAtual

				linhaAnterior.Value.size += linhaAtual.Value.size
			}

			linhaAtual.Value.root.next = linhaAtual.Value.root
			linhaAtual.Value.root.prev = linhaAtual.Value.root
			linhaAtual.Value.size = 0

			e.line = linhaAnterior

			e.cursor = novoNoCursor

			e.lines.Erase(linhaAtual)
		}

		return
	}

	e.line.Value.Erase(noDeTras)
	noDeTras.next = nil
	noDeTras.prev = nil
}

func (e *Editor) KeyDelete() {

	noDaFrente := e.cursor.next

	if noDaFrente == e.line.Value.root {

		if e.line.Next() != nil {

			linhaAtual := e.line
			linhaDaFrente := e.line.Next()

			if linhaAtual.Value.size > 0 {

				ultimoNOLinhaAtual := linhaAtual.Value.End().prev
				ultimoNOlinhaDaFrente := linhaDaFrente.Value.End().prev

				primeiroNOLinhaDaFrente := linhaDaFrente.Value.Front()

				ultimoNOLinhaAtual.next = primeiroNOLinhaDaFrente
				primeiroNOLinhaDaFrente.prev = ultimoNOLinhaAtual

				ultimoNOlinhaDaFrente.next = linhaAtual.Value.root
				linhaAtual.Value.root.prev = ultimoNOlinhaDaFrente.next

				linhaAtual.Value.size += linhaDaFrente.Value.size
			}

			linhaDaFrente.Value.root.next = linhaDaFrente.Value.root
			linhaDaFrente.Value.root.prev = linhaDaFrente.Value.root
			linhaDaFrente.Value.size = 0

			e.lines.Erase(linhaDaFrente)
		}

		return
	}

	e.line.Value.Erase(noDaFrente)
	noDaFrente.next = nil
	noDaFrente.prev = nil
}

func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.lines = NewList[*List[rune]]()
	e.lines.PushBack(NewList[rune]())
	e.line = e.lines.Front()
	e.cursor = e.line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.lines.Front(); line != e.lines.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.cursor {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
