package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(newCapacity int) *MultiSet {
	return &MultiSet{
		data:     make([]int, newCapacity),
		size:     0,
		capacity: newCapacity,
	}
}

func (ms *MultiSet) Search(value int) (bool, int) {

	low, high := 0, ms.size-1

	for low <= high {

		meio := (low + high) / 2

		if ms.data[meio] < value {
			low = meio + 1
		} else if ms.data[meio] > value {
			high = meio - 1
		} else {
			return true, meio
		}
	}

	return false, low
}

func (ms *MultiSet) insert(value int, index int) error {

	if ms.size == ms.capacity {
		ms.expand()
	}

	if index < 0 || index > ms.capacity {
		return errors.New("value not found")
	}

	for i := ms.size - 1; i >= index; i-- {
		ms.data[i+1] = ms.data[i]
	}
	ms.size++
	ms.data[index] = value
	return nil
}

func (ms *MultiSet) Insert(value int) {

	existe, index := ms.Search(value)
	if existe {
		//logica de adicionar mais 1 na qtd do número
		ms.insert(value, index)
	} else {
		ms.insert(value, index)
	}
}

func (ms *MultiSet) erase(index int) error {

	if index < 0 || index > ms.capacity {
		return errors.New("value not found")
	}

	for i := index; i < ms.size; i++ {
		ms.data[i] = ms.data[i+1]
	}
	ms.size--
	return nil
}

func (ms *MultiSet) Erase(value int) error {

	existe, index := ms.Search(value)
	if existe {
		err := ms.erase(index)
		if err != nil {
			return errors.New("value not found")
		} else {
			return nil
		}
	}
	return errors.New("value not found")
}

func (ms *MultiSet) Contains(value int) bool {

	existe, _ := ms.Search(value)
	if existe {
		return true
	}
	return false
}

func (ms *MultiSet) Count(value int) int {

	return 0
}

func (ms *MultiSet) Unique() int {

	return 0
}

func (ms *MultiSet) String() string {
	return "[" + Join(ms.data[:ms.size], ", ") + "]"
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func (ms *MultiSet) expand() {

	if ms.capacity == 0 {
		ms.capacity = 1
	} else {
		ms.capacity *= 2
	}

	dataRef := make([]int, ms.capacity)

	for i, value := range ms.data {
		dataRef[i] = value
	}

	ms.data = dataRef
}

func (ms *MultiSet) Clear() {
	ms.size = 0
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
