package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct {
	data []int
}

type Iterator struct {
	data  []int
	index int
}

type ReverseIterator struct {
	data  []int
	index int
}

type CyclicIterator struct {
	data  []int
	index int
}

func NewMyList(values []int) *MyList {
	return &MyList{data: values}
}

func (l *MyList) Iterator() *Iterator {
	return &Iterator{data: l.data, index: -1}
}

func (l *MyList) newReverseIterator() *ReverseIterator {

	if len(l.Iterator().data) == 0 {
		return &ReverseIterator{data: l.data}
	}

	newReversI := make([]int, 0)
	ultimoIndex := len(l.data) - 1
	for i := len(l.data) - 1; i > 0; i-- {
		newReversI = append(newReversI, l.data[i])
	}

	return &ReverseIterator{
		data:  newReversI,
		index: ultimoIndex, //indice temp
		//implementar
	}
}

func newCyclicIterator() *CyclicIterator {
	return &CyclicIterator{
		//implementar
	}
}

func (ri *ReverseIterator) HasNext() bool {
	return ri.index < len(ri.data)-1
}

func (ri *ReverseIterator) Next() int {
	if ri.index == len(ri.data) {
		panic(fmt.Errorf("No more elements"))
	}
	ri.index += 1
	return ri.data[ri.index]
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.data)-1
}

func (i *Iterator) Next() int {
	if i.index == len(i.data) {
		panic(fmt.Errorf("No more elements"))
	}
	i.index += 1
	return i.data[i.index]
}

func (c *CyclicIterator) HasNext() bool {
	return c.index < len(c.data)-1
}

func (c *CyclicIterator) Next() int {
	if c.index == len(c.data) {
		panic(fmt.Errorf("No more elements"))
	}
	c.index += 1
	return c.data[c.index]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			break
		case "read":
			for i := 1; i < len(args); i++ {
				slice := make([]int, len(args)-1)
				for i, value := range args[1:] {
					slice[i], _ = strconv.Atoi(value)
				}
				mylist = NewMyList(slice)
			}
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "reverse":
			// fmt.Print("[ ")
			// for it := mylist.ReverseIterator(); it.HasNext(); {
			// 	fmt.Printf("%v ", it.Next())
			// }
			// fmt.Println("]")
		case "cyclic":
			// qtd, _ := strconv.Atoi(args[1])
			// fmt.Print("[ ")
			// it := mylist.CyclicIterator()
			// for range qtd {
			// 	fmt.Printf("%v ", it.Next())
			// }
			// fmt.Println("]")
		}
	}

}
